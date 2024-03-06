// Package auth: handle user authentication
package auth

import (
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


	if ok := requests.ValidForm(&request, c, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// select database, and return response
	c.JSON(http.StatusOK, gin.H {
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
	c.JSON(http.StatusOK, gin.H {
		"exist": user.IsEmailExist(request.Email),
	})
}