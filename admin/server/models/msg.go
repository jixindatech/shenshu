package models

import "github.com/jinzhu/gorm"

type Msg struct {
	Model

	Name    string `json:"name" gorm:"column:name;unique;comment:'名称'"`
	Content string `json:"content" gorm:"column:content;comment:'内容'"`

	Remark string `json:"remark" gorm:"column:remark;comment:'备注'"`
}

func AddMsg(data map[string]interface{}) error {
	msg := Msg{
		Name:    data["name"].(string),
		Content: data["content"].(string),
		Remark:  data["remark"].(string),
	}
	return db.Create(&msg).Error
}

func DeleteMsg(id uint) error {
	return db.Where("id = ?", id).Delete(Msg{}).Error
}

func UpdateMsg(id uint, data map[string]interface{}) error {
	return db.Model(&Msg{}).Where("id = ?", id).Update(data).Error
}

func GetMsg(id uint) (*Msg, error) {
	var msg Msg
	err := db.Where("id = ?", id).First(&msg).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &msg, nil
}

func GetMsgs(query map[string]interface{}, page int, pageSize int) ([]*Msg, uint, error) {
	var msgs []*Msg
	var count uint
	var err error

	pageNum := (page - 1) * pageSize

	var name string
	if query["name"] != nil {
		name = query["name"].(string)
	}

	if len(name) > 0 {
		name = "%" + name + "%"
		err = db.Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&msgs).Count(&count).Error
	} else {
		err = db.Offset(pageNum).Limit(pageSize).Find(&msgs).Count(&count).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return msgs, count, nil
}
