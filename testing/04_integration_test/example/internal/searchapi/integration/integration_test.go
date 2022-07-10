package integration

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/olivere/elastic/v7"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi/elasticstore"
	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi/httphandler"
	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi/service"
)

// To run integration test:
// only integration test should have test name contains Integration
//   go test ./... -run Integration  -v
//
// To run unit test only:
//   go test ./... -v -short
//
// Alternate way is by using built tag
// // +build integration
// Run the test by calling
//   go test ./... -v -tags integration

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration test is skipped")
	}

	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		testDuration := deadline.Sub(time.Now())

		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), testDuration)
		defer cancel()
	}

	t.Log("=========================================================================")
	t.Log("INITIALIZE DEPENDENCY")
	container, err := NewContainer(ctx, t)
	if err != nil {
		t.Fatalf("\t%s\t%v", failed, err)
		return
	}

	defer func() {
		t.Logf("=========================================================================")
		t.Logf("TEAR DOWN DEPENDENCY")
		if err = container.Shutdown(t); err != nil {
			t.Errorf("\t%s\t%v", failed, err)
		}
	}()

	indexFilePath := filepath.Join("testdata", "index_music.json")
	if err = container.CreateIndex(ctx, t, "music", indexFilePath); err != nil {
		t.Errorf("\t%s\t%v", failed, err)
		return
	}

	dataFilePath := filepath.Join("testdata", "data_music.json")
	if err = container.DumpData(ctx, t, dataFilePath); err != nil {
		t.Errorf("\t%s\t%v", failed, err)
		return
	}

	elasticsearchClient = container.Elasticsearch.Client

	t.Logf("=========================================================================")
	t.Logf("EXECUTING TEST")
	testSearchAPI(t)
}

var elasticsearchClient *elastic.Client

func testSearchAPI(t *testing.T) {
	testCases := []struct {
		label              string
		keyword            string
		expectedResult     string
		expectedStatusCode int
	}{
		{
			label:              "foundTheMusic",
			keyword:            "love",
			expectedStatusCode: http.StatusOK,
			expectedResult:     fmt.Sprintf("[{\"id\":1,\"title\":\"How deep is your love\"},{\"id\":2,\"title\":\"La la la love song\"}]\n"),
		},
		{
			label:              "musicNotFound",
			keyword:            "lope",
			expectedStatusCode: http.StatusOK,
			expectedResult:     fmt.Sprintf("[]\n"),
		},
	}

	t.Logf("Given the need to test search API handler")
	for _, tc := range testCases {
		tf := func(t *testing.T) {
			store := elasticstore.NewStore(elasticsearchClient)
			svc := service.NewService(store)
			handler := httphandler.NewHandler(svc)

			targetURL, err := url.Parse("/test")
			if err != nil {
				t.Errorf("\t%s\tFailed parsing the target URL: %v", failed, err)
			}

			q := targetURL.Query()
			q.Set("title", tc.keyword)
			targetURL.RawQuery = q.Encode()

			mux := http.NewServeMux()
			mux.Handle(targetURL.Path, handler)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, targetURL.String(), nil)
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
