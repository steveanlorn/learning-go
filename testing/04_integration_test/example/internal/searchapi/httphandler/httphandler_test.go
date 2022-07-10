package httphandler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi"
	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi/mockservice"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestHandler_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := mockservice.NewMockService(ctrl)

	testCase := []struct {
		label              string
		setRequest         func() *http.Request
		setMock            func()
		expectedStatusCode int
		expectedResult     string
	}{
		{
			label: "success",
			setRequest: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/test?title=love", nil)
			},
			setMock: func() {
				ms.EXPECT().GetMusicByTitle(gomock.Any(), "love").Return(
					[]searchapi.Music{
						{
							ID:    1,
							Title: "How Deep Is Your Love",
						},
					}, nil)
			},
			expectedStatusCode: 200,
			expectedResult:     fmt.Sprintf("[{\"id\":1,\"title\":\"How Deep Is Your Love\"}]\n"),
		},
		{
			label: "notFound",
			setRequest: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/test?title=love", nil)
			},
			setMock: func() {
				ms.EXPECT().GetMusicByTitle(gomock.Any(), "love").Return(
					[]searchapi.Music{}, nil)
			},
			expectedStatusCode: 200,
			expectedResult:     fmt.Sprintf("[]\n"),
		},
	}

	t.Logf("Given the need to test Music Box HTTP Handler")
	for _, tc := range testCase {
		tf := func(t *testing.T) {
			tc.setMock()
			service := searchapi.Service(ms)
			handler := NewHandler(service)

			mux := http.NewServeMux()
			mux.Handle("/test", handler)

			r := tc.setRequest()
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, r)

			if tc.expectedStatusCode != w.Code {
				t.Errorf("\t%s\tShould get status code %d, but got %d", failed, tc.expectedStatusCode, w.Code)
			} else {
				t.Logf("\t%s\tShould get status code %d", succeed, tc.expectedStatusCode)
			}

			if diff := cmp.Diff(tc.expectedResult, w.Body.String()); diff != "" {
				t.Errorf("\t%s\tShould get response body %s\n\t%s", failed, tc.expectedResult, diff)
			} else {
				t.Logf("\t%s\tShould get the same response body", succeed)
			}
		}

		t.Run(tc.label, tf)
	}
}
