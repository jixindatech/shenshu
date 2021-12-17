package models

import (
	"github.com/jinzhu/gorm"
)

type RuleBatch struct {
	Model

	Name         string `json:"name" gorm:"column:name;not null"`
	BatchGroupID uint   `json:"ruleGroup" gorm:"not null"`
	Pattern      string `json:"pattern" gorm:"column:pattern;not null"`
	Action       int    `json:"action" gorm:"column:action;not null;default 0"`
	Status       int    `json:"status" gorm:"column:status;default:'0'"`
	Remark       string `json:"remark" gorm:"column:remark;comment:'备注'"`
}

func AddBatchRule(data map[string]interface{}) error {
	var ruleGroup BatchGroup
	ruleGroup.Model.ID = data["rulegroup"].(uint)
	rule := RuleBatch{
		Name:    data["name"].(string),
		Pattern: data["pattern"].(string),
		Action:  data["action"].(int),
		Status:  data["status"].(int),
		Remark:  data["remark"].(string),
	}

	return db.Debug().Model(&ruleGroup).Association("RuleBatchs").Append(&rule).Error
}

func UpdateBatchRule(id uint, data map[string]interface{}) error {
	return db.Model(&RuleBatch{}).Where("id = ?", id).Update(data).Error
}

func GetBatchRule(id uint) (*RuleBatch, error) {
	var rule RuleBatch

	err := db.Where("id = ?", id).Find(&rule).Error
	if err != nil {
		return &rule, err
	}

	return &rule, nil
}

func GetBatchRules(data map[string]interface{}) ([]*RuleBatch, int, error) {
	var rules []*RuleBatch
	rulegroup := data["rulegroup"].(uint)
	name := data["name"].(string)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	search := make(map[string]interface{})
	if data["status"] != nil {
		search["status"] = data["status"].(int)
	}

	var err error
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where(search).Where("batch_group_id = ?", rulegroup).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
		} else {
			err = db.Where(search).Where("batch_group_id = ?", rulegroup).Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
		}
	} else {
		err = db.Where(search).Where("batch_group_id = ?", rulegroup).Find(&rules).Count(&count).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []*RuleBatch{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return rules, count, nil
}

func DeleteBatchRule(id uint) error {
	return db.Where("id = ?", id).Delete(RuleBatch{}).Error
}
