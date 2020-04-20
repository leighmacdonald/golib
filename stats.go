package util

import (
	"runtime"
)

func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type SysStats struct {
	Alloc      uint64
	TotalAlloc uint64
	Sys        uint64
	NumGC      uint32
	GoRoutines int
}

func GetStats() SysStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	gor := runtime.NumGoroutine()
	return SysStats{
		Alloc:      m.Alloc,
		TotalAlloc: m.TotalAlloc,
		Sys:        m.Sys,
		NumGC:      m.NumGC,
		GoRoutines: gor,
	}
}
