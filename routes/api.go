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
			// 发送验证码
            vcc := new(auth.VerifyCodeController)
            // 图片验证码，需要加限流
            authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
            authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
		}
	}

}