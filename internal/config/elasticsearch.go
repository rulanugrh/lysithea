package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rulanugrh/lysithea/internal/entity/web"
)

func NewConnectionElastic() (*elasticsearch.Client, error) {
	conf := GetConfig()
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			conf.Elasticsearch.URL,
		},
		Username: conf.Elasticsearch.Username,
		Password: conf.Elasticsearch.Password,
	})

	if err != nil {
		return nil, web.StatusBadRequest("cannot connect to elasticsearch")
	}

	log.Println("Success connected to Elasticsearch")
	return client, nil
}
