package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/mem"
)

const (
	CRITICAL_EXIT = 2
	WARNING_EXIT  = 1
	OK_EXIT       = 0
)

func main() {
	warning := flag.Int("warning", 80, "Warning threshold for memory.")
	critical := flag.Int("critical", 90, "Critical threshold for memory.")
	flag.Parse()

	m, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Can't get information - error %v.\n", err)
		os.Exit(3)
	}

	total := m.Total / 1024 / 1024
	used := m.Used / 1024 / 1024
	usedPercent := int(m.UsedPercent)

	var status string
	var exitCode int

	switch {
	case usedPercent >= *critical:
		status = "CRITICAL"
		exitCode = CRITICAL_EXIT
	case usedPercent >= *warning:
		status = "WARNING"
		exitCode = WARNING_EXIT
	default:
		status = "OK"
		exitCode = OK_EXIT
	}

	fmt.Printf("%s - Memory percent usage %v%% - Total Memory %v Mo - Used Memory %v Mo | mem_percent = %v,%v,%v,0,100\n",
		status, usedPercent, total, used, usedPercent, *warning, *critical)
	os.Exit(exitCode)
}
