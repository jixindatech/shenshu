package service

import "admin/server/models"

type BatchRuleEvent struct {
	ID uint

	SiteID uint
	Start  int64
	End    int64

	Page     int
	PageSize int
}

func (c *BatchRuleEvent) GetList() (map[string]interface{}, error) {
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

	return models.GetBatchRuleEventList(query, c.Page, c.PageSize)
}

func (c *BatchRuleEvent) GetInfo() (map[string]interface{}, error) {
	var err error
	var list []*models.BatchGroup
	if c.SiteID == 0 {
		group := &BatchGroup{
			Page:     0,
			PageSize: 0,
		}
		list, _, err = group.GetList()
	} else {
		list, err = models.GetSiteBatchGroup(c.SiteID)
	}
	if err != nil {
		return nil, err
	}

	infos := make(map[string]interface{})
	for _, item := range list {
		ruleSrv := &RuleBatch{
			RuleGroup: item.ID,
			Page:      0,
			PageSize:  0,
		}
		rules, _, err := ruleSrv.GetList()
		if err != nil {
			return nil, err
		}
		var ids []uint
		for _, rule := range rules {
			ids = append(ids, rule.ID)
		}

		query := map[string]interface{}{
			"size": 0,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": []map[string]interface{}{
						{
							"terms": map[string]interface{}{
								"id": ids,
							},
						},
						{
							"range": map[string]interface{}{
								"timestamp": map[string]interface{}{
									"gte": c.Start / 1000,
									"lte": c.End / 1000,
								},
							},
						},
					},
				},
			},
		}

		res, err := models.GetBatchRuleEventInfo(query)
		if err != nil {
			return nil, err
		}
		infos[item.Name] = res
	}

	return infos, nil
}
