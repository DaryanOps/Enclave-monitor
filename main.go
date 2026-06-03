package main

import (
    "fmt"
    "runtime"
    "time"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/net"
)

type SystemInfo struct {
    OS           string
    Architecture string
    CPUs         int
    Time         string
    TotalMemory  uint64
    UsedMemory   uint64
    CPUUsage     float64
    TotalDisk    uint64
    UsedDisk     uint64
    DiskUsage    float64
    BytesSent    uint64
    BytesRecv    uint64
}

func getSystemInfo() SystemInfo {
    vmStat, _     := mem.VirtualMemory()
    cpuPercent, _ := cpu.Percent(time.Second, false)
    diskStat, _   := disk.Usage("/")
    netStat, _    := net.IOCounters(false)

    return SystemInfo{
        OS:           runtime.GOOS,
        Architecture: runtime.GOARCH,
        CPUs:         runtime.NumCPU(),
        Time:         time.Now().Format("2006-01-02 15:04:05"),
        TotalMemory:  vmStat.Total / 1024 / 1024 / 1024,
        UsedMemory:   vmStat.Used / 1024 / 1024 / 1024,
        CPUUsage:     cpuPercent[0],
        TotalDisk:    diskStat.Total / 1024 / 1024 / 1024,
        UsedDisk:     diskStat.Used / 1024 / 1024 / 1024,
        DiskUsage:    diskStat.UsedPercent,
        BytesSent:    netStat[0].BytesSent / 1024 / 1024,
        BytesRecv:    netStat[0].BytesRecv / 1024 / 1024,
    }
}

func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

func printInfo(info SystemInfo) {
    fmt.Println("=== Enclave Monitor ===")
    fmt.Println("OS:          ", info.OS)
    fmt.Println("Architecture:", info.Architecture)
    fmt.Println("CPUs:        ", info.CPUs)
    fmt.Println("Time:        ", info.Time)
    fmt.Println("Total Memory:", info.TotalMemory, "GB")
    fmt.Println("Used Memory: ", info.UsedMemory, "GB")
    fmt.Printf("CPU Usage:    %.2f%%\n", info.CPUUsage)
    fmt.Printf("Disk Usage:   %.2f%%\n", info.DiskUsage)
    fmt.Println("Total Disk:  ", info.TotalDisk, "GB")
    fmt.Println("Used Disk:   ", info.UsedDisk, "GB")
    fmt.Println("Network Sent:", info.BytesSent, "MB")
    fmt.Println("Network Recv:", info.BytesRecv, "MB")
    fmt.Println("=======================")
    fmt.Println("Press Ctrl+C to exit")
}

func main() {
    for {
        clearScreen()
        info := getSystemInfo()
        printInfo(info)
        time.Sleep(2 * time.Second)
    }
}
