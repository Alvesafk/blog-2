package serverstats

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/mem"
)

type ServerStats struct {
	Goroutines int `json:"goroutines"`
	UsedMemory float64 `json:"usedMemory"`
}

func NewServerStats() *ServerStats {
	return &ServerStats{
		Goroutines: runtime.NumGoroutine(),
		UsedMemory: getUsedMemory(),
	}
}

func getUsedMemory() float64 {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return -1
	}

	return mem.UsedPercent
}
