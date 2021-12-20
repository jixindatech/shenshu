package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type IP struct {
	Model

	SiteId uint           `json:"site" gorm:"not null"`
	Name   string         `json:"name" gorm:"column:name;not null"`
	Type   int            `json:"type" gorm:"column:type;not null"`
	IP     datatypes.JSON `json:"ip" gorm:"column:ip;not null"`
	Remark string         `json:"remark" gorm:"column:remark;"`
}

func (i *IP) AfterSave(tx *gorm.DB) (err error) {
	changeSiteTimestamp(i.SiteId, "IPTimestamp")
	return nil
}

func (i *IP) AfterDelete(tx *gorm.DB) (err error) {
	changeSiteTimestamp(i.SiteId, "IPTimestamp")
	return nil
}

func AddIP(data map[string]interface{}) error {
	var site Site
	site.Model.ID = data["site"].(uint)
	return db.Model(&site).Association("IPs").Append(&IP{
		Name:   data["name"].(string),
		Type:   data["type"].(int),
		IP:     data["ip"].([]byte),
		Remark: data["remark"].(string),
	}).Error
}

func GetIPs(data map[string]interface{}) ([]*IP, int, error) {
	var ips []*IP
	name := data["name"].(string)
	ipType := data["type"].(int)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	var err error
	site := data["site"].(uint)
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where("site_id = ?", site).Where("type = ?", ipType).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&ips).Count(&count).Error
		} else {
			err = db.Where("site_id = ?", site).Where("type = ?", ipType).Offset(offset).Limit(pageSize).Find(&ips).Count(&count).Error
		}
	} else {
		err = db.Where("site_id = ?", site).Where("type = ?", ipType).Find(&ips).Count(&count).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []*IP{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return ips, count, nil
}

func UpdateIP(id uint, data map[string]interface{}) error {
	ip, err := GetIP(id)
	if err != nil {
		return err
	}

	return db.Model(&ip).Update(data).Error
}

func DeleteIP(id uint) error {
	ip, err := GetIP(id)
	if err != nil {
		return err
	}

	return db.Delete(&ip).Error
}

func GetIP(id uint) (*IP, error) {
	var ip IP

	err := db.Where("id = ?", id).Find(&ip).Error
	if err != nil {
		return &ip, err
	}

	return &ip, nil
}
