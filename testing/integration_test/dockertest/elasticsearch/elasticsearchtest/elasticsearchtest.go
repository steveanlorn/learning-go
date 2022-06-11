// Package elasticsearchtest provide utilities for Elasticsearch testing
package elasticsearchtest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pkg/errors"
)

// Server ...
type Server struct {
	resource            *dockertest.Resource
	elasticSearchClient *elastic.Client
}

// Shutdown ...
func (s *Server) Shutdown() error {
	s.GetClient().DeleteIndex("*").Do(context.Background())
	return s.resource.Close()
}

// GetClient ...
func (s *Server) GetClient() *elastic.Client {
	return s.elasticSearchClient
}

// CreateIndex helper functions to create index which settings and mappings are fetched from a file.
func (s *Server) CreateIndex(indexName string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("CreateIndex(%s, %s) -> os.Open(%s)", indexName, fileName, fileName))
	}

	defer file.Close()

	jsonMap := make(map[string]interface{})
	decoder := json.NewDecoder(file)
	decoder.Decode(&jsonMap)

	svc := s.GetClient().CreateIndex(indexName)
	svc.BodyJson(jsonMap)
	_, err = svc.Do(context.Background())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("CreateIndex(%s, %s) -> svc.Do(context.Background())", indexName, fileName))
	}

	return nil
}

// NewServer runs a new Elasticsearch server in a container
// and initializes Elasticsearch client connecting to that container.
// This is a long process especially when container does not exist in local file.
func NewServer() (*Server, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, errors.Wrap(err, "NewServer() -> dockertest.NewPool(\"\")")
	}

	options := dockertest.RunOptions{
		Repository: "docker.elastic.co/elasticsearch/elasticsearch-oss",
		Tag:        "7.8.0",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9200": {{HostPort: "9200"}},
		},
		Env: []string{
			"cluster.name=elasticsearch",
			"bootstrap.memory_lock=true",
			"discovery.type=single-node",
			"network.publish_host=127.0.0.1",
			"logger.org.elasticsearch=warn",
			"ES_JAVA_OPTS=-Xms1g -Xmx1g",
		},
	}

	resource, err := pool.RunWithOptions(&options)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("NewServer() -> pool.RunWithOptions(%v)", options))
	}

	server := Server{
		resource: resource,
	}

	// initialize Elasticsearch client
	{
		endpoint := fmt.Sprintf("http://127.0.0.1:%s", resource.GetPort("9200/tcp"))
		var elasticClient *elastic.Client
		if err := pool.Retry(func() error {
			var err error

			clientOptions := []elastic.ClientOptionFunc{
				elastic.SetURL(endpoint),
				elastic.SetSniff(false),
				elastic.SetHealthcheck(false),
			}
			elasticClient, err = elastic.NewClient(clientOptions...)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("NewServer() -> elastic.NewClient(%v)", clientOptions))
			}

			_, _, err = elasticClient.Ping(endpoint).Do(context.Background())
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("NewServer() -> elasticClient.Ping(%s)", endpoint))
			}

			return nil
		}); err != nil {
			_ = server.Shutdown()
			return nil, err
		}
		server.elasticSearchClient = elasticClient
	}

	return &server, nil
}
