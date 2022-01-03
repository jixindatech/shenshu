package models

import (
	"admin/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v7"
	"io"
)

var esClient *es.Client
var esCfg *config.Elasticsearch

func SetupES(cfg *config.Elasticsearch) error {
	esCfg = cfg
	es7Config := es.Config{
		Addresses: []string{esCfg.Host},
		Username:  esCfg.User,
		Password:  esCfg.Password,
		CloudID:   "",
		APIKey:    "",
	}

	var err error
	esClient, err = es.NewClient(es7Config)
	if err != nil {
		return err
	}

	res, err := esClient.Info()
	if err != nil {
		return fmt.Errorf("es error: %s\n", err)
	}
	defer res.Body.Close()

	return nil
}

func esSearchList(index string, query map[string]interface{}, page, pageSize int) (map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithFrom((page-1)*pageSize),
		esClient.Search.WithSize(pageSize),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			fmt.Println(e)
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}

	return r, nil
}
func esSearch(index string, query map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	fmt.Println(res)

	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}

	return r, nil
}

func GetCCEventList(query map[string]interface{}, page, pageSize int) (map[string]interface{}, error) {
	data, err := esSearchList(esCfg.CCIndex, query, page, pageSize)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	res["count"] = data["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	res["data"] = data["hits"].(map[string]interface{})["hits"]
	return res, nil
}

func GetBatchRuleEventList(query map[string]interface{}, page, pageSize int) (map[string]interface{}, error) {
	data, err := esSearchList(esCfg.BatchRuleIndex, query, page, pageSize)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	res["count"] = data["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	res["data"] = data["hits"].(map[string]interface{})["hits"]
	return res, nil
}

func GetSpecificRuleEventList(query map[string]interface{}, page, pageSize int) (map[string]interface{}, error) {
	data, err := esSearchList(esCfg.SpecificRuleIndex, query, page, pageSize)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	res["count"] = data["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	res["data"] = data["hits"].(map[string]interface{})["hits"]
	return res, nil
}

func GetEventInfo(index string, query map[string]interface{}) (map[string]interface{}, error) {
	var esIndex string
	switch index {
	case "cc":
		esIndex = esCfg.CCIndex
	case "batch":
		esIndex = esCfg.BatchRuleIndex
	case "specific":
		esIndex = esCfg.SpecificRuleIndex
	}

	data, err := esSearch(esIndex, query)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	res["count"] = data["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	// res["data"] = data["hits"].(map[string]interface{})["hits"]
	res["aggregations"] = data["aggregations"]
	return res, nil
}
