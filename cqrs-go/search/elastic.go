package search

import (
	"context"
	"encoding/json"
	"github.com/obarbier/cqrs-go/schema"
	"github.com/olivere/elastic"
	"log"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	_, err := r.client.Index().
		Index("meows").
		Type("meow").
		Id(meow.ID).
		BodyJson(meow.Body).
		Refresh("wait_for").
		Do(ctx)
	//params := url.Values{}
	//params.Set("pretty", fmt.Sprint(true))
	//path, err := uritemplates.Expand("/{index}", map[string]string{
	//	"index": "meows",
	//})
	//if err != nil {
	//	return err
	//}
	//
	//_, err = r.client.PerformRequest(ctx, elastic.PerformRequestOptions{
	//	Method: "PUT",
	//	Path:   path,
	//	Params: params,
	//	Body:   meow,
	//})
	return err
}

func (r *ElasticRepository) SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	result, err := r.client.Search().
		Index("meows").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	meows := []schema.Meow{}
	for _, hit := range result.Hits.Hits {
		var meow schema.Meow
		if err = json.Unmarshal(*hit.Source, &meow); err != nil {
			log.Println(err)
		}
		meows = append(meows, meow)
	}
	return meows, nil
}
