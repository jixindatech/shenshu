package service

import "admin/server/models"

type Rule struct {
	ID uint

	RuleGroup uint
	Name      string
	Remark    string

	Page     int
	PageSize int
}

func (r *Rule) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = r.Name
	data["remark"] = r.Remark

	if r.ID > 0 {
		err = models.UpdateRule(r.ID, data)
	} else {
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
