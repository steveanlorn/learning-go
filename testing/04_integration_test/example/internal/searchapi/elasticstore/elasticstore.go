package elasticstore

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi/service"
)

// Store ...
type Store struct {
	esClient *elastic.Client
}

// Music ...
type Music struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// NewStore ...
func NewStore(esClient *elastic.Client) *Store {
	store := Store{
		esClient: esClient,
	}

	return &store
}

// GetMusicByTitle ...
func (s *Store) GetMusicByTitle(ctx context.Context, title string) ([]service.Music, error) {
	query := constructMusicQuery(title)
	response, err := s.esClient.Search("music").Query(query).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("s.esClient.Search(\"music\").Query(query).Do(ctx)"))
	}

	var music Music
	results := make([]service.Music, 0, response.TotalHits())
	for _, item := range response.Each(reflect.TypeOf(music)) {
		m := item.(Music)

		id, err := strconv.ParseInt(m.Id, 10, 64)
		if err != nil {
			log.Printf("failed to parse string '%s' to int64: %v\n", m.Id, err)
			continue
		}

		results = append(results, service.Music{
			ID:    id,
			Title: m.Title,
		})
	}

	return results, nil
}

func constructMusicQuery(title string) elastic.Query {
	query := elastic.NewTermQuery("title", title)
	return query
}
