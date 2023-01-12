package commands

import (
	"reflect"
	"sort"
	"strings"

	"github.com/Gophercraft/core/realm"
)

type CommandProvider struct {
}

func (c *CommandProvider) Activated() (bool, error) {
	return true, nil
}

func (c *CommandProvider) Terminate() error {
	return nil
}

func (c *CommandProvider) Init(s *realm.Server, info *realm.PluginInfo) error {
	info.ID = "gophercraft.commands"
	info.Name = "Gophercraft Core commands"

	s.Cmd(realm.GameMaster, "help", "Show this menu", cmdHelp)

	s.Cmd(realm.GameMaster, "modify scale", "Changes the scale of your player model.\n  .modify scale 1.5", cmdScale)
	s.Cmd(realm.GameMaster, "modify speed", "Changes all movement speeds of your character. For 2X speed:\n  .modify speed 2.0", cmdSpeed)
	s.Cmd(realm.GameMaster, "modify coinage", "Add an amount of gold to your balance.\n  To remove 500 gold, 10 silver: .modify gold -500g10s", cmdMoney)
	s.Cmd(realm.GameMaster, "modify level", "Change your character's level", cmdModLevel)
	s.Cmd(realm.GameMaster, "morph", "Changes your displayID to the supplied integer.", cmdMorph)
	s.Cmd(realm.GameMaster, "set fly", "Enable flight", cmdFly)
	s.Cmd(realm.GameMaster, "set gamemode", "Sets a gamemode.", cmdSetGamemode)

	s.Cmd(realm.PhaseBuilder|realm.PhaseOwner, "npc add", "Spawn a Creature at your current position.", cmdAddNPC)

	s.Cmd(realm.PhaseBuilder|realm.PhaseOwner, "gobject add", "Spawn a GameObject at your current position.", cmdSpawnGameobject)

	s.Cmd(realm.PhaseOwner, "play music", "Play music throughout the current map.", cmdPlayMusic)
	s.Cmd(realm.PhaseOwner, "play sound", "Play sound throughout the current map.", cmdPlaySound)
	s.Cmd(realm.PhaseOwner, "play visual", "Play a spell visual ID for the targeted unit", cmdPlayVisual)

	s.Cmd(realm.GameMaster, "tele", "Teleport to a Port ID", cmdTele)

	s.Cmd(realm.GameMaster, "xgps", "To move 300 forward:\n  .xgps 300 f", cmdXGPS)
	s.Cmd(realm.GameMaster, "gps", "Query current position and area information.", cmdGPS)
	s.Cmd(realm.GameMaster, "summon", "Summons the named player to your location with their consent.", cmdSummon)
	s.Cmd(realm.GameMaster, "appear", "Teleport to the named player.", cmdAppear)

	s.Cmd(realm.GameMaster, "lookup teleport", "Find a teleport location.", cmdLookupTeleport)
	s.Cmd(realm.GameMaster, "lookup item", "Find a teleport location.", cmdLookupItem)
	s.Cmd(realm.GameMaster, "lookup gobject", "Find a gameobject.", cmdLookupGameObject)

	s.Cmd(realm.GameMaster, "additem", "Add an item.", cmdAddItem)

	s.Cmd(realm.GameMaster, "teach", "Teaches a spell to the target", cmdTeach)

	s.Cmd(realm.GameMaster, "aura list", "List aura info of your target", cmdAuraList)
	s.Cmd(realm.GameMaster, "aura add", "Add an aura to your target", cmdAuraAdd)

	s.Cmd(realm.GameMaster, "revive", "Revive yourself or a target", cmdRevive)

	// Admin utilities
	s.Cmd(realm.Admin, "debug pupd8", "(Debug) Send an update packet to your client", cmdRefreshPlayer)
	s.Cmd(realm.Admin, "debug trackedguids", "(Debug) List currently tracked objects", cmdTrackedGUIDs)
	s.Cmd(realm.Admin, "debug map", "(Debug) Show map info", cmdDebugMapInfo)
	s.Cmd(realm.Admin, "value", "(Debug) Modify a value of your Player object directly", cmdVmod)
	s.Cmd(realm.Admin, "sstats", "(Debug) Query current server usage statistics", cmdStats)
	s.Cmd(realm.Admin, "funcdump", "(Debug) Dump currently running goroutine stack traces to stdout", cmdGoroutines)
	s.Cmd(realm.Admin, "showsql", "(Debug) Display XORM queries to stdout", cmdShowSQL)
	s.Cmd(realm.Admin, "debug inv", "(Debug) Show contents of inventory manager", cmdDebugInv)
	s.Cmd(realm.Admin, "prop list", "(Debug)", cmdListProps)
	s.Cmd(realm.Admin, "prop toggle", "(Debug)", cmdToggleBoolProp)
	s.Cmd(realm.Admin, "prop list", "(Debug)", cmdRemoveProp)

	return nil
}

func init() {
	realm.RegisterPlugin("commands", &CommandProvider{})
}

func getCommandPrefixes(s *realm.Session) []string {
	var set = make(map[string]bool)
	var out []string
	for _, ch := range s.Server.CommandHandlers {
		if !strings.Contains(ch.Signature, " ") {
			out = append(out, ch.Signature)
		} else {
			split := strings.Split(ch.Signature, " ")

			prefix := split[0]

			if !set[prefix] {
				set[prefix] = true
				out = append(out, prefix+"...")
			}
		}
	}

	sort.Strings(out)
	return out
}

var (
	listColor = "ff00ffea"
)

func cmdHelp(c *realm.Session, prefix string) {
	var perLine int
	var line string

	if prefix == "" {
		prefixes := getCommandPrefixes(c)

		for _, command := range prefixes {
			if perLine == 0 {
				line += "|c" + listColor
			}

			if perLine > 0 {
				line += "   "
			}

			line += command

			// refl := reflect.ValueOf(command.Function)

			// for x := 1; x < refl.Type().NumIn(); x++ {
			// 	parameters += " <" + refl.Type().In(x).String() + ">"
			// }

			if perLine > 3 {
				c.SystemChat(line + "|r")
				perLine = 0
				line = ""
			} else {
				perLine++
			}

			// c.ColorPrintf(realm.DemoColor, "  %s", command.Description)
		}

		if line != "" {
			c.SystemChat(line + "|r")
		}
	} else {
		cmdHelpSpecific(c, prefix)
	}
}

func cmdHelpSpecific(c *realm.Session, prefix string) {
	var matches []*realm.Command

	for i := range c.Server.CommandHandlers {
		cmd := &c.Server.CommandHandlers[i]
		if strings.HasPrefix(cmd.Signature, prefix) {
			matches = append(matches, cmd)
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Signature < matches[j].Signature
	})

	for _, cmd := range matches {
		parameters := ""
		refl := reflect.ValueOf(cmd.Function)

		for x := 1; x < refl.Type().NumIn(); x++ {
			parameters += " <" + refl.Type().In(x).String() + ">"
		}

		c.ColorPrintf(realm.HelpColor, ".%s%s", cmd.Signature, parameters)
		c.ColorPrintf(realm.DemoColor, "  %s", cmd.Description)
	}
}
