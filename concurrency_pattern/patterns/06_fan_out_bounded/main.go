package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// tasks denotes piece of work to do
	tasks := []string{"sales1", "sales2", "sales3", "sales4", 100: "sales99"}

	// number of worker goroutines that will process the task
	g := runtime.NumCPU()

	// wg denotes waitGroup to orchestrate the work
	var wg sync.WaitGroup
	wg.Add(g)

	// channel denotes a buffered channel of type string which provides signaling semantics
	channel := make(chan string, g)

	// create and launch worker goroutines
	for e := 0; e < g; e++ {
		go func(employee int) {
			defer wg.Done()

			// for-range loop used to check for new task on communication channel `channel`
			for t := range channel {
				fmt.Printf("employee %d: received signal: %s\n", employee, t)
			}

			// printed when communication channel is closed
			fmt.Printf("employee %d: received shutdown\n", employee)
		}(e)
	}

	// range over collection of task, one value at the time
	for _, task := range tasks {
		// signal/send task into channel
		// if buffer is full, this operation blocks
		channel <- task
	}

	// once last piece of task is submitted, close the channel
	// worker goroutines will process everything from the buffer
	close(channel)

	// guarantee point, wait for all worker goroutines to finish the work
	wg.Wait()
}
