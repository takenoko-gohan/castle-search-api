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
	Name        string   `json:"name"`
	Prefecture  string   `json:"prefecture"`
	Rulers      []string `json:"rulers"`
	Description string   `json:"description"`
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

	//r.Name = b["hits"].(map[string]interface{})["hits"].([]interface{})["_source"].(map[string]interface{})["name"].(string)
	//r.Prefecture = b["hits"].(map[string]interface{})["hits"].([]interface{})["_source"].(map[string]interface{})["prefecture"].(string)
	//r.Rulers = b["hits"].(map[string]interface{})["hits"].([]interface{})["_source"].(map[string]interface{})["rulers"].(string)
	//r.Description = b["hits"].(map[string]interface{})["hits"].([]interface{})["_source"].(map[string]interface{})["description"].(string)

	r := make([]Response, 3)

	for i, hit := range b["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		//i := strconv.Itoa(index)

		//r[i] = Response{
		//	Name:       doc["_source"].(map[string]interface{})["name"].(string),
		//	Prefecture: doc["_source"].(map[string]interface{})["prefecture"].(string),
		//	//Rulers:      doc["_source"].(map[string]interface{})["rulers"].([]string),
		//	Description: doc["_source"].(map[string]interface{})["description"].(string),
		//}
		fmt.Println("test1")
		r[i].Name = doc["_source"].(map[string]interface{})["name"].(string)
		fmt.Println("test2")
		r[i].Prefecture = doc["_source"].(map[string]interface{})["prefecture"].(string)
		fmt.Println("test3")
		for _, str := range doc["_source"].(map[string]interface{})["rulers"].([]interface{}) {
			r[i].Rulers = append(r[i].Rulers, str.(string))
		}
		fmt.Println("test4")
		r[i].Description = doc["_source"].(map[string]interface{})["description"].(string)
		fmt.Println(r[i])
	}

	//fmt.Println(r)

	return c.JSON(http.StatusOK, r)
}
