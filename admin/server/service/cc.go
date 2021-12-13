package service

import "admin/server/models"

type CC struct {
	ID uint

	Name      string
	Site      uint
	Mode      string
	Method    string
	URI       string
	Match     string
	Threshold int
	Duration  int
	Action    string
	Remark    string

	Page     int
	PageSize int
}

func (c *CC) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = c.Name
	data["mode"] = c.Mode
	data["method"] = c.Method
	data["uri"] = c.URI
	data["match"] = c.Match
	data["threshold"] = c.Threshold
	data["duration"] = c.Duration
	data["action"] = c.Action
	data["remark"] = c.Remark

	if c.ID > 0 {
		err = models.UpdateCC(c.ID, data)
	} else {
		err = models.AddCC(data)
	}

	return err
	// return SetupSites()
}

func (c *CC) Get() (*models.CC, error) {
	return models.GetCC(c.ID)
}

func (c *CC) GetList() ([]*models.CC, int, error) {
	data := make(map[string]interface{})
	data["site"] = c.Site
	data["name"] = c.Name
	data["page"] = c.Page
	data["pagesize"] = c.PageSize

	return models.GetCCs(data)
}
