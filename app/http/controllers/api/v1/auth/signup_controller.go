// Package auth: handle user authentication
package auth

import (
	"fmt"
	v1 "goapihub/app/http/controllers/api/v1"
	"goapihub/app/models/user"
	"goapihub/app/requests"
	"goapihub/pkg/response"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist: judge whether phone number is registered
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	
	// init request 
	request := requests.SignupPhoneExistRequest{}


	if ok := requests.ValidForm(&request, c, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// select database, and return response
	response.JSON(c, gin.H {
		"exist": user.IsPhoneExist(request.Phone),
	})
}



// IsEmailExist: judge whether email is registered
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	
	// init request 
	request := requests.SignupEmailExistRequest{}

	if ok := requests.ValidForm(&request, c, requests.ValidateSignupEmailExist); !ok {
		return
	}

	// select database, and return response
	response.JSON(c, gin.H {
		"exist": user.IsEmailExist(request.Email),
	})
}

// SignupUsingPhone: use phone and captcha to register
func(sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 1.验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.ValidForm(&request, c, requests.SignupUsingPhone); !ok {
		return
	}

	// 2.验证成功，创建数据
	_user := user.User{
		Name: request.Name,
		Phone: request.Phone,
		Password: request.Password,
	}
	err := _user.Create()
	if err != nil {
		// 创建失败
		fmt.Println("创建用户失败:", err)
	} else {
		fmt.Println("用户创建成功")
	}

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}