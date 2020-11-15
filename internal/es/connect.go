package es

import (
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func ConnectElasticsearch() (*elasticsearch.Client, error) {
	var addr string
	if os.Getenv("ES_ADDRESS") != "" {
		addr = os.Getenv("ES_ADDRESS")
	} else {
		addr = "http://localhost:9200"
	}
	cfg := elasticsearch.Config{
		Addresses: []string{
			addr,
		},
	}
	es, err := elasticsearch.NewClient(cfg)

	return es, err
}
