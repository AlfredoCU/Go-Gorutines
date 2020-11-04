package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		opc = ""
		condition = true
		flag = make(chan bool)
		scan = bufio.NewScanner(os.Stdin)

		processId uint64
		newProcess *Process
	 	processIdCreate uint64
	 	processIdDelete string

		process []*Process
		processAdmin = &ProcessManager {
			Processes: process,
		}
	)

	for exit := true; exit; exit = condition {
		fmt.Println("******* Process Manager *******")
		fmt.Println("[1] Add Process")
		fmt.Println("[2] Show Process")
		fmt.Println("[3] Delete Process")
		fmt.Println("[4] Exit System")
		fmt.Print("-Option: ")
		scan.Scan()
		opc = scan.Text()

		switch opc {
			case "1":
				newProcess = NewProcess(processIdCreate)
				processAdmin.AddProcess(newProcess)
				processIdCreate += 1
				go newProcess.Start()
				fmt.Print("\n-Add process #", processIdCreate, "\n\n")
				break
			case "2":
				if len(processAdmin.Processes) != 0 {
					go Concurrently(processAdmin, flag)
					scan.Scan()
					flag <- true
				} else {
					fmt.Print("\n-No process!\n\n")
				}
				break
			case "3":
				fmt.Print("\n-Enter the process ID to remove: ")
				scan.Scan()
				processIdDelete = scan.Text()
				processId, _ = strconv.ParseUint(processIdDelete, 10, 64)
				if processAdmin.KillProcess(processId) {
					fmt.Print("-Process ", processId, " removed successfully.\n\n")
				} else {
					fmt.Print("-No process found!\n\n")
				}
				processId = 0
				break
			case "4":
				processAdmin.exited()
				condition = false
				break
			default:
				invalidOptions()
		}
	}
}

// Process struct.
type Process struct {
	Id uint64
	Task uint64
	IsRunning bool
}

// Start function.
func (process *Process) Start() {
	process.Task = 0
	process.IsRunning = true

	for {
		process.Task += 1
		time.Sleep(time.Millisecond * 500)
		if !process.IsRunning {
			break
		}
	}
}

// Stop function.
func (process *Process) Stop() {
	process.IsRunning = false
}

// NewProcess function.
func NewProcess(id uint64) *Process {
	return &Process{
		Id: id,
	}
}

// ProcessManager struct.
type ProcessManager struct {
	Processes []*Process
}

// AddProcess function.
func (processAdmin *ProcessManager) AddProcess(process *Process) {
	processAdmin.Processes = append(processAdmin.Processes, process)
}

// KillProcess function.
func (processAdmin *ProcessManager) KillProcess(processId uint64) bool {
	var newProcess []*Process
	deleted := false

	for _, process := range processAdmin.Processes {
		if process.Id != processId {
			newProcess = append(newProcess, process)
		}

		if process.Id == processId {
			deleted = true
			process.Stop()
		}
	}

	processAdmin.Processes = newProcess
	return deleted
}

// ShowProcess function.
func (processAdmin *ProcessManager) ShowProcess() {
	for _, process := range processAdmin.Processes {
		fmt.Print("\nId ", process.Id, " : ", process.Task)
	}
	fmt.Println()
}

// Concurrently function.
func Concurrently(processAdmin *ProcessManager, flag chan bool) {
	for {
		select {
		case <-flag:
			return
		default:
			processAdmin.ShowProcess()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// invalidOptions function.
func invalidOptions() {
	fmt.Print("\n-Invalid Option!\n\n")
}

// exited function.
func (processAdmin *ProcessManager) exited() {
	for _, process := range processAdmin.Processes {
		process.Stop()
	}
	fmt.Println("\n-System exited...")
}
