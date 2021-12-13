package models

import "github.com/jinzhu/gorm"

type Site struct {
	Model

	Name   string `json:"name" gorm:"column:name;not null;unique"`
	Host   string `json:"host" gorm:"column:host;not null;"`
	Path   string `json:"path" gorm:"column:path;not null"`
	Remark string `json:"remark" gorm:"column:remark"`

	Upstreams []*Upstream `json:"upstreamRef" gorm:"many2many:site_upstream;"`
	IPs       []IP        `json:"ips"`
}

func AddSite(data map[string]interface{}) error {
	site := &Site{
		Name:   data["name"].(string),
		Host:   data["host"].(string),
		Path:   data["path"].(string),
		Remark: data["remark"].(string),
	}

	err := db.Create(&site).Error
	if err != nil {
		return err
	}

	/* for global site */
	_, ok := data["upstreamRef"]
	if !ok {
		return nil
	}

	var upstreams []*Upstream
	temp := Upstream{}
	temp.Model.ID = data["upstreamRef"].(uint)
	upstreams = append(upstreams, &temp)

	err = db.Debug().Model(&site).Association("Upstreams").Replace(upstreams).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteSite(id uint) error {
	site := Site{}
	site.Model.ID = id
	/* clear Associations with upstreams*/
	db.Model(&site).Association("Upstreams").Clear()

	/* clear Associations with upstreams*/
	db.Model(&site).Association("IPs").Clear()

	err := db.Delete(&site).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateSite(id uint, data map[string]interface{}) error {
	site := Site{}
	site.Model.ID = id

	var upstreams []*Upstream
	temp := Upstream{}
	temp.Model.ID = data["upstreamRef"].(uint)
	upstreams = append(upstreams, &temp)

	delete(data, "upstreamRef")

	err := db.Model(&site).Update(data).Error
	if err != nil {
		return err
	}

	err = db.Debug().Model(&site).Association("Upstreams").Replace(upstreams).Error
	if err != nil {
		return err
	}

	return nil
}

func GetSite(id uint) (*Site, error) {
	var site Site

	err := db.Preload("Upstreams").Where("id = ?", id).Find(&site).Error
	if err != nil {
		return &site, err
	}

	return &site, nil
}

func GetSites(data map[string]interface{}) ([]*Site, int, error) {
	var sites []*Site
	name := data["name"].(string)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err := db.Preload("Upstreams").Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&sites).Count(&count).Error
			if err != nil {
				return nil, 0, err
			}
		} else {
			err := db.Preload("Upstreams").Offset(offset).Limit(pageSize).Find(&sites).Count(&count).Error
			if err != nil {
				return nil, 0, err
			}
		}
	} else {
		err := db.Preload("Upstreams").Find(&sites).Count(&count).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	return sites, count, nil
}
