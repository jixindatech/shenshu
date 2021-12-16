package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type BatchGroup struct {
	Model

	Name     string         `json:"name" gorm:"column:name;not null"`
	Status   int            `json:"status" gorm:"column:status;default:'2';comment:'状态'"`
	Action   int            `json:"action" gorm:"column:action;default:'1';comment:'动作'"`
	Level    int            `json:"level" gorm:"column:level;default:'0';comment:'等级'"`
	Decoder  datatypes.JSON `json:"decoder" gorm:"type:VARBINARY(1024);column:decoder"`
	Priority int            `json:"priority" gorm:"column:priority;not null"`
	Remark   string         `json:"remark" gorm:"column:remark;"`

	RuleBatchs []*RuleBatch
}

func AddBatchGroup(data map[string]interface{}) error {
	batchGroup := BatchGroup{
		Name:     data["name"].(string),
		Priority: data["priority"].(int),
		Status:   data["status"].(int),
		Action:   data["action"].(int),
		Level:    data["level"].(int),
		Decoder:  data["decoder"].(datatypes.JSON),

		Remark: data["remark"].(string),
	}
	return db.Create(&batchGroup).Error
}

func DeleteBatchGroup(id uint) error {
	var batchGroup BatchGroup
	batchGroup.Model.ID = id
	ruleCount := db.Model(&batchGroup).Association("Rules").Count()
	ruleBatchCount := db.Model(&batchGroup).Association("RuleBatchs").Count()
	if ruleCount != 0 || ruleBatchCount != 0 {
		return fmt.Errorf("%s", "rule exist in this group")
	}

	return db.Delete(&batchGroup).Error
}

func UpdateBatchGroup(id uint, data map[string]interface{}) error {
	return db.Model(&BatchGroup{}).Where("id = ?", id).Update(data).Error
}

func GetBatchGroup(id uint) (*BatchGroup, error) {
	var batchGroup BatchGroup
	err := db.Where("id = ?", id).First(&batchGroup).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &batchGroup, nil
}

func GetBatchGroups(query map[string]interface{}, page int, pageSize int) ([]*BatchGroup, uint, error) {
	var batchGroups []*BatchGroup
	var count uint
	var err error

	var name string
	if query["name"] != nil {
		name = query["name"].(string)
	}
	var status int
	if query["status"] != nil {
		status = query["status"].(int)
	}

	if page == 0 {
		if len(name) > 0 {
			name = "%" + name + "%"
			if status != 0 {
				err = db.Order("priority DESC", true).Where("status = ?", status).Where("name like ?", name).Find(&batchGroups).Count(&count).Error
			} else {
				err = db.Order("priority DESC", true).Where("name like ?", name).Find(&batchGroups).Count(&count).Error
			}
		} else {
			err = db.Order("priority DESC", true).Find(&batchGroups).Count(&count).Error
		}
	} else {
		pageNum := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Order("priority DESC", true).Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&batchGroups).Count(&count).Error

		} else {
			err = db.Order("priority DESC", true).Offset(pageNum).Limit(pageSize).Find(&batchGroups).Count(&count).Error
		}
	}

	if err == gorm.ErrRecordNotFound {
		return []*BatchGroup{}, 0, nil
	}

	if err != nil {
		return nil, count, err
	}

	return batchGroups, count, nil
}
