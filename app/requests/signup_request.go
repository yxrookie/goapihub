// Package requests: handle request data and form validtion
package requests

import (
	"goapihub/app/requests/validators"

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
	
	return validate(data, rules, messages)
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
	
	return validate(data, rules, messages)
}

// SignupUsingPhoneRequest: the request information by using phone
type SignupUsingPhoneRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Name string `valid:"name" json:"name"`
	Password string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
} 

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11", "not_exists:users,phone"},
		"name": []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password": []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData {
		"phone": []string{
			"required:手机号为必填项,参数名称为 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"name":[]string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许英文和数字",
			"between:用户名长度需要 3~20 之间",
		},
		"password": []string {
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"verify_code":[]string {
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

