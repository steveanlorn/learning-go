package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// duration denotes waiting time for the manager
	duration := 200 * time.Millisecond

	// set a context with specified timeout based on the duration
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Buffered channel ensures that the worker goroutine can perform the send operation
	// and complete even if there is no-one on the reception side.
	channel := make(chan string, 1)

	go func() {
		// simulating the time need to complete the task.
		time.Sleep(time.Duration(rand.Intn(220)) * time.Millisecond)

		// send signal when gask is done
		channel <- "sales"
	}()

	// select-case allow us to perform multiple channel operations
	// at the same time, on the same goroutine
	select {
	// best case scenario:
	// receive a result from worker goroutine in under the 200 ms
	case done := <-channel:
		fmt.Println("work is complete ", done)
	// ctx.Done() call starts the 200ms duration clock ticking.
	// If 200 ms passes before the worker goroutine finishes, this println will be executed
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
