package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type SpecificGroup struct {
	Model

	Name     string         `json:"name" gorm:"column:name;not null"`
	Type     int            `json:"type" gorm:"column:type;not null"`
	Status   int            `json:"status" gorm:"column:status;default:'2';comment:'状态'"`
	Action   int            `json:"action" gorm:"column:action;default:'1';comment:'动作'"`
	Level    int            `json:"level" gorm:"column:level;default:'0';comment:'等级'"`
	Decoder  datatypes.JSON `json:"decoder" gorm:"type:VARBINARY(1024);column:decoder"`
	Priority int            `json:"priority" gorm:"column:priority;not null"`
	Remark   string         `json:"remark" gorm:"column:remark;"`

	SpecificRules []*RuleSpeicifc
}

func AddSpecificGroup(data map[string]interface{}) error {
	specificGroup := SpecificGroup{
		Name:     data["name"].(string),
		Type:     data["type"].(int),
		Priority: data["priority"].(int),
		Status:   data["status"].(int),
		Action:   data["action"].(int),
		Level:    data["level"].(int),
		Decoder:  data["decoder"].(datatypes.JSON),

		Remark: data["remark"].(string),
	}
	return db.Create(&specificGroup).Error
}

func DeleteSpecificGroup(id uint) error {
	var specificGroup SpecificGroup
	specificGroup.Model.ID = id
	ruleCount := db.Model(&specificGroup).Association("Rules").Count()
	ruleBatchCount := db.Model(&specificGroup).Association("RuleBatchs").Count()
	if ruleCount != 0 || ruleBatchCount != 0 {
		return fmt.Errorf("%s", "rule exist in this group")
	}

	return db.Delete(&specificGroup).Error
}

func UpdateSpecificGroup(id uint, data map[string]interface{}) error {
	return db.Model(&SpecificGroup{}).Where("id = ?", id).Update(data).Error
}

func GetSpecificGroup(id uint) (*SpecificGroup, error) {
	var specificGroup SpecificGroup
	err := db.Where("id = ?", id).First(&specificGroup).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &specificGroup, nil
}

func GetSpecificGroups(query map[string]interface{}, page int, pageSize int) ([]*SpecificGroup, uint, error) {
	var ruleGroups []*SpecificGroup
	var count uint
	var err error

	var name string
	if query["name"] != nil {
		name = query["name"].(string)
	}
	var ruleType int
	if query["type"] != nil {
		ruleType = query["type"].(int)
	}
	var status int
	if query["status"] != nil {
		status = query["status"].(int)
	}

	if page == 0 {
		if len(name) > 0 {
			name = "%" + name + "%"
			if ruleType != 0 {
				if status != 0 {
					err = db.Where("status = ?", status).Where("type = ?", ruleType).Where("name like ?", name).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
				} else {
					err = db.Where("type = ?", ruleType).Where("name like ?", name).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
				}
			} else {
				if status != 0 {
					err = db.Where("status = ?", status).Where("name like ?", name).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
				} else {
					err = db.Where("name like ?", name).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
				}
			}
		} else {
			if ruleType != 0 {
				err = db.Where("type = ?", ruleType).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
			} else {
				err = db.Find(&ruleGroups).Order("priority DESC").Count(&count).Error
			}
		}
	} else {
		pageNum := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			if ruleType != 0 {
				err = db.Where("type = ?", ruleType).Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
			} else {
				err = db.Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
			}
		} else {
			if ruleType != 0 {
				err = db.Where("type = ?", ruleType).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
			} else {
				err = db.Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Order("priority DESC").Count(&count).Error
			}
		}
	}

	if err == gorm.ErrRecordNotFound {
		return []*SpecificGroup{}, 0, nil
	}

	if err != nil {
		return nil, count, err
	}

	return ruleGroups, count, nil
}
