package auth

import (
	v1 "goapihub/app/http/controllers/api/v1"
	"goapihub/app/models/user"
	"goapihub/app/requests"
	"goapihub/pkg/response"

	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	//1. 验证表单
	request := requests.ResetByPhoneRequest{}
	if ok := requests.ValidForm(&request, c, requests.ResetByPhone); !ok {
		return 
	}

	//2. update password
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}