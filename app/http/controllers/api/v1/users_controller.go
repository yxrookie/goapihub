package v1

import (
    "goapihub/pkg/auth"
    "goapihub/pkg/response"

    "github.com/gin-gonic/gin"
)

type UsersController struct {
    BaseAPIController
}


// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
    userModel := auth.CurrentUser(c)
    response.Data(c, userModel)
}