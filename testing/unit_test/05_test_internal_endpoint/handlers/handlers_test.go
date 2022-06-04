package handlers

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestGetUser(t *testing.T) {
	testCases := []struct {
		label              string
		url                string
		method             string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			label:              "success",
			url:                "/user",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"id":1,"username":"ag","firstname":"Andrew","lastname":"Garfield"},{"id":2,"username":"th","firstname":"Tom","lastname":"Hiddleston"}]`,
		},
	}

	for _, tc := range testCases {
		tf := func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.url, nil)
			w := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(w, r)

			if tc.expectedStatusCode != w.Code {
				t.Errorf("\t%s\tShould get status code %d, but got %d", failed, tc.expectedStatusCode, w.Code)
			} else {
				t.Logf("\t%s\tShould get status code %d", succeed, tc.expectedStatusCode)
			}

			if diff := cmp.Diff(tc.expectedResponse, w.Body.String()); diff != "" {
				t.Errorf("\t%s\tShould get response code %s\n\t%s", failed, tc.expectedResponse, diff)
			} else {
				t.Logf("\t%s\tShould get the same response code", succeed)
			}
		}
		t.Run(tc.label, tf)
	}
}
