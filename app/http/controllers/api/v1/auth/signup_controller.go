// Package auth: handle user authentication
package auth

import (
	"fmt"
	v1 "goapihub/app/http/controllers/api/v1"
	"goapihub/app/models/user"
	"goapihub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist: judge whether phone number is registered
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	
	// init request 
	request := requests.SignupPhoneExistRequest{}


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

	// form validation
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// select database, and return response
	c.JSON(http.StatusOK, gin.H {
		"exist": user.IsPhoneExist(request.Phone),
	})
}