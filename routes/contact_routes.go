package routes

import (
	"phonebook/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, contactController *controllers.ContactController) {
	group := router.Group("/contacts")
	group.POST("/", contactController.CreateContact)
	group.GET("", contactController.GetContacts)
	group.GET("/search", contactController.SearchContacts)
	group.PUT("/:id", contactController.UpdateContact)
	group.DELETE("/:id", contactController.DeleteContact)
}
