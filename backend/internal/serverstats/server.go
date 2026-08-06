package serverstats

import (
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type ServerStats struct {
	Goroutines       int     `json:"goroutines"`
	UsedMemory       float64 `json:"usedMemory"`
	UsedCPU          float64 `json:"usedCPU"`
	TotalConnections int64   `json:"totalConnections"`
}

func NewServerStats(cw *ConnectionWatcher) *ServerStats {
	return &ServerStats{
		Goroutines: runtime.NumGoroutine(),
		UsedMemory: getUsedMemory(),
		UsedCPU:    getUsedCPU(),
		TotalConnections: cw.Load(),
	}
}

func getUsedMemory() float64 {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return -1
	}

	return mem.UsedPercent
}

func getUsedCPU() float64 {
	cpuSlice, err := cpu.Percent(0, true)
	if err != nil {
		return -1
	}

	var total float64
	for _, v := range cpuSlice {
		total += v
	}

	return total / float64(len(cpuSlice))
}
