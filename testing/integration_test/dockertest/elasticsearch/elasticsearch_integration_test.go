package elasticsearch

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/steveanlorn/learning-go/testing/integration_test/dockertest/elasticsearch/elasticsearchtest"
)

const succeed = "\u2713"
const failed = "\u2717"
const loading = "\u29D6"

var elasticServer *elasticsearchtest.Server

func TestMain(m *testing.M) {
	var err error
	elasticServer, err = elasticsearchtest.NewServer()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	code := m.Run()

	err = elasticServer.Shutdown()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(code)
}

func TestQuery(t *testing.T) {
	resp, err := elasticClient.NodesInfo().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.ClusterName)
}
