package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID        string
	Username  string
	FirstName string
	LastName  string
}

type apiResponse struct {
	Results []struct {
		Login struct {
			Username string `json:"username"`
			UUID     string `json:"uuid"`
		} `json:"login"`
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
	} `json:"results"`
}

// LoadUserData loads the user data from an API and returns slice of User.
func LoadUserData(url string) (users []User, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status response %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	defer func() {
		errClose := resp.Body.Close()
		if err == nil {
			err = errClose
		}
	}()

	var ar apiResponse

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&ar); err != nil {
		return nil, err
	}

	users = make([]User, len(ar.Results))
	for i, result := range ar.Results {
		user := User{
			ID:        result.Login.UUID,
			Username:  result.Login.Username,
			FirstName: result.Name.First,
			LastName:  result.Name.Last,
		}

		users[i] = user
	}

	return users, nil
}
