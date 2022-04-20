package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := registerUser("ADMIN", "password"); err != nil {
		switch errors.Cause(err) {
		case context.DeadlineExceeded:
			// Got the custom error
			log.Println("Deadline exceeded error")
		default:
			// Got the default error
			log.Println("Default error")

		}

		log.Printf("%v", err)
	}
}

func registerUser(username string, password string) error {
	if err := validateUser(username); err != nil {
		return errors.Wrapf(err, "registerUser(%s,%s)->validateUser(%s)", username, password, username)
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
		return errors.Wrapf(err, "validateUser(%s)->connectToUserAPI(%s)", username, usernameToLower)
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
		return errors.Wrapf(ctx.Err(), "error request:'%s'; timeout is set to %s by config %#v", url, c.timeout.String(), c)
	}
}
