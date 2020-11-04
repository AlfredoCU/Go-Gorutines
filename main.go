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
		processAdmin = &ProcessAdmin {
			ProcessAdmin: process,
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
				if processAdmin.ProcessLength != 0 {
					go Concurrently(processAdmin, flag)
					scan.Scan()
					flag <- true
				}
				break
			case "3":
				fmt.Print("\n-Enter the process ID to remove: ")
				scan.Scan()
				processIdDelete = scan.Text()
				processId, _ = strconv.ParseUint(processIdDelete, 10, 64)
				if processAdmin.KillProcess(processId) {
					fmt.Print("-Process",processId,"removed successfully.\n\n")
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

func Concurrently(processAdmin *ProcessAdmin, flag chan bool) {
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

// ProcessAdmin function.
type ProcessAdmin struct {
	ProcessAdmin []*Process
	ProcessLength uint64
}

// AddProcess function.
func (processAdmin *ProcessAdmin) AddProcess(process *Process) {
	processAdmin.ProcessAdmin = append(processAdmin.ProcessAdmin, process)
	processAdmin.ProcessLength += 1
}

// KillProcess function.
func (processAdmin *ProcessAdmin) KillProcess(processId uint64) bool {
	var newProcess []*Process
	deleted := false

	for _, process := range processAdmin.ProcessAdmin {
		if process.Id != processId {
			newProcess = append(newProcess, process)
		}

		if process.Id == processId {
			deleted = true
			process.Stop()
			processAdmin.ProcessLength -= 1
		}
	}

	processAdmin.ProcessAdmin = newProcess
	return deleted
}

// ShowProcess function.
func (processAdmin *ProcessAdmin) ShowProcess() {
	for _, process := range processAdmin.ProcessAdmin {
		fmt.Print("\nId ", process.Id, " : ", process.Task)
	}
	fmt.Println()
}

// invalidOptions function.
func invalidOptions() {
	fmt.Print("\n-Invalid Option!\n\n")
}

// exited function.
func (processAdmin *ProcessAdmin) exited() {
	for _, process := range processAdmin.ProcessAdmin {
		process.Stop()
	}
	fmt.Println("\n-System exited...")
}
