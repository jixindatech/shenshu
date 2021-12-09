package models

import "github.com/jinzhu/gorm"

type Email struct {
	Model

	Host     string `json:"host" gorm:"column:host;unique;comment:'用户'"`
	Port     int    `json:"port" gorm:"column:port;comment:'用户'"`
	Sender   string `json:"sender" gorm:"column:sender;comment:'用户'"`
	Password string `json:"password" gorm:"column:password;comment:'用户'"`
}

func AddEmail(data map[string]interface{}) error {
	email := &Email{
		Host:     data["host"].(string),
		Port:     data["port"].(int),
		Sender:   data["sender"].(string),
		Password: data["password"].(string),
	}

	return db.Create(&email).Error
}

func GetEmail() (*Email, error) {
	var email Email
	err := db.Where("id = ?", 1).First(&email).Error
	if err == gorm.ErrRecordNotFound {
		return &email, nil
	}

	return &email, err
}

func UpdateEmail(id uint, data map[string]interface{}) error {
	return db.Model(&Email{}).Where("id = ?", id).Update(data).Error
}
