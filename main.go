package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tutofox.com/m/configs"
	"tutofox.com/m/routes"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
