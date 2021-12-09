package models

import "github.com/jinzhu/gorm"

type Ldap struct {
	Model

	Type     string `json:"type" gorm:"column:type;comment:'类型'"`
	Host     string `json:"host" gorm:"column:host;comment:'主机'"`
	Port     int    `json:"port" gorm:"column:port;comment:'端口'"`
	DN       string `json:"dn" gorm:"column:dn;comment:'用户'"`
	BaseDN   string `json:"basedn" gorm:"column:basedn;comment:'dn'"`
	Password string `json:"password" gorm:"column:password;comment:'密码'"`
}

func AddLdap(data map[string]interface{}) error {
	ldap := &Ldap{
		Type:     data["type"].(string),
		Host:     data["host"].(string),
		Port:     data["port"].(int),
		DN:       data["dn"].(string),
		BaseDN:   data["basedn"].(string),
		Password: data["password"].(string),
	}

	return db.Create(&ldap).Error
}

func GetLdap() (*Ldap, error) {
	var ldap Ldap
	err := db.Where("id = ?", 1).First(&ldap).Error
	if err == gorm.ErrRecordNotFound {
		return &ldap, nil
	}

	return &ldap, err
}

func UpdateLdap(id uint, data map[string]interface{}) error {
	return db.Model(&Ldap{}).Where("id = ?", id).Update(data).Error
}
