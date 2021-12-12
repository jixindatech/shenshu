package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type IP struct {
	Model

	Name   string         `json:"name" gorm:"column:name;not null"`
	Type   int            `json:"type" gorm:"column:type;not null"`
	IP     datatypes.JSON `json:"ip" gorm:"column:ip;not null"`
	Remark string         `json:"remark" gorm:"column:remark;"`

	Sites []Site `json:"sites" gorm:"many2many:site_ip;"`
}

func AddIP(data map[string]interface{}) error {
	ip := &IP{
		Name:   data["name"].(string),
		Type:   data["type"].(int),
		IP:     data["ip"].([]byte),
		Remark: data["remark"].(string),
	}

	err := db.Create(&ip).Error
	if err != nil {
		return err
	}

	return nil
}

func GetIPs(data map[string]interface{}) ([]*IP, int, error) {
	var ips []*IP
	name := data["name"].(string)
	ipType := data["type"].(int)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err := db.Preload("Sites").Where("type = ?", ipType).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&ips).Count(&count).Error
			if err != nil {
				return nil, 0, err
			}
		} else {
			err := db.Preload("Sites").Where("type = ?", ipType).Offset(offset).Limit(pageSize).Find(&ips).Count(&count).Error
			if err != nil {
				return nil, 0, err
			}
		}
	} else {
		err := db.Preload("Sites").Where("type = ?", ipType).Find(&ips).Count(&count).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	return ips, count, nil
}

func UpdateIP(id uint, data map[string]interface{}) error {
	ip := IP{}
	ip.Model.ID = id

	var sites []*Site
	for _, item := range data["sites"].([]uint) {
		temp := Site{}
		temp.Model.ID = item
		sites = append(sites, &temp)
	}

	return db.Model(&ip).Association("Sites").Replace(sites).Error
}

func DeleteIP(id uint) error {
	ip := IP{}
	ip.Model.ID = id
	/* clear Associations with sites*/
	db.Model(&ip).Association("Sites").Clear()

	err := db.Delete(&ip).Error
	if err != nil {
		return err
	}
	return nil
}

func GetIP(id uint) (*IP, error) {
	var ip IP

	err := db.Preload("Sites").Where("id = ?", id).Find(&ip).Error
	if err != nil {
		return &ip, err
	}

	return &ip, nil
}
