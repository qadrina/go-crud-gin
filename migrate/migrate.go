package main

import (
	"github.com/qadrina/go-crud-gin/initializers"
	"github.com/qadrina/go-crud-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

// time stamp 15:43
