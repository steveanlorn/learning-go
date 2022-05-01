package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// Imagine we have an app that logging to a device.
// Someday we detect that the device is full and causing our app become deadlock (blocked by the full disk).
// How we can solve with the concurrency pattern?

// device notes device mocker where we write logs to.
type device struct {
	problem bool
}

// Write implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	// g denotes number of goroutines that will be writing logs.
	const g = 10

	// d denotes a device where the log goes to.
	var d device
	l := log.New(&d, "prefix", 0)

	// Generates goroutines,  each writing to disk.
	for i := 0; i < g; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(20 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking.
	// Capture interrupt signals to toggle device issues.
	// Use <ctrl> z to kill the program.

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan
		d.problem = !d.problem
	}
}
