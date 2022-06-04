package user

import (
	"path/filepath"
	"reflect"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestLoadUserData(t *testing.T) {
	// Arrange
	fileName := filepath.Join("testdata", "username.csv")
	expectedResults := []User{
		{
			ID:        9012,
			Username:  "booker12",
			FirstName: "Rachel",
			LastName:  "Booker",
		},
		{
			ID:        2070,
			Username:  "grey07",
			FirstName: "Laura",
			LastName:  "Grey",
		},
	}
	expectedErr := false

	// Act
	gotResult, gotErr := LoadUserData(fileName)

	// Assert
	t.Logf("Given the need to load user data from a '%s' CSV file\n", fileName)
	if (gotErr != nil) != expectedErr {
		t.Errorf("\t%s\tShould get an error is %v: %v", failed, expectedErr, gotErr)
	} else {
		t.Logf("\t%s\tShould get an error is %v", succeed, expectedErr)
	}

	if !reflect.DeepEqual(gotResult, expectedResults) {
		t.Errorf("\t%s\tShould get an equal []User result %v", failed, gotResult)
	} else {
		t.Logf("\t%s\tShould get an equal User result", succeed)
	}
}
