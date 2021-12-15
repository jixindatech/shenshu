package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type RuleGroup struct {
	Model

	Name       string         `json:"name" gorm:"column:name;not null"`
	Type       int            `json:"type" gorm:"column:type;not null"`
	Status     int            `json:"status" gorm:"column:status;default:'0';comment:'模式'"`
	Level      int            `json:"level" gorm:"column:level;default:'0';comment:'等级'"`
	Decoder    datatypes.JSON `json:"decoder" gorm:"type:VARBINARY(1024);column:decoder"`
	Priority   int            `json:"priority" gorm:"column:priority;not null"`
	Remark     string         `json:"remark" gorm:"column:remark;"`
	Rules      []*Rule
	RuleBatchs []*RuleBatch
}

func AddRuleGroup(data map[string]interface{}) error {
	ruleGroup := RuleGroup{
		Name:     data["name"].(string),
		Type:     data["type"].(int),
		Priority: data["priority"].(int),
		Status:   data["status"].(int),
		Level:    data["level"].(int),
		Decoder:  data["decoder"].(datatypes.JSON),

		Remark: data["remark"].(string),
	}
	return db.Create(&ruleGroup).Error
}

func DeleteRuleGroup(id uint) error {
	var ruleGroup RuleGroup
	ruleGroup.Model.ID = id
	ruleCount := db.Model(&ruleGroup).Association("Rules").Count()
	ruleBatchCount := db.Model(&ruleGroup).Association("RuleBatchs").Count()
	if ruleCount != 0 || ruleBatchCount != 0 {
		return fmt.Errorf("%s", "rule exist in this group")
	}

	return db.Delete(&ruleGroup).Error
}

func UpdateRuleGroup(id uint, data map[string]interface{}) error {
	return db.Model(&RuleGroup{}).Where("id = ?", id).Update(data).Error
}

func GetRuleGroup(id uint) (*RuleGroup, error) {
	var ruleGroup RuleGroup
	err := db.Where("id = ?", id).First(&ruleGroup).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &ruleGroup, nil
}

func GetRuleGroups(query map[string]interface{}, page int, pageSize int) ([]*RuleGroup, uint, error) {
	var ruleGroups []*RuleGroup
	var count uint
	var err error

	pageNum := (page - 1) * pageSize

	var name string
	if query["name"] != nil {
		name = query["name"].(string)
	}
	var ruleType int
	if query["type"] != nil {
		ruleType = query["type"].(int)
	}
	if len(name) > 0 {
		name = "%" + name + "%"
		if ruleType != 0 {
			err = db.Where("type = ?", ruleType).Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority").Count(&count).Error
		} else {
			err = db.Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority").Count(&count).Error
		}
	} else {
		if ruleType != 0 {
			err = db.Where("type = ?", ruleType).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority").Count(&count).Error
		} else {
			err = db.Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority").Count(&count).Error
		}
	}

	if err == gorm.ErrRecordNotFound {
		return []*RuleGroup{}, 0, nil
	}

	if err != nil {
		return nil, count, err
	}

	return ruleGroups, count, nil
}
