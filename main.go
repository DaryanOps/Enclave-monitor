package main

import (
    "fmt"
    "runtime"
    "time"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/cpu"
)

type SystemInfo struct {
    OS           string
    Architecture string
    CPUs         int
    Time         string
    TotalMemory  uint64
    UsedMemory   uint64
    CPUUsage     float64
}

func getSystemInfo() SystemInfo {
    vmStat, _ := mem.VirtualMemory()
    cpuPercent, _ := cpu.Percent(time.Second, false)

    return SystemInfo{
        OS:           runtime.GOOS,
        Architecture: runtime.GOARCH,
        CPUs:         runtime.NumCPU(),
        Time:         time.Now().Format("2006-01-02 15:04:05"),
        TotalMemory:  vmStat.Total / 1024 / 1024 / 1024,
        UsedMemory:   vmStat.Used / 1024 / 1024 / 1024,
        CPUUsage:     cpuPercent[0],
    }
}

func main() {
    info := getSystemInfo()

    fmt.Println("=== Enclave Monitor ===")
    fmt.Println("OS:          ", info.OS)
    fmt.Println("Architecture:", info.Architecture)
    fmt.Println("CPUs:        ", info.CPUs)
    fmt.Println("Time:        ", info.Time)
    fmt.Println("Total Memory:", info.TotalMemory, "GB")
    fmt.Println("Used Memory: ", info.UsedMemory, "GB")
    fmt.Printf("CPU Usage:    %.2f%%\n", info.CPUUsage)
    fmt.Println("=======================")
}
