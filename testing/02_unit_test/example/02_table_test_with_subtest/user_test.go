package user

import (
	"path/filepath"
	"reflect"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestLoadUserDataWithTable(t *testing.T) {
	testCases := []struct {
		label           string
		fileName        string
		expectedResults []User
		expectedErr     bool
	}{
		{
			label:    "withExistCSV",
			fileName: filepath.Join("testdata", "username.csv"),
			expectedResults: []User{
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
			},
			expectedErr: false,
		},
		{
			label:           "withNotExistCSV",
			fileName:        filepath.Join("testdata", "nonexistent.csv"),
			expectedResults: []User{},
			expectedErr:     true,
		},
	}

	t.Log("Given the need to load user data from a CSV file")
	for i, tc := range testCases {
		t.Logf("\tTest %d:\t%s\n", i, tc.label)
		{
			gotResult, gotErr := LoadUserData(tc.fileName)

			if (gotErr != nil) != tc.expectedErr {
				t.Errorf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, gotErr)
			} else {
				t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)
			}

			if !reflect.DeepEqual(gotResult, tc.expectedResults) {
				t.Errorf("\t%s\tShould get an equal []User result %v", failed, gotResult)
			} else {
				t.Logf("\t%s\tShould get an equal User result", succeed)
			}
		}
	}
}

func TestLoadUserDataWithTableMap(t *testing.T) {
	testCases := map[string]struct {
		fileName        string
		expectedResults []User
		expectedErr     bool
	}{
		"withExistCSV": {
			fileName: filepath.Join("testdata", "username.csv"),
			expectedResults: []User{
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
			},
			expectedErr: false,
		},
		"withNotExistCSV": {
			fileName:        filepath.Join("testdata", "nonexistent.csv"),
			expectedResults: []User{},
			expectedErr:     true,
		},
	}

	t.Log("Given the need to load user data from a CSV file")
	for name, tc := range testCases {
		t.Logf("\tTest:\t%s\n", name)
		{
			gotResult, gotErr := LoadUserData(tc.fileName)

			if (gotErr != nil) != tc.expectedErr {
				t.Errorf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, gotErr)
			} else {
				t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)
			}

			if !reflect.DeepEqual(gotResult, tc.expectedResults) {
				t.Errorf("\t%s\tShould get an equal []User result %v", failed, gotResult)
			} else {
				t.Logf("\t%s\tShould get an equal User result", succeed)
			}
		}
	}
}

func TestLoadUserDataWithTableSubTest(t *testing.T) {
	testCases := []struct {
		label           string
		fileName        string
		expectedResults []User
		expectedErr     bool
	}{
		{
			label:    "withExistCSV",
			fileName: filepath.Join("testdata", "username.csv"),
			expectedResults: []User{
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
			},
			expectedErr: false,
		},
		{
			label:           "withNotExistCSV",
			fileName:        filepath.Join("testdata", "nonexistent.csv"),
			expectedResults: []User{},
			expectedErr:     true,
		},
	}

	t.Log("Given the need to load user data from a CSV file")
	for i, tc := range testCases {
		tf := func(t *testing.T) {
			t.Logf("\tTest %d:\t%s\n", i, tc.label)
			{
				gotResult, gotErr := LoadUserData(tc.fileName)

				if (gotErr != nil) != tc.expectedErr {
					t.Errorf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, gotErr)
				} else {
					t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)
				}

				if !reflect.DeepEqual(gotResult, tc.expectedResults) {
					t.Errorf("\t%s\tShould get an equal []User result %v", failed, gotResult)
				} else {
					t.Logf("\t%s\tShould get an equal User result", succeed)
				}
			}
		}
		t.Run(tc.label, tf)
	}
}

func TestLoadUserDataWithTableSubTestParallel(t *testing.T) {
	testCases := []struct {
		label           string
		fileName        string
		expectedResults []User
		expectedErr     bool
	}{
		{
			label:    "withExistCSV",
			fileName: filepath.Join("testdata", "username.csv"),
			expectedResults: []User{
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
			},
			expectedErr: false,
		},
		{
			label:           "withNotExistCSV",
			fileName:        filepath.Join("testdata", "nonexistent.csv"),
			expectedResults: []User{},
			expectedErr:     true,
		},
	}

	t.Log("Given the need to load user data from a CSV file")
	for i := range testCases {
		tc := testCases[i]
		tf := func(t *testing.T) {
			t.Parallel()
			t.Logf("\tTest:\t%s\n", tc.label)
			{
				gotResult, gotErr := LoadUserData(tc.fileName)

				if (gotErr != nil) != tc.expectedErr {
					t.Fatalf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, gotErr)
				} else {
					t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)
				}

				if !reflect.DeepEqual(gotResult, tc.expectedResults) {
					t.Errorf("\t%s\tShould get an equal []User result %v", failed, gotResult)
				} else {
					t.Logf("\t%s\tShould get an equal User result", succeed)
				}
			}
		}
		t.Run(tc.label, tf)
	}
}
