package elasticsearch

import (
	"log"
	"os"
	"testing"

	"github.com/steveanlorn/learning-go/testing/integration_test/dockertest/elasticsearch/elasticsearchtest"
)

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
	if !testing.Short() {
		t.Log("TestQuery is skipped. To run this test, provide the -test.short")
	}
	// Integration test...
}
