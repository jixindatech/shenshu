package service

import "admin/server/models"

type SpecificRuleEvent struct {
	ID uint

	SiteID   uint
	SiteName string
	Start    int64
	End      int64

	Page     int
	PageSize int
}

func (c *SpecificRuleEvent) GetList() (map[string]interface{}, error) {
	start := c.Start / 1000
	end := c.End / 1000
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
			"term": map[string]interface{}{
				"router": c.SiteID,
			},
		})
	}
	if len(c.SiteName) > 0 {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{
				"host": c.SiteName,
			},
		})
	}
	query := map[string]interface{}{
		"sort": map[string]interface{}{
			"timestamp": map[string]interface{}{
				"order": "desc",
			},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
			},
		},
	}

	return models.GetSpecificRuleEventList(query, c.Page, c.PageSize)
}

func (c *SpecificRuleEvent) GetInfo() (map[string]interface{}, error) {
	var err error
	var list []*models.SpecificGroup
	if c.SiteID == 0 {
		group := &SpecificGroup{
			Page:     0,
			PageSize: 0,
		}
		list, _, err = group.GetList()
	} else {
		list, err = models.GetSiteSpecificGroup(c.SiteID)
	}
	if err != nil {
		return nil, err
	}

	start := c.Start / 1000
	end := c.End / 1000
	interval := (end - start) / maxItemGroup

	infos := make(map[string]interface{})
	for _, item := range list {
		ruleSrv := &RuleSpecific{
			RuleGroup: item.ID,
			Page:      0,
			PageSize:  0,
		}
		rules, _, err := ruleSrv.GetList()
		if err != nil {
			return nil, err
		}
		ids := []uint{}
		for _, rule := range rules {
			ids = append(ids, rule.ID)
		}

		filter := []map[string]interface{}{
			{
				"terms": map[string]interface{}{
					"id": ids,
				},
			},
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
				"term": map[string]interface{}{
					"router": c.SiteID,
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
		res, err := models.GetEventInfo("specific", query)
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

		infos[item.Name] = res
	}

	return infos, nil
}
