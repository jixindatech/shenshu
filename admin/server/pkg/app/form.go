package app

import (
	"admin/core/rbac"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net"
	"regexp"
	"strings"
)

type IDForm struct {
	ID uint `uri:"id" validate:"required,gte=1"`
}

var validate *validator.Validate

var phoneReg *regexp.Regexp

func ValidatePhone(fl validator.FieldLevel) bool {
	var err error
	if phoneReg == nil {
		phoneReg, err = regexp.Compile("^1[3456789]\\d{9}$")
		if err != nil {
			return false
		}
	}

	phone := fl.Field().String()
	return phoneReg.Match([]byte(phone))
}

func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	if rbac.ROLES != nil && rbac.ROLES[role] != false {
		return true
	}

	return false
}

func ValidateLB(fl validator.FieldLevel) bool {
	lb := fl.Field().String()
	if lb != "" && (lb == "chash" || lb == "roundrobin") {
		return true
	}

	return false
}

func ValidateIP(fl validator.FieldLevel) bool {
	ip := fl.Field().String()

	ok := strings.Contains(ip, "/")
	if ok {
		_, _, err := net.ParseCIDR(ip)
		if err != nil {
			return false
		}
		return true
	}

	if net.ParseIP(ip) != nil {
		return true
	}

	return false
}

func SetupValidate() error {
	var err error

	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()

	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}

	err = validate.RegisterValidation("phone", ValidatePhone)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("role", ValidateRole)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("lb", ValidateLB)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("ip", ValidateIP)
	if err != nil {
		return err
	}

	return nil
}

func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return err
	}

	if validate != nil {
		return validate.Struct(form)
	}

	return fmt.Errorf("%s", "invalid validate")
}

func BindUriAndValid(c *gin.Context, form interface{}) error {
	err := c.BindUri(form)
	if err != nil {
		return err
	}

	if validate != nil {
		return validate.Struct(form)
	}

	return fmt.Errorf("%s", "invalid validate")
}
