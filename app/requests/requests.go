package requests

import (
	"fmt"
	"net/http"

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
		// parse failure, return 422 status code and failure information
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		// end request
		return false
	}

	// form validation
	errs := handler(obj, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return false
	}
	return true
}