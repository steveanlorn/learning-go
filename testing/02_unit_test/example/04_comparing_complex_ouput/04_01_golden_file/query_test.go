package _4_01_golden_file

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/olivere/elastic/v7"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestGenerateMusicQuery(t *testing.T) {
	testCases := []struct {
		label   string
		keyword string
	}{
		{
			label:   "godFather",
			keyword: "god father",
		}, {
			label:   "sweetChildOMine",
			keyword: "sweet child o mine",
		},
	}

	t.Log("Given the need to generate Elasticsearch music query")
	for i, tc := range testCases {
		tf := func(t *testing.T) {
			t.Logf("\tTest %d:\t%s\n", i, tc.label)

			queryResult := GenerateMusicQuery(tc.keyword)
			querySource := getQuerySource(t, queryResult)

			want := goldenValue(t, tc.label, querySource, *update)
			if !bytes.Equal(want, querySource) {
				t.Fatalf("\t%s\tShould get the same query:\n%v", failed, string(querySource))
			}
			t.Logf("\t%s\tShould get the same query", succeed)
		}
		t.Run(tc.label, tf)
	}
}

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

func goldenValue(t *testing.T, goldenFile string, actual []byte, update bool) []byte {
	t.Helper()
	goldenPath := filepath.Join("testdata", goldenFile+".golden")

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

func TestCompareDiff(t *testing.T) {
	a := &http.Server{Addr: ":7070"}
	b := &http.Server{Addr: ":7071"}

	t.Log("With 'reflect.DeepEqual'")
	if !reflect.DeepEqual(a, b) {
		t.Errorf("\t%s\t'%+v' is not equal with '%+v'", failed, a, b)
	} else {
		t.Logf("\t%s\t'%+v' is equal with '%+v'", succeed, a, b)
	}

	t.Log("With 'cmp'")
	if diff := cmp.Diff(a, b, cmpopts.IgnoreUnexported(http.Server{})); diff != "" {
		t.Logf("\t%s\ta & b are not equal:", failed)
		t.Errorf(diff)
	} else {
		t.Logf("\t%s\ta & b are equal:", succeed)
	}
}
