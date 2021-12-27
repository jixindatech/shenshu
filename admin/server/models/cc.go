package models

import (
	"github.com/jinzhu/gorm"
)

type CC struct {
	Model

	Name      string `json:"name" gorm:"column:name;not null"`
	SiteId    uint   `json:"site" gorm:"not null"`
	Mode      string `json:"mode" gorm:"column:mode;not null"`
	Method    string `json:"method" gorm:"column:method;not null"`
	URI       string `json:"uri" gorm:"column:uri;not null"`
	Threshold int    `json:"threshold" gorm:"column:threshold;not null"`
	Action    string `json:"action" gorm:"column:action;not null"`
	Remark    string `json:"remark" gorm:"column:remark;not null"`
}

func (c *CC) AfterSave(tx *gorm.DB) (err error) {
	changeSiteTimestamp(c.SiteId, "CCTimestamp")
	return nil
}

func (c *CC) AfterDelete(tx *gorm.DB) (err error) {
	changeSiteTimestamp(c.SiteId, "CCTimestamp")
	return nil
}

func AddCC(data map[string]interface{}) error {
	var site Site
	site.Model.ID = data["site"].(uint)
	return db.Model(&site).Association("CCs").Append(&CC{
		Name:      data["name"].(string),
		Mode:      data["mode"].(string),
		Method:    data["method"].(string),
		URI:       data["uri"].(string),
		Threshold: data["threshold"].(int),
		Action:    data["action"].(string),
		Remark:    data["remark"].(string),
	}).Error
}

func UpdateCC(id uint, data map[string]interface{}) error {
	cc, err := GetCC(id)
	if err != nil {
		return err
	}

	return db.Model(&cc).Update(data).Error
}

func GetCC(id uint) (*CC, error) {
	var cc CC

	err := db.Where("id = ?", id).Find(&cc).Error
	if err != nil {
		return &cc, err
	}

	return &cc, nil
}

func GetCCs(data map[string]interface{}) ([]*CC, int, error) {
	var ccs []*CC
	site := data["site"].(uint)
	name := data["name"].(string)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	var err error
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where("site_id = ?", site).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&ccs).Count(&count).Error
		} else {
			err = db.Where("site_id = ?", site).Offset(offset).Limit(pageSize).Find(&ccs).Count(&count).Error
		}
	} else {
		err = db.Where("site_id = ?", site).Find(&ccs).Count(&count).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []*CC{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return ccs, count, nil
}

func DeleteCC(id uint) error {
	cc, err := GetCC(id)
	if err != nil {
		return err
	}

	return db.Delete(&cc).Error
}
