package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type GlobalIP struct {
	Model

	Name   string         `json:"name" gorm:"column:name;not null"`
	Type   int            `json:"type" gorm:"column:type;not null"`
	IP     datatypes.JSON `json:"ip" gorm:"column:ip;not null"`
	Remark string         `json:"remark" gorm:"column:remark;"`
}

func AddGlobalIP(data map[string]interface{}) error {
	return db.Debug().Create(&GlobalIP{
		Name:   data["name"].(string),
		Type:   data["type"].(int),
		IP:     data["ip"].([]byte),
		Remark: data["remark"].(string)}).Error
}

func GetGlobalIPs(data map[string]interface{}) ([]*GlobalIP, int, error) {
	var ips []*GlobalIP
	name := data["name"].(string)
	ipType := data["type"].(int)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	var err error
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where("type = ?", ipType).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&ips).Count(&count).Error
		} else {
			err = db.Where("type = ?", ipType).Offset(offset).Limit(pageSize).Find(&ips).Count(&count).Error
		}
	} else {
		err = db.Where("type = ?", ipType).Find(&ips).Count(&count).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []*GlobalIP{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return ips, count, nil
}

func UpdateGlobalIP(id uint, data map[string]interface{}) error {
	ip, err := GetGlobalIP(id)
	if err != nil {
		return err
	}

	return db.Model(&ip).Update(data).Error
}

func DeleteGlobalIP(id uint) error {
	ip, err := GetGlobalIP(id)
	if err != nil {
		return err
	}

	return db.Delete(&ip).Error
}

func GetGlobalIP(id uint) (*GlobalIP, error) {
	var ip GlobalIP

	err := db.Where("id = ?", id).Find(&ip).Error
	if err != nil {
		return &ip, err
	}

	return &ip, nil
}
