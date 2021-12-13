package service

import (
	"admin/server/models"
)

type RuleGroup struct {
	ID uint

	Name   string
	Type   int
	Remark string

	Page     int
	PageSize int
}

func (r *RuleGroup) Save() error {
	data := map[string]interface{}{
		"name":   r.Name,
		"remark": r.Remark,
	}

	if r.ID > 0 {
		return models.UpdateRuleGroup(r.ID, data)
	}

	data["type"] = r.Type
	return models.AddRuleGroup(data)
}

func (r *RuleGroup) Delete() error {
	return models.DeleteRuleGroup(r.ID)
}

func (r *RuleGroup) Get() (*models.RuleGroup, error) {
	return models.GetRuleGroup(r.ID)
}

func (r *RuleGroup) GetList() ([]*models.RuleGroup, uint, error) {
	var query = make(map[string]interface{})
	if len(r.Name) > 0 {
		query["name"] = r.Name
	}
	if r.Type != 0 {
		query["type"] = r.Type
	}

	return models.GetRuleGroups(query, r.Page, r.PageSize)
}
