package elasticstore

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/olivere/elastic/v7"
)

func getQuerySource(t *testing.T, query elastic.Query) []byte {
	t.Helper()
	sourceQuery, err := query.Source()
	if err != nil {
		t.Fatalf("Error getting source of the query: %v", err)
	}

	queryJSON, err := json.MarshalIndent(sourceQuery, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling query: %v", err)
	}

	return queryJSON
}

func goldenValue(t *testing.T, actual []byte, update bool) []byte {
	t.Helper()
	goldenPath := filepath.Join("testdata", "query.golden")

	if update {
		f, err := os.Create(goldenPath)
		defer f.Close()
		_, err = f.Write(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	f, err := os.Open(goldenPath)
	if err != nil {
		t.Fatalf("Error opening file %s: %v", goldenPath, err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}

	return content
}

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
