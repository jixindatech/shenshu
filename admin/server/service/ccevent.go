package service

import "admin/server/models"

type CCEvent struct {
	ID uint

	SiteID uint
	Start  int64
	End    int64

	Page     int
	PageSize int
}

func (c *CCEvent) GetList() (map[string]interface{}, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": c.Start / 1000,
					"lte": c.End / 1000,
				},
			},
		},
	}

	return models.GetCCEventList(query, c.Page, c.PageSize)
}

func (c *CCEvent) GetInfo() (map[string]interface{}, error) {
	start := c.Start / 1000
	end := c.End / 1000
	interval := (end - start) / maxItemGroup

	filter := []map[string]interface{}{
		{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": start,
					"lte": end,
				},
			},
		},
	}
	if c.SiteID > 0 {
		filter = append(filter, map[string]interface{}{
			"terms": map[string]interface{}{
				"router": []uint{c.SiteID},
			},
		})
	}
	query := map[string]interface{}{
		"size": 0,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
			},
		},
		"aggs": map[string]interface{}{
			"by_timestamp": map[string]interface{}{
				"histogram": map[string]interface{}{
					"field":    "timestamp",
					"interval": interval,
					"extended_bounds": map[string]interface{}{
						"min": start,
						"max": end,
					},
				},
			},
		},
	}

	res, err := models.GetEventInfo("cc", query)
	if err != nil {
		return nil, err
	}

	docCount := res["aggregations"].(map[string]interface{})["by_timestamp"].(map[string]interface{})["buckets"]
	var intervalData []int64
	for _, item := range docCount.([]interface{}) {
		intervalData = append(intervalData, int64(item.(map[string]interface{})["doc_count"].(float64)))
	}
	res["interval"] = intervalData
	delete(res, "aggregations")

	return res, nil
}
