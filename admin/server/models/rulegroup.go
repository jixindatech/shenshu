package models

import "github.com/jinzhu/gorm"

type RuleGroup struct {
	Model

	Name   string `json:"name" gorm:"column:name;not null"`
	Remark string `json:"remark" gorm:"column:remark;"`
	Rules  []*Rule
}

func AddRuleGroup(data map[string]interface{}) error {
	ruleGroup := RuleGroup{
		Name:   data["name"].(string),
		Remark: data["remark"].(string),
	}
	return db.Create(&ruleGroup).Error
}

func DeleteRuleGroup(id uint) error {
	// association delete sites
	return db.Where("id = ?", id).Delete(RuleGroup{}).Error
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

	if len(name) > 0 {
		name = "%" + name + "%"
		err = db.Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Count(&count).Error
	} else {
		err = db.Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Count(&count).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []*RuleGroup{}, 0, nil
	}

	if err != nil {
		return nil, count, err
	}

	return ruleGroups, count, nil
}
