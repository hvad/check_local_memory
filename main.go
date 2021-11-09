package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/mem"
)

func main() {

	warning := flag.Int("warning", 80, "Warning thresold for memory.")
	critical := flag.Int("critical", 90, "Critical thresold for memory.")
	flag.Parse()

	m, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Can't get information - error %v.\n", err)
		os.Exit(3)
	}

	total := m.Total / 1024 / 1024

	used := m.Used / 1024 / 1024

	if int(m.UsedPercent) >= *critical {
		fmt.Printf("CRITICAL - Memory percent usage %v%% - Total Memory %v Mo - Used Memory %v Mo\n", int(m.UsedPercent), total, used)
		os.Exit(2)

	} else if int(m.UsedPercent) >= *warning {
		fmt.Printf("WARNING - Memory percent usage %v%% - Total Memory %v Mo - Used Memory %v Mo\n", int(m.UsedPercent), total, used)
		os.Exit(1)

	} else {
		fmt.Printf("OK - Memory percent usage %v%% - Total Memory %v Mo - Used Memory %v Mo\n", int(m.UsedPercent), total, used)
		os.Exit(0)
	}

}
