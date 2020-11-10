package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/takenoko-gohan/castle-search-api/internal/es"
)

type Query struct {
	Keyword    string `json:"keyword" query:"keyword"`
	Prefecture string `json:"prefecture" query:"prefecture"`
}

type Response struct {
	Name        string
	Prefecture  string
	Rulers      string
	Description string
}

func CastleSearch(c echo.Context) (err error) {
	fmt.Println("検索API")

	q := new(Query)
	if err = c.Bind(q); err != nil {
		return
	}

	//r := new(Response)
	var (
		b   map[string]interface{}
		buf bytes.Buffer
	)
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"name": q.Keyword,
						},
					},
					{
						"match": map[string]interface{}{
							"rulers": q.Keyword,
						},
					},
					{
						"match": map[string]interface{}{
							"description": q.Keyword,
						},
					},
				},
				"minimum_should_match": 1,
			},
		},
	}

	if q.Prefecture != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = []map[string]interface{}{
			{
				"term": map[string]interface{}{
					"prefecture": q.Prefecture,
				},
			},
		}
	} else {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["minimum_should_match"] = 1
	}

	fmt.Println(query)

	json.NewEncoder(&buf).Encode(query)

	es := es.ConnectElasticsearch()
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("castle"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		log.Fatal(err)
	}

	var r interface{} = b["hits"].(map[string]interface{})["hits"].(map[string]interface{})["_source"].(map[string]interface{})

	return c.JSON(http.StatusOK, r)
}
