package requests

import (
	"goapihub/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// init configuration 
	opts := govalidator.Options {
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}
	//start validate
	return govalidator.New(opts).ValidateStruct()
}

// validation function type
type validFunc func(data any, c *gin.Context) map[string][]string

func ValidForm(obj any, c *gin.Context,handler validFunc) bool {
	// parse JSON request
	if err := c.ShouldBindJSON(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式")
		
		// end request
		return false
	}

	// form validation
	errs := handler(obj, c)

	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}