package main

import (
	"tutofox.com/m/initializers"
	"tutofox.com/m/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Person{})
}
