package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// channel provides signaling semantics
	// unbuffered channel provides a guarantee that the
	// signal being sent is received
	channel := make(chan string)

	// a goroutine that waits for some task
	go func() {
		// employee waits for signal that it has some work to do
		t := <-channel
		fmt.Println("employee: received signal: ", t)
	}()

	// simulating the time need to prepare the task.
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// when the task is ready, send signal form manager to the employee
	// sender (employee) has a guarantee that the worker (employee)
	// has received a signal
	channel <- "sales"
	fmt.Println("manager: sent signal")

	time.Sleep(time.Second)
}
