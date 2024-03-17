package v1

import (
	"goapihub/app/models/user"
	"goapihub/app/requests"

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

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
    request := requests.PaginationRequest{}
    if ok := requests.ValidGetForm(&request, c, requests.Pagination); !ok {
        return
    }
    data, pager := user.Paginate(c, 10)
    response.JSON(c, gin.H{
        "data":  data,
        "pager": pager,
    })
}