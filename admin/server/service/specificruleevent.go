package service

import "admin/server/models"

type SpecificRuleEvent struct {
	ID uint

	Start int64
	End   int64

	Page     int
	PageSize int
}

func (c *SpecificRuleEvent) GetList() (map[string]interface{}, error) {
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

	return models.GetSpecificRuleEventList(query, c.Page, c.PageSize, c.Start, c.End)
}
