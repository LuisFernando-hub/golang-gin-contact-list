package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Api OK!",
		})
	})

	api := app.Group("/api")

	ContactRoutes(api)
}
