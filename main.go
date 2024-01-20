package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qadrina/go-crud-gin/controllers"
	"github.com/qadrina/go-crud-gin/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/api/posts", controllers.PostsCreate) // time stamp 17.43
	r.PUT("/api/posts/:id", controllers.PostsUpdate)
	r.GET("/api/posts", controllers.PostsIndex)
	r.GET("/api/posts/:id", controllers.PostsSingular)
	r.DELETE("/api/posts/:id", controllers.PostsDelete)
	r.Run()
}
