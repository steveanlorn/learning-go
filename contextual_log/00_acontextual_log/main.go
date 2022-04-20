package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	c := config{
		timeout: 200 * time.Millisecond,
	}
	if err := connectToUserAPI(c, "admin"); err != nil {
		log.Println(err)
	}
}

type config struct {
	timeout time.Duration
}

// TIMEOUT ERROR EXAMPLE
func connectToUserAPI(c config, username string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), c.timeout)
	defer cancel()

	var done chan bool
	url := fmt.Sprintf("https://myapi.com?user=%s", username)

	go func() {
		// simulating to call the url
		_, _ = http.Get(url)
		time.Sleep(2 * c.timeout)
		done <- true
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		// How can we improve this log to have more context?
		return ctx.Err()
		//return fmt.Errorf("error request:'%s'; err:%w; timeout is set to %s by config %#v", url, ctx.Err(), c.timeout.String(), c)
	}
}
