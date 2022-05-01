package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// channel provides signaling semantics
	channel := make(chan string)

	// a goroutine that does some work
	go func() {
		// simulating the time need to complete the task.
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		// when task is done send signal
		channel <- "sales"

		// we don't know which print statement is going to be executed first
		fmt.Println("employee: sent signal")
	}()

	// wait for and receive signal from goroutine.
	// blocking operation
	t := <-channel

	// we don't know which print statement is going to be executed first
	fmt.Println("manager: received signal: ", t)

	// ensure enough time to get the result (for demo only)
	time.Sleep(time.Second)
}
