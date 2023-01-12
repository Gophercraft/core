package commands

import (
	"runtime"

	"github.com/Gophercraft/core/realm"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func cmdStats(s *realm.Session) {
	stats := s.Server.GetServerStats()
	s.SystemChat("|cff34dceb|r ~~ Server Statistics ~~")
	s.Kv("System", "%s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	s.Kv("CPUs", "%d", runtime.NumCPU())
	s.Kv("Server uptime", "%v", stats.Uptime)
	s.Kv("Number of active goroutines", "%d", stats.Goroutines)
	s.Kv("Total bytes allocated by heap", "%d MiB", bToMb(stats.TotalAllocated))
	s.Kv("Current memory allocated for server", "%d MiB", bToMb(stats.SystemMemory))
	s.Kv("Current memory usage of server", "%d MiB", bToMb(stats.Allocated))
	s.Kv("Total GC cycles", "%d", stats.NumGCCycles)
	s.Kv("Cache size on disk", "%d", stats.CacheSize)
}
