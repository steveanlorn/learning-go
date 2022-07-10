package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const succeed = "\u2713"
const failed = "\u2717"

type Container struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource

	Elasticsearch Elasticsearch
}

type Elasticsearch struct {
	url    string
	Client *elastic.Client
}

// NewContainer starts and run Elasticsearch.
func NewContainer(ctx context.Context, t *testing.T) (*Container, error) {
	t.Helper()
	var container Container

	// =========================================================================
	// Container connection pool
	t.Log("\tCreating Docker's connection pool")
	{
		pool, err := dockertest.NewPool("")
		if err != nil {
			return nil, err
		}

		container.pool = pool
	}
	t.Logf("\t%s\tConnection pool is created", succeed)

	// =========================================================================
	// Elasticsearch container
	t.Log("\tStarting Elasticsearch Docker container")
	{
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

		resource, err := container.pool.RunWithOptions(&options)
		if err != nil {
			return nil, err
		}

		container.resource = resource
	}
	t.Logf("\t%s\tDocker container is created", succeed)

	// =========================================================================
	// Elasticsearch client
	t.Log("\tWaiting for Elasticsearch cluster to boot up")
	{
		container.Elasticsearch.url = fmt.Sprintf("http://127.0.0.1:%s", container.resource.GetPort("9200/tcp"))

		f := func() error {
			clientOptions := []elastic.ClientOptionFunc{
				elastic.SetURL(container.Elasticsearch.url),
				elastic.SetSniff(false),
				elastic.SetHealthcheck(false),
			}

			client, err := elastic.NewClient(clientOptions...)
			if err != nil {
				return err
			}

			_, _, err = client.Ping(container.Elasticsearch.url).Do(ctx)
			if err != nil {
				return err
			}

			container.Elasticsearch.Client = client
			return nil
		}

		if err := container.pool.Retry(f); err != nil {
			_ = container.Shutdown(t)
			return nil, err
		}
	}
	t.Logf("\t%s\tConnected to Elasticsearch cluster", succeed)

	return &container, nil
}

func (c *Container) CreateIndex(ctx context.Context, t *testing.T, indexName string, indexFilePath string) error {
	t.Helper()
	t.Log("\tCreating Elasticsearch index")

	file, err := os.Open(indexFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	jsonMap := make(map[string]interface{})
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&jsonMap); err != nil {
		return err
	}

	service := c.Elasticsearch.Client.CreateIndex(indexName)
	service.BodyJson(jsonMap)
	_, err = service.Do(ctx)
	if err != nil {
		return err
	}

	t.Logf("\t%s\tIndex `%s` is created", succeed, indexName)
	return nil
}

func (c *Container) DumpData(ctx context.Context, t *testing.T, dataFilePath string) error {
	t.Helper()
	t.Log("\tDumping the data")

	file, err := os.Open(dataFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	bulkAPI := c.Elasticsearch.url + "/_bulk?refresh=true"
	req, err := http.NewRequest(http.MethodPost, bulkAPI, file)
	if err != nil {
		return err
	}

	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status code %d: %s", resp.StatusCode, string(body))
	}

	t.Logf("\t%s\tData is dumped", succeed)
	return nil
}

func (c *Container) Shutdown(t *testing.T) error {
	t.Helper()
	t.Log("\tShutting down the container")
	if err := c.resource.Close(); err != nil {
		return err
	}
	t.Logf("\t%s\tContainer and volume is removed", succeed)

	return nil
}
