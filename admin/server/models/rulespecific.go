package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type RuleSpeicifc struct {
	Model
	Name        string `json:"name" gorm:"column:name;not null"`
	RuleGroupId uint   `json:"ruleGroup" gorm:"not null"`

	Rules    datatypes.JSON `json:"rules" gorm:"column:rules;type:VARBINARY(1024);not null;comment:'规则'"`
	Action   int            `json:"action" gorm:"column:action;default:'0'"` // deny score
	Priority int            `json:"priority" gorm:"column:priority;default:'0'"`
	Status   int            `json:"status" gorm:"column:status;default:'0'"`
	Remark   string         `json:"remark" gorm:"column:remark;"`
}

func AddRuleSpecific(data map[string]interface{}) error {
	var ruleGroup SpecificGroup
	ruleGroup.Model.ID = data["rulegroup"].(uint)
	rule := RuleSpeicifc{
		Name:     data["name"].(string),
		Rules:    data["rules"].(datatypes.JSON),
		Priority: data["priority"].(int),
		Action:   data["action"].(int),
		Status:   data["status"].(int),
		Remark:   data["remark"].(string),
	}

	return db.Model(&ruleGroup).Association("Rules").Append(&rule).Error
}

func UpdateRuleSpecific(id uint, data map[string]interface{}) error {
	return db.Model(&RuleSpeicifc{}).Where("id = ?", id).Update(data).Error
}

func GetRuleSpecific(id uint) (*RuleSpeicifc, error) {
	var rule RuleSpeicifc

	err := db.Where("id = ?", id).Find(&rule).Error
	if err != nil {
		return &rule, err
	}

	return &rule, nil
}

func GetRuleSpecifics(data map[string]interface{}, page, pageSize int) ([]*RuleSpeicifc, int, error) {
	var rules []*RuleSpeicifc
	rulegroup := data["rulegroup"].(uint)
	name := data["name"].(string)
	status := data["status"].(int)

	var err error
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			if status != 0 {
				err = db.Where("status = ?", status).Where("rule_group_id = ?", rulegroup).Where("name LIKE ?", name).Order("priority DESC").Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
			} else {
				err = db.Where("rule_group_id = ?", rulegroup).Where("name LIKE ?", name).Order("priority DESC").Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
			}
		} else {
			if status != 0 {
				err = db.Where("status = ?", status).Where("rule_group_id = ?", rulegroup).Order("priority DESC").Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
			} else {
				err = db.Where("rule_group_id = ?", rulegroup).Order("priority DESC").Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
			}
		}
	} else {
		if len(name) > 0 {
			name = "%" + name + "%"
			if status != 0 {
				err = db.Where("status = ?", status).Where("rule_group_id = ?", rulegroup).Where("name LIKE ?", name).Order("priority DESC").Find(&rules).Count(&count).Error
			} else {
				err = db.Where("rule_group_id = ?", rulegroup).Where("name LIKE ?", name).Order("priority DESC").Find(&rules).Count(&count).Error
			}
		} else {
			if status != 0 {
				err = db.Debug().Where("status = ?", status).Where("rule_group_id = ?", rulegroup).Order("priority DESC").Find(&rules).Count(&count).Error
			} else {
				err = db.Where("rule_group_id = ?", rulegroup).Order("priority DESC").Find(&rules).Count(&count).Error
			}
		}
	}

	if err == gorm.ErrRecordNotFound {
		return []*RuleSpeicifc{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}

	return rules, count, nil
}

func DeleteRuleSpecific(id uint) error {
	return db.Where("id = ?", id).Delete(RuleSpeicifc{}).Error
}
