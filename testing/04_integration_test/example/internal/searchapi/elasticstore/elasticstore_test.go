package elasticstore

import (
	"bytes"
	"context"
	"flag"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/olivere/elastic/v7"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi/service"
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

func Test_constructMusicQuery(t *testing.T) {
	title := "butter"

	query := constructMusicQuery(title)

	{
		querySource := getQuerySource(t, query)
		want := goldenValue(t, querySource, *update)

		t.Log("Given the need to generate Elasticsearch music query")
		if !bytes.Equal(want, querySource) {
			t.Fatalf("\t%s\tShould get the same query:\n%v", failed, string(querySource))
		}
		t.Logf("\t%s\tShould get the same query", succeed)
	}
}

func TestStore_GetMusicByTitle(t *testing.T) {
	testcases := []struct {
		label      string
		title      string
		payload    string
		statusCode int
		wantMusics []service.Music
		wantErr    bool
	}{
		{
			label: "matchedWithASong",
			title: "love",
			payload: ` 
				{
				  "took": 4,
				  "timed_out": false,
				  "_shards": {
					"total": 1,
					"successful": 1,
					"skipped": 0,
					"failed": 0
				  },
				  "hits": {
					"total": 1,
					"max_score": 1,
					"hits": [
					  {
						"_index": "music",
						"_id": "1",
						"_score": 1,
						"_routing": "1",
						"_source": {
							"id": "1",
							"title": "How Deep Is Your Love"
						}
					  }
					]
				  }
				}
			`,
			statusCode: 200,
			wantMusics: []service.Music{
				{
					ID:    1,
					Title: "How Deep Is Your Love",
				},
			},
			wantErr: false,
		},
	}

	t.Logf("Given the need to test GetMusicByTitle")
	for _, tc := range testcases {
		tf := func(t *testing.T) {
			server := mockServer(t, tc.payload, tc.statusCode)
			defer func() {
				server.Close()
			}()

			esClient, err := elastic.NewSimpleClient(elastic.SetURL(server.URL))
			if err != nil {
				t.Fatalf("\t%s\tShould able to create ES simple client: %v", failed, err)
			}
			t.Logf("\t%s\tShould able to create ES simple client", succeed)

			store := NewStore(esClient)
			result, err := store.GetMusicByTitle(context.Background(), tc.title)
			if (err != nil) != tc.wantErr {
				t.Fatalf("\t%s\tShould able to get music by title: %v", failed, err)
			}
			t.Logf("\t%s\tShould able to get music by title", succeed)

			if diff := cmp.Diff(result, tc.wantMusics); diff != "" {
				t.Logf("\t%s\tShould get expected music", failed)
				t.Fatal(diff)
			}
			t.Logf("\t%s\tShould get expected music", succeed)
		}

		t.Run(tc.label, tf)
	}
}
