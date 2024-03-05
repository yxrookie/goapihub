// Package requests: handle request data and form validtion
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// custom validation rules
	rules := govalidator.MapData {
		"phone": []string{"required", "digits:11"},
	}
	// custom error message
	messages := govalidator.MapData {
		"phone": []string {
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	// init configuration 
	opts := govalidator.Options {
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}
	//start validate
	return govalidator.New(opts).ValidateStruct()
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	// custom validation rules
	rules := govalidator.MapData {
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	// custom error message
	messages := govalidator.MapData {
		"email": []string {
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	// init configuration 
	opts := govalidator.Options {
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}
	//start validate
	return govalidator.New(opts).ValidateStruct()
}