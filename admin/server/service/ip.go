package service

import (
	"admin/server/models"
	"encoding/json"
)

type IP struct {
	ID uint

	Name   string
	Type   int
	IP     []string
	Remark string

	Page     int
	PageSize int
}

func (i *IP) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = i.Name
	data["type"] = i.Type

	ip, err := json.Marshal(&i.IP)
	if err != nil {
		return err
	}
	data["ip"] = ip
	data["remark"] = i.Remark

	if i.ID > 0 {
		//err = models.PutSite(i.ID, data)
	} else {
		data["name"] = i.Name
		err = models.AddIP(data)
	}

	if err != nil {
		return err
	}

	// return SetupSites()
	return nil
}

func (i *IP) GetList() ([]*models.IP, int, error) {
	data := make(map[string]interface{})
	data["name"] = i.Name
	data["type"] = i.Type
	data["page"] = i.Page
	data["pagesize"] = i.PageSize

	return models.GetIPs(data)
}
