package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type RuleSpeicifc struct {
	Model
	Name            string `json:"name" gorm:"column:name;not null"`
	SpecificGroupId uint   `json:"ruleGroup" gorm:"not null"`

	Rules    datatypes.JSON `json:"rules" gorm:"column:rules;type:VARBINARY(1024);not null;comment:'规则'"`
	Action   int            `json:"action" gorm:"column:action;default:'0'"`
	Priority int            `json:"priority" gorm:"column:priority;default:'0'"`
	Status   int            `json:"status" gorm:"column:status;default:'0'"`
	Remark   string         `json:"remark" gorm:"column:remark;"`
}

func (r *RuleSpeicifc) AfterSave(tx *gorm.DB) (err error) {
	return changeSpecificGroupTimestamp(r.SpecificGroupId)
}

func (r *RuleSpeicifc) AfterDelete(tx *gorm.DB) (err error) {
	return changeSpecificGroupTimestamp(r.SpecificGroupId)
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

	return db.Model(&ruleGroup).Association("RuleSpecifics").Append(&rule).Error
}

func UpdateRuleSpecific(id uint, data map[string]interface{}) error {
	rule, err := GetRuleSpecific(id)
	if err != nil {
		return err
	}

	return db.Model(&rule).Update(data).Error
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

	search := make(map[string]interface{})
	if status != 0 {
		search["status"] = status
	}
	if rulegroup != 0 {
		search["specific_group_id"] = rulegroup
	}

	var err error
	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Order("priority DESC", true).Where(search).Where("name LIKE ?", name).Order("priority DESC").Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error

		} else {
			err = db.Order("priority DESC", true).Where(search).Order("priority DESC").Offset(offset).Limit(pageSize).Find(&rules).Count(&count).Error
		}
	} else {
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where(search).Where("name LIKE ?", name).Order("priority DESC").Find(&rules).Count(&count).Error

		} else {
			err = db.Order("priority DESC", true).Where(search).Order("priority DESC").Find(&rules).Count(&count).Error
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
	rule, err := GetRuleSpecific(id)
	if err != nil {
		return err
	}

	return db.Delete(&rule).Error
}
