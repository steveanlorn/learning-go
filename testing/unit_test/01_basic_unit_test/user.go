package user

import (
	"encoding/csv"
	"os"
	"strconv"
)

type User struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
}

// LoadUserData loads the user data from a CSV file and returns slice of User.
func LoadUserData(csvFile string) (users []User, err error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		errClose := file.Close()
		if err == nil {
			err = errClose
		}
	}()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	recordsWithoutHeader := records[1:]
	users = make([]User, len(recordsWithoutHeader))

	for i, record := range recordsWithoutHeader {
		userName := record[0]
		userID, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			return nil, err
		}

		firstName := record[2]
		lastName := record[3]

		user := User{
			ID:        userID,
			Username:  userName,
			FirstName: firstName,
			LastName:  lastName,
		}

		users[i] = user
	}

	return users, nil
}
