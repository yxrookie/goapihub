// Package auth: handle user authentication
package auth

import (
	"fmt"
	v1 "goapihub/app/http/controllers/api/v1"
	"goapihub/app/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist: judge whether phone number is registered
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// request object
	type PhoneExistRequest struct {
		// use the struct tag, so parse the JSON phone data to struct field Phone
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	// parse JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		// parse failure, return 422 status code and failure information
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		// end request
		return
	}

	// select database, and return response
	c.JSON(http.StatusOK, gin.H {
		"exist": user.IsPhoneExist(request.Phone),
	})
}