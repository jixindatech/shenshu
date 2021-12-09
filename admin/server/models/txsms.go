package models

import "github.com/jinzhu/gorm"

type Txsms struct {
	Model

	Host      string `json:"host" gorm:"column:host;comment:'host'"`
	SecretID  string `json:"secretId" gorm:"column:secret_id;comment:'secretId'"`
	SecretKey string `json:"secretKey" gorm:"column:secret_key;comment:'secretKey'"`
}

func AddTxsms(data map[string]interface{}) error {
	txsms := &Txsms{
		Host:      data["host"].(string),
		SecretID:  data["secretId"].(string),
		SecretKey: data["secretKey"].(string),
	}

	return db.Create(&txsms).Error
}

func GetTxsms() (*Txsms, error) {
	var txsms Txsms
	err := db.Where("id = ?", 1).First(&txsms).Error
	if err == gorm.ErrRecordNotFound {
		return &txsms, nil
	}

	return &txsms, err
}

func UpdateTxsms(id uint, data map[string]interface{}) error {
	return db.Model(&Txsms{}).Where("id = ?", id).Update(data).Error
}
