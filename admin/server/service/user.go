package service

import (
	"admin/core/log"
	"admin/server/models"
	"admin/server/util"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"go.uber.org/zap"
)

type User struct {
	ID uint

	Username    string
	DisplayName string
	LoginType   string
	Password    string
	Salt        string
	Email       string
	Phone       string
	Status      int
	Role        string

	Remark string

	Page     int
	PageSize int
}

func (u *User) Save() error {
	data := map[string]interface{}{
		"username":    u.Username,
		"displayName": u.DisplayName,
		"loginType":   u.LoginType,
		"email":       u.Email,
		"phone":       u.Phone,
		"status":      u.Status,
		"role":        u.Role,
		"remark":      u.Remark,
	}

	if u.ID > 0 {
		if len(u.Password) > 0 {
			salt, password := util.GetSaltAndEncodedPassword(u.Password)
			data["salt"] = salt
			data["password"] = password
		}

		// id:1 admin only support standard and always enabled
		if u.ID == 1 {
			delete(data, "loginType")
			data["status"] = 1
			data["role"] = "admin"
		}

		return models.UpdateUser(u.ID, data)
	}

	return models.AddUser(data)
}

func (u *User) UpdatePassword() error {
	data := make(map[string]interface{})

	user, err := models.GetUser(u.ID)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return fmt.Errorf("%s", "invalid user id")
	}

	passwordStr := util.GeneratePassword(16)
	fmt.Println(string(passwordStr))
	salt, password := util.GetSaltAndEncodedPassword(string(passwordStr))
	data["salt"] = salt
	data["password"] = password

	err = models.UpdateUser(u.ID, data)
	if err != nil {
		return err
	}

	return SendMail(user.Email, "重置密码", passwordStr)
}

func (u *User) UpdateUserInfo() error {
	data := make(map[string]interface{})
	if len(u.DisplayName) > 0 {
		data["displayName"] = u.DisplayName
		data["email"] = u.Email
		data["phone"] = u.Phone
	}

	if len(u.Password) > 0 {
		salt, password := util.GetSaltAndEncodedPassword(u.Password)
		data["salt"] = salt
		data["password"] = password
	}

	return models.UpdateUser(u.ID, data)
}

func (u *User) Delete() error {
	if u.ID == 1 {
		return fmt.Errorf("%s", "invalid user id")
	}
	return models.DeleteUser(u.ID)
}

func (u *User) Get() (*models.User, error) {
	return models.GetUser(u.ID)
}

func (u *User) GetList() ([]*models.User, uint, error) {
	var query = make(map[string]interface{})
	if len(u.Username) > 0 {
		query["username"] = u.Username
	}

	return models.GetUsers(query, u.Page, u.PageSize)
}

func (u *User) GetLoginUser(test bool) (*models.User, error) {
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("invalid username")
	}

	if user.Status == util.USER_DENY {
		return nil, fmt.Errorf("invalid user status")
	}

	if test {
		if util.VerifyRawPassword(u.Password, user.Password, user.Salt) {
			return user, nil
		}

		return nil, fmt.Errorf("invalid password")
	}

	if user.LoginType == "standard" {
		if util.VerifyRawPassword(u.Password, user.Password, user.Salt) {
			return user, nil
		}

		return nil, fmt.Errorf("invalid password")
	} else if user.LoginType == "ldap" {
		ldapConfig, err := getLdapConfig()
		if err != nil {
			log.Logger.Error("ldap", zap.String("err", err.Error()))
		}
		if ldapConfig != nil {
			conn, err := ldap.DialURL("ldap://" + ldapConfig.Host + fmt.Sprintf(":%d", ldapConfig.Port))
			if err != nil {
				return nil, err
			}

			if conn != nil {
				defer conn.Close()
				err = conn.Bind(ldapConfig.DN, ldapConfig.Password)
				if err != nil {
					return nil, err
				}
				filter := fmt.Sprintf("(%s=%s)", util.LDAP_USERNAME, u.Username)
				searchRequest := ldap.NewSearchRequest(
					ldapConfig.BaseDN,
					ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
					filter,
					[]string{"dn", "cn", "objectClass"},
					nil,
				)

				searchResult, err := conn.Search(searchRequest)
				if err != nil {
					return nil, err
				}
				if searchRequest == nil {
					return nil, fmt.Errorf("%s")
				}
				if searchResult != nil && len(searchResult.Entries) == 0 {
					return nil, fmt.Errorf("%s", "invalid username")
				}

				entry := searchResult.Entries[0]
				err = conn.Bind(entry.DN, u.Password)
				if err != nil {
					return nil, err
				}

				return user, nil
			}
		}
	}

	return nil, fmt.Errorf("login failed")
}

func SaveAdmin(id uint, username, role, password string) error {
	admin, err := models.GetUser(id)
	if err != nil {
		return err
	}

	if admin.ID > 0 {
		salt, encodedPassword := util.GetSaltAndEncodedPassword(password)
		data := make(map[string]interface{})
		data["salt"] = salt
		data["password"] = encodedPassword

		return models.UpdateUser(admin.ID, data)
	}

	data := map[string]interface{}{
		"username":    username,
		"displayName": "admin",
		"loginType":   "standard",
		"email":       "admin@admin.com",
		"phone":       "13200000000",
		"status":      1,
		"role":        role,
		"remark":      "administrator",
	}

	salt, encodedPassword := util.GetSaltAndEncodedPassword(password)
	data["salt"] = salt
	data["password"] = encodedPassword

	return models.AddUser(data)
}

func (u *User) SaveSelf() error {
	data := map[string]interface{}{
		"displayName": u.DisplayName,
		"email":       u.Email,
		"phone":       u.Phone,
	}

	if u.ID > 0 {
		if len(u.Password) > 0 {
			salt, password := util.GetSaltAndEncodedPassword(u.Password)
			data["salt"] = salt
			data["password"] = password
		}

		return models.UpdateUser(u.ID, data)
	}

	return fmt.Errorf("%s", "invalid user id")
}
