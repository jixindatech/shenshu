package service

import (
	"admin/server/models"
	"admin/server/pkg/rule"
	"gorm.io/datatypes"
)

type Rule struct {
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

func (r *Rule) Save() (err error) {
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
		err = models.UpdateRule(r.ID, data)
	} else {
		data["rulegroup"] = r.RuleGroup
		err = models.AddRule(data)
	}

	return err
	// return SetupSites()
}

func (r *Rule) Get() (*models.Rule, error) {
	return models.GetRule(r.ID)
}

func (r *Rule) GetList() ([]*models.Rule, int, error) {
	data := make(map[string]interface{})
	data["rulegroup"] = r.RuleGroup
	data["name"] = r.Name
	data["page"] = r.Page
	data["pagesize"] = r.PageSize

	return models.GetRules(data)
}

func (r *Rule) Delete() error {
	return models.DeleteRule(r.ID)
}
