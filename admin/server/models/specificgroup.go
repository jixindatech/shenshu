package models

import (
	"admin/core/log"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"time"
)

type SpecificGroup struct {
	Model

	Name     string         `json:"name" gorm:"column:name;not null"`
	Status   int            `json:"status" gorm:"column:status;default:'2';comment:'状态'"`
	Action   int            `json:"action" gorm:"column:action;default:'1';comment:'动作'"`
	Level    int            `json:"level" gorm:"column:level;default:'0';comment:'等级'"`
	Decoder  datatypes.JSON `json:"decoder" gorm:"type:VARBINARY(1024);column:decoder"`
	Priority int            `json:"priority" gorm:"column:priority;not null"`
	Remark   string         `json:"remark" gorm:"column:remark;"`

	Sites         []*Site `json:"sites" gorm:"many2many:site_specificgroup;"`
	RuleSpecifics []*RuleSpeicifc
}

func changeSpecificGroupTimestamp(id uint) error {
	group := SpecificGroup{}
	group.Model.ID = id

	err := db.Model(&group).Update("updated_at", time.Now()).Error
	if err != nil {
		log.Logger.Error("specific_group", zap.String("err", err.Error()))
	}

	return err
}

func AddSpecificGroup(data map[string]interface{}) error {
	specificGroup := SpecificGroup{
		Name:     data["name"].(string),
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

	count := db.Model(&specificGroup).Association("Sites").Count()
	if count != 0 {
		return fmt.Errorf("%s", "site exist link in this group")
	}

	count = db.Model(&specificGroup).Association("RuleSpecifics").Count()
	if count != 0 {
		return fmt.Errorf("%s", "rule exist in this group")
	}
	return db.Delete(&specificGroup).Error
}

func UpdateSpecificGroup(id uint, data map[string]interface{}) error {
	group, err := GetSpecificGroup(id)
	if err != nil {
		return err
	}

	err = db.Model(&group).Update(data).Error
	if err != nil {
		return err
	}

	var sites []*Site
	err = db.Model(&group).Association("Sites").Find(&sites).Error
	if err != nil {
		return err
	}

	for _, site := range sites {
		err := db.Model(&site).Update("rule_timestamp", time.Now().Unix()).Error
		if err != nil {
			log.Logger.Error("batch_group", zap.String("err", err.Error()))
		}
	}

	return nil
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

	search := make(map[string]interface{})
	if query["status"] != nil {
		search["status"] = query["status"].(int)
	}

	if page == 0 {
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Order("priority DESC", true).Where("name like ?", name).Where(search).Find(&ruleGroups).Count(&count).Error
		} else {
			err = db.Order("priority DESC", true).Where(search).Find(&ruleGroups).Count(&count).Error
		}
	} else {
		pageNum := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Order("priority DESC", true).Where("name like ?", name).Where(search).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Count(&count).Error
		} else {
			err = db.Order("priority DESC", true).Where(search).Offset(pageNum).Limit(pageSize).Find(&ruleGroups).Count(&count).Error
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
