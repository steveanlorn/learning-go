package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	if err := registerUser("admin", "password"); err != nil {

		if errors.Is(err, context.DeadlineExceeded) {
			// Got the custom error
			log.Println("Deadline exceeded error")
		} else {
			// Got the default error
			log.Println("Default error")
		}

		log.Println(err)
	}
}

func registerUser(username string, password string) error {
	if err := validateUser(username); err != nil {
		return err
	}

	// save username & password in db
	// .....

	return nil
}

func validateUser(username string) error {
	// string normalisation
	usernameToLower := strings.ToLower(username)

	c := config{
		timeout: 200 * time.Millisecond,
	}

	if err := connectToUserAPI(usernameToLower, c); err != nil {
		return err
	}

	return nil
}

type config struct {
	timeout time.Duration
}

// TIMEOUT ERROR EXAMPLE
func connectToUserAPI(username string, c config) error {
	ctx, cancel := context.WithTimeout(context.TODO(), c.timeout)
	defer cancel()

	var done chan bool
	url := fmt.Sprintf("https://myapi.com?user=%s", username)

	go func() {
		// simulating to call the API to check the username
		_, _ = http.Get(url)
		time.Sleep(2 * c.timeout)
		done <- true
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("error request:'%s'; err:%w; timeout is set to %s by config %#v", url, ctx.Err(), c.timeout.String(), c)
	}
}
