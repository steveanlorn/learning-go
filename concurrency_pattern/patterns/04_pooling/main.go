package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// channel provides signaling semantics
	// unbuffered channel provides a guarantee that the
	// signal being sent is received
	channel := make(chan string)

	// g denotes number of goroutines to create, numCPU() is a good starting point
	g := runtime.NumCPU()
	for e := 0; e < g; e++ {
		// a new goroutine is created for each employee
		go func(employee int) {
			// employee waits for the signal that there is some task to do
			// all goroutines are blocked on the same channel `channel` receive
			for t := range channel {
				fmt.Printf("employee %d: received signal: %s\n", employee, t)
			}

			// when all task is sent, manager notifies all employees by closing the channel
			// once the channel is closed, employee breaks out of the for-range loop
			fmt.Printf("employee %d : revieved shutdown signal\n", employee)
		}(e)
	}

	// amount of task to be done
	const task = 100
	for t := 0; t < task; t++ {
		// when task is ready, we send signal from the manager to the employee
		// sender (manager) has a guarantee that the worker (employee) has received the signal
		// manager doesn't care about which employee received a signal,
		// since all employees are capable of doing the task
		channel <- "sales"
		fmt.Println("manager: sent signal", t)
	}

	// when all task is sent the manager notifies all employees by closing the channel
	// unbuffered channel provides a guarantee that all work has been sent
	close(channel)
	fmt.Println("manager: sent shutdown signal")
	time.Sleep(time.Second)
}
