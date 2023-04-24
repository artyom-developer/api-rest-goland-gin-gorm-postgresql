package main

import (
	"tutofox.com/m/configs"
	"tutofox.com/m/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}
