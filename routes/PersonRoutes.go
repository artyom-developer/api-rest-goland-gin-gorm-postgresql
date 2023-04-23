package routes

import (
	"github.com/gin-gonic/gin"
	"tutofox.com/m/controllers"
)

func PersonRouter(router *gin.Engine) {

	routes := router.Group("api/v1/persons")
	routes.POST("", controllers.PersonCreate)
	routes.GET("", controllers.PersonGet)
	routes.GET("/:id", controllers.PersonGetById)
	routes.PUT("/:id", controllers.PersonUpdate)
	routes.DELETE("/:id", controllers.PersonDelete)

}
