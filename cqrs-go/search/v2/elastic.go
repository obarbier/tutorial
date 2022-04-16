package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/obarbier/cqrs-go/schema"
	"github.com/obarbier/cqrs-go/search"
)

type ElasticRepository struct {
	client *elasticsearch.Client
}

var _ search.Repository = &ElasticRepository{}

func NewElastic(url string) (*ElasticRepository, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			url,
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (e *ElasticRepository) Close() {
	//TODO implement me
	panic("implement me")
}

func (e *ElasticRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	b, err := json.Marshal(meow)
	res, err := e.client.Index(
		"meows",                            // Index name
		bytes.NewReader(b),                 // Document body 		// Document ID
		e.client.Index.WithRefresh("true"), // Refresh
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (e *ElasticRepository) SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	//TODO implement me
	panic("implement me")
}
