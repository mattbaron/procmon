package main

import (
	"fmt"
	"os"

	"github.com/mattbaron/procmon/influx"
)

func main() {
	processList, err := NewProcessList()
	if err != nil {
		os.Exit(2)
	}

	for _, processName := range os.Args[1:] {
		line := influx.NewLine("procmon")
		line.AddTag("name", processName)
		line.AddField("count", processList.GetProcessCount(processName))
		fmt.Println(line)
	}
}
