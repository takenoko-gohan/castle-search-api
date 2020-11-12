package search

func createQuery(q *Query) map[string]interface{} {
	query := map[string]interface{}{}
	if q.Keyword != "" && q.Prefecture != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"bool": map[string]interface{}{
								"should": []map[string]interface{}{
									{
										"match": map[string]interface{}{
											"name": map[string]interface{}{
												"query": q.Keyword,
												"boost": 3,
											},
										},
									},
									{
										"match": map[string]interface{}{
											"rulers": map[string]interface{}{
												"query": q.Keyword,
												"boost": 2,
											},
										},
									},
									{
										"match": map[string]interface{}{
											"description": map[string]interface{}{
												"query": q.Keyword,
												"boost": 1,
											},
										},
									},
								},
								"minimum_should_match": 1,
							},
						},
						{
							"bool": map[string]interface{}{
								"must": []map[string]interface{}{
									{
										"term": map[string]interface{}{
											"prefecture": q.Prefecture,
										},
									},
								},
							},
						},
					},
				},
			},
		}
	} else if q.Keyword != "" && q.Keyword != "城" && q.Prefecture == "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": []map[string]interface{}{
						{
							"match": map[string]interface{}{
								"name": map[string]interface{}{
									"query": q.Keyword,
									"boost": 3,
								},
							},
						},
						{
							"match": map[string]interface{}{
								"rulers": map[string]interface{}{
									"query": q.Keyword,
									"boost": 2,
								},
							},
						},
						{
							"match": map[string]interface{}{
								"description": map[string]interface{}{
									"query": q.Keyword,
									"boost": 1,
								},
							},
						},
					},
					"minimum_should_match": 1,
				},
			},
		}
	} else if q.Keyword == "" && q.Prefecture != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"prefecture": q.Prefecture,
							},
						},
					},
				},
			},
		}
	}

	return query
}
