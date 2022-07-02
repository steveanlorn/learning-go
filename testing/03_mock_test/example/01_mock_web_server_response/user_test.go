package user

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const succeed = "\u2713"
const failed = "\u2717"

func mockServer(t *testing.T, payload string, statusCode int) *httptest.Server {
	t.Helper()
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(statusCode)
		_, err := io.WriteString(w, payload)
		if err != nil {
			t.Fatalf("could not write to string from the mock server: %v", err)
		}
	}

	mockServer := httptest.NewServer(http.HandlerFunc(mockHandler))
	return mockServer
}

func TestLoadUserData(t *testing.T) {
	testCases := []struct {
		label        string
		payload      string
		statusCode   int
		expectedErr  bool
		expectedUser []User
	}{
		{
			label: "oneUser",
			payload: `
			{
			  "results": [
				{"name":{"first":"brad","last":"gibson"},"login":{"uuid":"155e","username":"silverswan131"}}
			  ]
			}`,
			statusCode:  200,
			expectedErr: false,
			expectedUser: []User{
				{
					ID:        "155e",
					Username:  "silverswan131",
					FirstName: "brad",
					LastName:  "gibson",
				},
			},
		},
		{
			label: "twoUsers",
			payload: `
			{
			  "results": [
				{"name":{"first":"brad","last":"gibson"},"login":{"uuid":"155e","username":"silverswan131"}},
				{"name":{"first":"بیتا","last":"کوتی"},"login":{"uuid":"4ef9","username":"tinykoala692"}}
			  ]
			}`,
			statusCode:  200,
			expectedErr: false,
			expectedUser: []User{
				{
					ID:        "155e",
					Username:  "silverswan131",
					FirstName: "brad",
					LastName:  "gibson",
				},
				{
					ID:        "4ef9",
					Username:  "tinykoala692",
					FirstName: "بیتا",
					LastName:  "کوتی",
				},
			},
		},
		{
			label:        "apiError",
			payload:      `{error:"Uh oh, something has gone wrong. Please tweet us @randomapi about the issue. Thank you."}`,
			statusCode:   500,
			expectedErr:  true,
			expectedUser: nil,
		},
	}

	t.Logf("Given the need to test LoadUserData")
	for _, tc := range testCases {
		tf := func(t *testing.T) {
			server := mockServer(t, tc.payload, tc.statusCode)
			defer server.Close()

			users, err := LoadUserData(server.URL)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, err)
			}
			t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)

			if diff := cmp.Diff(tc.expectedUser, users); diff != "" {
				t.Logf("\t%s\tShould get list of users", failed)
				t.Fatal(diff)
			}
			t.Logf("\t%s\tShould get list of users", succeed)
		}
		t.Run(tc.label, tf)

	}
}
