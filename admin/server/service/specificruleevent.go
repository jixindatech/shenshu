package service

import "admin/server/models"

type SpecificRuleEvent struct {
	ID uint

	SiteID uint
	Start  int64
	End    int64

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
		res, err := models.GetSpecificRuleEventInfo(query)
		if err != nil {
			return nil, err
		}
		infos[item.Name] = res
	}

	return infos, nil
}
