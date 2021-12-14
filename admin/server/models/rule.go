package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type Rule struct {
	Model
	Name        string `json:"name" gorm:"column:name;not null"`
	RuleGroupId uint   `json:"ruleGroup" gorm:"not null"`

	Rules    datatypes.JSON `json:"rules" gorm:"column:rules;type:VARBINARY(1024);not null;comment:'规则'"`
	Action   int            `json:"action" gorm:"column:action;default:'0'"` // deny score
	Priority int            `json:"priority" gorm:"column:priority;default:'0'"`
	Status   int            `json:"status" gorm:"column:status;default:'0'"`
	Remark   string         `json:"remark" gorm:"column:remark;"`
}

func AddRule(data map[string]interface{}) error {
	var ruleGroup RuleGroup
	ruleGroup.Model.ID = data["rulegroup"].(uint)
	rule := Rule{
		Name:     data["name"].(string),
		Rules:    data["rules"].(datatypes.JSON),
		Priority: data["priority"].(int),
		Action:   data["action"].(int),
		Status:   data["status"].(int),
		Remark:   data["remark"].(string),
	}

	return db.Model(&ruleGroup).Association("Rules").Append(&rule).Error
}

func UpdateRule(id uint, data map[string]interface{}) error {
	return db.Model(&Rule{}).Where("id = ?", id).Update(data).Error
}

func GetRule(id uint) (*Rule, error) {
	var rule Rule

	err := db.Where("id = ?", id).Find(&rule).Error
	if err != nil {
		return &rule, err
	}

	return &rule, nil
}

func GetRules(data map[string]interface{}) ([]*Rule, int, error) {
	var rules []*Rule
	rulegroup := data["rulegroup"].(uint)
	name := data["name"].(string)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	var err error
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where("rule_group_id = ?", rulegroup).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
		} else {
			err = db.Where("rule_group_id = ?", rulegroup).Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
		}
	} else {
		err = db.Where("rule_group_id = ?", rulegroup).Find(&rules).Count(&count).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []*Rule{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return rules, count, nil
}

func DeleteRule(id uint) error {
	return db.Where("id = ?", id).Delete(Rule{}).Error
}
