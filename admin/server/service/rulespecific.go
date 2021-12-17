package service

import (
	"admin/server/models"
	"admin/server/pkg/rule"
	"gorm.io/datatypes"
)

type RuleSpecific struct {
	ID uint

	RuleGroup uint
	Name      string
	Rules     datatypes.JSON
	Action    int
	Priority  int
	Status    int
	Remark    string

	Page     int
	PageSize int
}

func (r *RuleSpecific) Save() (err error) {
	data := make(map[string]interface{})
	if err := rule.ValidateRule(r.Rules); err != nil {
		return err
	}
	data["name"] = r.Name
	data["rules"] = r.Rules
	data["action"] = r.Action
	data["priority"] = r.Priority
	data["status"] = r.Status
	data["remark"] = r.Remark

	if r.ID > 0 {
		err = models.UpdateRuleSpecific(r.ID, data)
	} else {
		data["rulegroup"] = r.RuleGroup
		err = models.AddRuleSpecific(data)
	}

	return err
}

func (r *RuleSpecific) Get() (*models.RuleSpeicifc, error) {
	return models.GetRuleSpecific(r.ID)
}

func (r *RuleSpecific) GetList() ([]*models.RuleSpeicifc, int, error) {
	data := make(map[string]interface{})
	data["rulegroup"] = r.RuleGroup
	data["name"] = r.Name
	data["status"] = r.Status

	return models.GetRuleSpecifics(data, r.Page, r.PageSize)
}

func (r *RuleSpecific) Delete() error {
	return models.DeleteRuleSpecific(r.ID)
}
