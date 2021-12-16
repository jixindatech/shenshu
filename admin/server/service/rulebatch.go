package service

import (
	"admin/server/models"
)

type RuleBatch struct {
	ID uint

	RuleGroup uint
	Name      string
	Pattern   string
	Action    int
	Status    int
	Remark    string

	Page     int
	PageSize int
}

func (b *RuleBatch) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = b.Name
	data["pattern"] = b.Pattern
	data["action"] = b.Action
	data["status"] = b.Status
	data["remark"] = b.Remark

	if b.ID > 0 {
		err = models.UpdateBatchRule(b.ID, data)
	} else {
		data["rulegroup"] = b.RuleGroup
		err = models.AddBatchRule(data)
	}

	return err
}

func (b *RuleBatch) Get() (*models.RuleBatch, error) {
	return models.GetBatchRule(b.ID)
}

func (b *RuleBatch) GetList() ([]*models.RuleBatch, int, error) {
	data := make(map[string]interface{})
	data["rulegroup"] = b.RuleGroup
	data["name"] = b.Name
	data["page"] = b.Page
	data["pagesize"] = b.PageSize

	return models.GetBatchRules(data)
}

func (b *RuleBatch) Delete() error {
	return models.DeleteBatchRule(b.ID)
}
