// register routes
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// register web page route
func RegisterAPIRoutes(r *gin.Engine) {
	// register v1 version route group
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "Api!",
			})
		})
	}

}