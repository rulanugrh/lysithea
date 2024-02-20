package service

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type InterfaceElasticsearch interface {
	GetProductBySearch(page int, perPage int, search string, buf bytes.Buffer) (map[string]interface{}, error)
}

type Elasticsearch struct {
	es *elasticsearch.Client
}

func NewElasticSearch(es *elasticsearch.Client) InterfaceElasticsearch {
	return &Elasticsearch{
		es: es,
	}
}

func (e *Elasticsearch) GetProductBySearch(page int, perPage int, search string, buf bytes.Buffer) (map[string]interface{}, error) {
	query := map[string]interface{}{
		"from":  page,
		"limit": perPage,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"name": map[string]interface{}{
								"query":         search,
								"fuzziness":     "AUTO",
								"prefix_length": 1,
							},
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"expire_at": map[string]interface{}{
								"query":         search,
								"fuzziness":     "AUTO",
								"prefix_length": 1,
							},
						},
					},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(&query); err != nil {
		log.Fatalf("Encoding error, %v", err.Error())
	}

	var resposeBind map[string]interface{}
	res, err := e.es.Search(
		e.es.Search.WithBody(&buf),
		e.es.Search.WithPretty(),
		e.es.Search.WithIndex("products"),
		e.es.Search.WithTrackTotalHits(false),
		e.es.Search.WithFilterPath("hits.hits._source.name",
			"hits.hits._source.expire_at", "hits.hits._source.id"),
	)

	if err != nil {
		log.Fatalf("Error getting response, %v", err.Error())
	}

	if err = json.NewDecoder(res.Body).Decode(&resposeBind); err != nil {
		log.Fatalf("Error Decoder, %v", err.Error())
	}

	return resposeBind, nil

}
