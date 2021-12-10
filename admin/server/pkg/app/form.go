package app

import (
	"admin/core/rbac"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"regexp"
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

func SetupValidate() error {
	var err error

	validate = validator.New()
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
