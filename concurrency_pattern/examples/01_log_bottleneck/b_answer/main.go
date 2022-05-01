package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	log "github.com/steveanlorn/learning-go/concurrency_pattern/examples/01_log_bottleneck/b_answer/logger"
)

// ANSWER:
// 1. Create a custom logger to handle the potential issue.
// 2. Use one goroutine to write to the device to easier detection of the problem.
// 3. Use signaling and a buffered channel to deal with the other goroutine that try to write to the device.
// 4. The buffered channel acts as a capacity control.
//		We set the capacity in a way that it can hold the normal volume of the task with additional buffer size.
//      If the buffered channel is full, we assume that there is a problem with the device that makes the buffer full.
// 5. If the buffered channel is full of capacity, then drop the task.

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
	l := log.New(&d, g)

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
