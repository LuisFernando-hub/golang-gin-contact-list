package routes

import (
	"gin-contact-list/handlers"

	"github.com/gin-gonic/gin"
)

func ContactRoutes(app *gin.RouterGroup) {
	router := app.Group("/contacts")

	router.GET("/", handlers.FetchAllContacts)
	router.POST("/", handlers.CreateContacts)
	router.GET("/:contactId", handlers.FetchContact)
	router.DELETE("/:contactId", handlers.DeleteContacts)
	router.PUT("/:contactId", handlers.UpdateContacts)
}
