// register routes
package routes

import (
	"goapihub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// register web page route
func RegisterAPIRoutes(r *gin.Engine) {
	// register v1 version route group
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// judge whether phone number is registered
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}

}