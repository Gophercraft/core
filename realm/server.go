package realm

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/Gophercraft/core/packet/auth"
	"github.com/Gophercraft/core/vsn"

	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
)

type Server struct {
	Config          *config.World
	DB              *wdb.Core
	GuardPhases     sync.Mutex
	Phases          map[string]*Phase
	GuardPlayerList sync.Mutex
	PlayerList      map[string]*Session
	PackLoader      *datapack.Loader
	Plugins         []*LoadedPlugin

	eventMgr          sync.Map
	ProtocolHandlers  ProtocolHandlers
	CommandHandlers   []Command
	HomeServiceClient rpcnet.HomeServiceClient
	WaitQueue         WaitQueue
	tlsConfig         *tls.Config
	DynamicCounters   map[guid.TypeID]uint64
	GuardCounters     sync.Mutex
	StartTime         time.Time
	// TerrainMgr
	// Misc data stores
	LevelExperience models.LevelExperience
	// PlayerCreateInfo          []models.PlayerCreateInfo
	// PlayerCreateItems         []models.PlayerCreateItem
	// PlayerCreateActionButtons []models.PlayerCreateActionButton
	// PlayerCreateAbilities     []models.PlayerCreateAbility

	SpellEffects []SpellEffect
	AuraEffects  []AuraEffect
}

func Start(opts *config.World) error {
	if opts.CPUProfile != "" {
		f, err := os.Create(opts.CPUProfile)
		if err != nil {
			return err
		}

		pprof.StartCPUProfile(f)
	}

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Warn("Stopping realm server")
		if opts.CPUProfile != "" {
			log.Warn("Closing CPU profile")
			pprof.StopCPUProfile()
		}
		os.Exit(0)
	}()

	ws := &Server{}
	ws.InitConfig(opts)
	ws.Phases = make(map[string]*Phase)
	ws.PlayerList = make(map[string]*Session)
	ws.DynamicCounters = make(map[guid.TypeID]uint64)
	ws.StartTime = time.Now()
	// Open database
	core, err := wdb.NewCore(opts.DBDriver, opts.DBURL)
	if err != nil {
		return err
	}

	ws.DB = core

	ws.initHandlers()
	ws.initSpellEffects()
	ws.initAuraEffects()

	if err := ws.loadPlugins(); err != nil {
		return err
	}

	// Open handles to ZIP archives and indices of flat folders.
	ws.PackLoader, err = datapack.Open(filepath.Join(ws.Config.Dir, "Datapacks"))
	if err != nil {
		return err
	}

	if err := ws.LoadDatapacks(); err != nil {
		return err
	}

	// Register info with Homeserver
	go ws.phoneHome()

	if opts.Version == vsn.Alpha {
		go ws.serveRedirect()
	}

	characterCt, err := ws.DB.Engine.Count(new(models.Character))
	if err != nil {
		return err
	}

	log.Println("Gophercraft Core World Server successfully initialized!")
	log.Println(characterCt, "characters exist on this realm.")

	l, err := net.Listen("tcp", opts.Listen)
	if err != nil {
		return err
	}

	log.Println("Worldserver started to listen at", opts.Listen)

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go ws.HandleConn(c)
	}
}

func (ws *Server) UptimeMS() uint32 {
	return uint32(time.Since(ws.StartTime) / time.Millisecond)
}

func (ws *Server) Build() vsn.Build {
	return vsn.Build(ws.Config.Version)
}

func (ws *Server) RealmID() uint64 {
	return ws.Config.RealmID
}

func (s *Server) GetSessionByGUID(g guid.GUID) (*Session, error) {
	s.GuardPlayerList.Lock()
	defer s.GuardPlayerList.Unlock()
	for _, v := range s.PlayerList {
		if v.HasState(InWorld) {
			if v.GUID() == g {
				return v, nil
			}
		}
	}

	return nil, fmt.Errorf("could not find session corresponding to input")
}

func (ws *Server) GetPlayerNameByGUID(g guid.GUID) (string, error) {
	plyr, err := ws.GetSessionByGUID(g)
	if err == nil {
		return plyr.PlayerName(), nil
	}

	var chr models.Character

	found, err := ws.DB.Where("id = ?", g.Counter()).Get(&chr)
	if err != nil {
		return "", err
	}

	if !found {
		return "", fmt.Errorf("no character found for guid %s", g)
	}

	return chr.Name, nil
}

func (s *Server) GetUnitNameByGUID(g guid.GUID) (string, error) {
	if g == guid.Nil {
		return "", nil
	}

	switch g.HighType() {
	case guid.Player:
		plyr, err := s.GetSessionByGUID(g)
		if err != nil {
			return "", err
		}
		return plyr.PlayerName(), nil
	case guid.Creature:
		return "", fmt.Errorf("npc names nyi")
	default:
		return "", fmt.Errorf("cannot name this type (%s)", g.HighType())
	}
}

func (s *Server) GetSessionByPlayerName(playerName string) (*Session, error) {
	s.GuardPlayerList.Lock()
	session := s.PlayerList[playerName]
	s.GuardPlayerList.Unlock()
	if session != nil {
		return session, nil
	}

	return nil, fmt.Errorf("no session for player '%s'", playerName)
}

func (s *Server) GetGUIDByPlayerName(playerName string) (guid.GUID, error) {
	s.GuardPlayerList.Lock()
	session, found := s.PlayerList[playerName]
	s.GuardPlayerList.Unlock()
	if found {
		return session.GUID(), nil
	}

	var chr models.Character
	found, err := s.DB.Where("name = ?", playerName).Get(&chr)
	if err != nil {
		return guid.Nil, err
	}

	if !found {
		return guid.Nil, fmt.Errorf("no player by the name of %s", playerName)
	}

	return guid.RealmSpecific(guid.Player, s.RealmID(), chr.ID), nil
}

func (s *Session) SendAuthWaitQueue(position uint32) {
	s.Send(&auth.Response{
		WaitInfo: &auth.WaitInfo{
			WaitCount: position,
		},
	})
}

type ServerStats struct {
	Allocated      uint64
	TotalAllocated uint64
	SystemMemory   uint64
	NumGCCycles    uint32
	Goroutines     int
	Uptime         time.Duration
	CacheSize      uint64
}

func (ws *Server) GetServerStats() *ServerStats {
	sstats := &ServerStats{}
	sstats.Goroutines = runtime.NumGoroutine()
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	sstats.Allocated = memStats.Alloc
	sstats.TotalAllocated = memStats.TotalAlloc
	sstats.SystemMemory = memStats.Sys
	sstats.NumGCCycles = memStats.NumGC

	sstats.CacheSize = ws.DB.Cache.Size()
	sstats.Uptime = time.Since(ws.StartTime)
	return sstats
}
