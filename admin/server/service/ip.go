package service

import (
	"admin/server/models"
	"encoding/json"
)

type IP struct {
	ID uint

	Site   uint
	Name   string
	Type   int
	IP     []string
	Remark string

	Sites []uint

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
		return models.UpdateIP(i.ID, data)
	}

	data["site"] = i.Site
	err = models.AddIP(data)

	if err != nil {
		return err
	}

	// return SetupSites()
	return nil
}

func (i *IP) GetList() ([]*models.IP, int, error) {
	data := make(map[string]interface{})
	data["site"] = i.Site
	data["name"] = i.Name
	data["type"] = i.Type
	data["page"] = i.Page
	data["pagesize"] = i.PageSize

	return models.GetIPs(data)
}

func (i *IP) Get() (*models.IP, error) {
	return models.GetIP(i.ID)
}

func (i *IP) Delete() error {
	err := models.DeleteIP(i.ID)
	if err != nil {
		return err
	}

	return nil
	//return SetupSites()
}
