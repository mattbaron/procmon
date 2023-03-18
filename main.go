package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mattbaron/procmon/influx"
)

const Version = "0.1.0"

func main() {
	versionFlag := flag.Bool("version", false, "Display version and exit")
	if *versionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}

	processList, err := NewProcessList()
	if err != nil {
		os.Exit(2)
	}

	for _, processName := range os.Args[1:] {
		processName := strings.TrimSpace(processName)
		if len(processName) == 0 {
			continue
		}

		line := influx.NewLine("procmon")
		line.AddTag("name", processName)
		line.AddField("count", processList.GetProcessCount(processName))

		fmt.Println(line)
	}
}
