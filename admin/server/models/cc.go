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
	Match     string `json:"match" gorm:"column:match;not null"`
	Threshold int    `json:"threshold" gorm:"column:threshold;not null"`
	Duration  int    `json:"duration" gorm:"column:duration;not null"`
	Action    string `json:"action" gorm:"column:action;not null"`
	Remark    string `json:"remark" gorm:"column:remark;not null"`
}

func AddCC(data map[string]interface{}) error {
	var site Site
	site.Model.ID = data["site"].(uint)
	return db.Model(&site).Association("CCs").Append(&CC{
		Name:      data["name"].(string),
		Mode:      data["mode"].(string),
		Method:    data["method"].(string),
		URI:       data["uri"].(string),
		Match:     data["match"].(string),
		Threshold: data["threshold"].(int),
		Duration:  data["duration"].(int),
		Action:    data["action"].(string),
		Remark:    data["remark"].(string),
	}).Error
}

func UpdateCC(id uint, data map[string]interface{}) error {
	return db.Model(&CC{}).Where("id = ?", id).Update(data).Error
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
	return db.Where("id = ?", id).Delete(CC{}).Error
}
