package es

import (
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func ConnectElasticsearch() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	return es
}
