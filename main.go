package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		opc string
		id int64 = -1
		condition = true
		task = make(chan Process)
	)

	for exit := true; exit; exit = condition {
		fmt.Println("******* Process Manager *******")
		fmt.Println("[1] Add Process")
		fmt.Println("[2] Show Process")
		fmt.Println("[3] Delete Process")
		fmt.Println("[0] Exit System")
		fmt.Print("-Option: ")
		fmt.Scan(&opc)

		switch opc {
			case "1":
				id = id + 1
				go addProcess(uint64(id), task)
				fmt.Print("\n-Add process #", uint64(id), "\n\n")
				go showProcess(task, false)
				break
			case "2":
				showProcess(task, true)
				break
			case "3":
				//go deleteProcess(1, task)
				break
			case "0":
				exited()
				condition = false
				break
			default:
				invalidOptions()
		}
	}
}

// Process struct
type Process struct {
	Id uint64
	Task uint64
}

// addProcess function
func addProcess(id uint64, task chan Process) {
	i := uint64(0)
	for {
		i = i + 1
		task <- Process{Id: id, Task: i}
		time.Sleep(time.Millisecond * 500)
	}
}

// showProcess function
func showProcess(task chan Process, show bool) {
	for {
		data := Process{}
		data = <- task
		if show {
			fmt.Print("\nId ", data.Id, " : ", data.Task)
		}
	}
}

//// stopProcess function
//func stopProcess(task chan Process) {
//	var stop string
//	fmt.Scanln(&stop)
//	if stop == "s" {
//		fmt.Println("msj",stop)
//		return
//	}
//}

//// deleteProcess function
//func deleteProcess(id uint64, task chan Process) {
//	data := <- task
//	if id == data.Id {
//		fmt.Println("Hello")
//		return
//	}
//}

// invalidOptions function.
func invalidOptions() {
	fmt.Print("\n-Invalid Option!\n\n")
}

// exited function.
func exited() {
	fmt.Println("\n-System exited...")
}

//package main
//
//import (
//"fmt"
//"time"
//)
//
//type Process struct {
//	Id int
//	Task int
//}
//
//func addProcess(id int, task chan Process) {
//	i := 0
//	for {
//		fmt.Println("Id", id, ":", i)
//		i = i + 1
//		task <- Process{Id: id, Task: i}
//		time.Sleep(time.Millisecond * 500)
//	}
//}
//
//func showProcess(task chan Process) {
//	for {
//		<-task
//	}
//}
//
//func stopProcess() {
//	var input string
//	fmt.Scan(&input)
//}
//
//func deleteProcess(id int, task chan Process) {
//	data := <- task
//	if id == data.Id {
//		fmt.Println("Hello")
//		close(task)
//	}
//}
//
//func main() {
//	task := make(chan Process)
//	stop := 1
//
//	for p := 0; p <= stop; p++ {
//		go addProcess(p, task)
//	}
//
//	go showProcess(task)
//	go deleteProcess(1, task)
//	stopProcess()
//}

//func main() {
//	var (
//		opc string
//		id int64 = -1
//		show = true
//		condition = true
//		task = make(chan Process)
//	)
//
//	for exit := true; exit; exit = condition {
//		fmt.Println("******* Process Manager *******")
//		fmt.Println("[1] Add Process")
//		fmt.Println("[2] Show Process")
//		fmt.Println("[3] Delete Process")
//		fmt.Println("[0] Exit System")
//		fmt.Print("-Option: ")
//		fmt.Scan(&opc)
//
//		if opc == "1" {
//			id = id + 1
//			go addProcess(uint64(id), task)
//			fmt.Print("\n-Add process #", uint64(id), "\n\n")
//			go showProcess(task, false)
//		} else if opc == "2" {
//			if show {
//				go showProcess(task, show)
//			} else {
//				go showProcess(task, show)
//			}
//			show = false
//		} else if opc == "3" {
//			//	go deleteProcess(1, task)
//		} else if opc == "0" {
//			exited()
//			condition = false
//		} else {
//			invalidOptions()
//		}
//	}
//}
