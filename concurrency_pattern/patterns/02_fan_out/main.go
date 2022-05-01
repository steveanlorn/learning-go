package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	employees := 10

	// channel provides signaling semantics.
	// buffered channel is used so no goroutine blocks a sending operation
	// if two goroutines send a signal at the same time, channel performs synchronization
	channel := make(chan string, employees)

	for e := 0; e < employees; e++ {

		// start goroutine that does some task for employee e
		go func(employee int) {
			// simulating the time need to complete the task.
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			// when task is done send signal
			channel <- "sales"
			fmt.Println("employee: sent signal: ", employee)
		}(e)
	}

	// wait for all employee task to be done
	for employees > 0 {
		// receive signal sent from the employee
		t := <-channel
		employees--
		fmt.Println("manager: received signal: ", t)
	}

	// after receiving all the task, then move on

	time.Sleep(time.Second)
}
