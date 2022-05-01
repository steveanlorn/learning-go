package main

import (
	"fmt"
	"time"
)

func main() {
	// cap denotes max number of active requests at any given moment
	const cap = 10

	// channel denotes buffered channel is used to determine when we are at capacity
	channel := make(chan string, cap)

	// a worker goroutine
	go func() {
		// range loop used to check for new work on communication channel `channel`
		for t := range channel {
			fmt.Println("employee: received signal: ", t)
		}
	}()

	// task denotes amount of work to do
	const task = 100

	// range over collection of task, one value at the time
	for t := 0; t < task; t++ {
		// select-case allow us to perform multiple channel operations
		// at the same time, on the same goroutine
		select {
		// sends task into channel
		// if buffer is full, default case is executed
		case channel <- "sales":
			fmt.Println("manager: sent signal: ", t)
		// if channel buffer is full, drop the message
		// allow us to detect that we are at capacity
		default:
			fmt.Println("manager: dropped data: ", t)
		}
	}

	// once last piece of work is submitted, close the channel
	// worker goroutines will process everything from the buffer
	close(channel)
	fmt.Println("manager: sent shutdown signal")
	time.Sleep(time.Second)
}
