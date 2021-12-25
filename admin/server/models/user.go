package models

import "github.com/jinzhu/gorm"

type User struct {
	Model
	Username    string `json:"username" gorm:"column:username;unique;comment:'用户'"`
	DisplayName string `json:"displayName" gorm:"column:display_name;comment:'昵称'"`
	LoginType   string `json:"loginType" gorm:"column:login_type;comment:'登陆类型'"`
	Password    string `json:"-" gorm:"column:password;not null;comment:'密码'"`
	Salt        string `json:"-" gorm:"column:salt;not null;comment:'盐'"`
	Email       string `json:"email" gorm:"column:email;not null;comment:'用户'"`
	Phone       string `json:"phone" gorm:"column:phone;comment:'手机'"`
	Status      int    `json:"status" gorm:"column:status;default:'0';comment:'状态: 1-normal,0-locked'"`
	Role        string `json:"role" gorm:"column:role;not null;comment:'角色'"`

	Remark string `json:"remark" gorm:"column:remark;comment:'备注'"`
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Username:    data["username"].(string),
		DisplayName: data["displayName"].(string),
		LoginType:   data["loginType"].(string),
		Password:    data["password"].(string),
		Salt:        data["salt"].(string),
		Email:       data["email"].(string),
		Phone:       data["phone"].(string),
		Status:      data["status"].(int),
		Role:        data["role"].(string),
		Remark:      data["remark"].(string),
	}
	return db.Create(&user).Error
}

func DeleteUser(id uint) error {
	return db.Where("id = ?", id).Delete(User{}).Error
}

func UpdateUser(id uint, data map[string]interface{}) error {
	return db.Model(&User{}).Where("id = ?", id).Update(data).Error
}

func GetUser(id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUsers(query map[string]interface{}, page int, pageSize int) ([]*User, uint, error) {
	var users []*User
	var count uint
	var err error

	pageNum := (page - 1) * pageSize

	var username string
	if query["username"] != nil {
		username = query["username"].(string)
	}

	if len(username) > 0 {
		username = "%" + username + "%"
		err = db.Where("username like ?", username).Offset(pageNum).Limit(pageSize).Find(&users).Count(&count).Error
	} else {
		err = db.Offset(pageNum).Limit(pageSize).Find(&users).Count(&count).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return users, count, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetBatchUser(ids []uint) ([]*User, error) {
	var users []*User
	err := db.Find(&users, ids).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}
