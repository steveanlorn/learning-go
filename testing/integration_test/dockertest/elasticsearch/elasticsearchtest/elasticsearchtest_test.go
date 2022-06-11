package elasticsearchtest

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestServer_CreateIndex(t *testing.T) {
	server, err := NewServer()
	if err != nil {
		t.Fatalf("%v", err)
	}

	defer server.Shutdown()

	err = server.CreateIndex("song", filepath.Join("testdata", "indexconfiguration", "index_1.json"))
	if err != nil {
		t.Fatalf("%v", err)
	}

	svc := server.GetClient().IndexExists("song")
	result, err := svc.Do(context.Background())
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !result {
		t.Fatal("index song is not exist")
	}

	resp, err := server.GetClient().Index().Index("song").Id("123").BodyString(`{"title": "ateez"}`).Refresh("true").Do(context.Background())
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(resp.Id)

	respSearch, err := server.GetClient().Search("song").Query(elastic.NewTermQuery("title", "ateez")).Do(context.Background())
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(respSearch.TotalHits())

	getResp, err := server.GetClient().Get().Index("song").Id("123").Do(context.Background())
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(getResp.Found)
}
