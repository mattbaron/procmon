package main

import (
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessList struct {
	Procs                []*process.Process
	CountByName          map[string]int
	CountByNameLowerCase map[string]int
}

func NewProcessList() (*ProcessList, error) {
	processList := &ProcessList{
		Procs:                make([]*process.Process, 0),
		CountByName:          make(map[string]int),
		CountByNameLowerCase: make(map[string]int),
	}

	allProcesses, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, proc := range allProcesses {
		processList.Add(proc)
	}

	return processList, nil
}

func (processList *ProcessList) Add(process *process.Process) {
	processList.Procs = append(processList.Procs, process)

	name, err := process.Name()
	if err != nil {
		return
	}

	processList.CountByName[name] += 1
	processList.CountByNameLowerCase[strings.ToLower(name)] += 1
}

func (processList *ProcessList) GetProcessCount(name string) int {
	lowerCaseName := strings.ToLower(name)
	value, ok := processList.CountByNameLowerCase[lowerCaseName]
	if ok {
		return value
	} else {
		return 0
	}
}
