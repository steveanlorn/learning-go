package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	employees := 50

	// channel denotes a buffered channel, one slot for every goroutine
	// sender side can complete without receive (non-blocking)
	channel := make(chan string, employees)

	// g denotes max number of RUNNING goroutines at any given time
	g := runtime.NumCPU()

	// sem denotes a buffered channel, based on the max number of the goroutines in RUNNING state
	// added to CONTROL the number of goroutines in RUNNING state
	sem := make(chan bool, g)

	for e := 0; e < employees; e++ {
		// create 50 goroutines in the RUNNABLE state
		// one for each employee
		go func(employee int) {

			// when goroutine moves from RUNNABLE to RUNNING state
			// send signal/value inside a `sem` channel
			// if `sem` channel buffer is full, this will block
			sem <- true
			{
				// simulating the time need to complete the task.
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

				// once task is done, signal on ch channel
				channel <- "sales"
				fmt.Println("employee: sent signal: ", employee)
			}

			// once all task is done pull the value from the `sem` channel
			// give place to another goroutine to do the work
			<-sem
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
