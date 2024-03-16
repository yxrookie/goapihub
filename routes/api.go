// register routes
package routes

import (
	controllers "goapihub/app/http/controllers/api/v1"
	"goapihub/app/http/controllers/api/v1/auth"
	"goapihub/app/http/middlewares"

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
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
            // 使用手机号，短信验证码进行登录
            authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
            authGroup.POST("/login/using-password", lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			 // 重置密码
			 pwc := new(auth.PasswordController)
			 authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
			 authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)


			 uc := new(controllers.UsersController)

			 // 获取当前用户
			 v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		}
	}

}