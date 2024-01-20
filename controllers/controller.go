package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qadrina/go-crud-gin/initializers"
	"github.com/qadrina/go-crud-gin/models"
)

func PostsCreate(c *gin.Context) {
	// get data from req body
	// create a struct to store it
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// return them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsSingular(c *gin.Context) {
	// get the ID from url
	id := c.Param("id")
	// get the posts
	var post models.Post
	initializers.DB.First(&post, id)
	// return them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get the ID from the URL
	id := c.Param("id")

	// get the data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the updated post
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	// kalau sebaris, tidak perlu koma setelah field terakhir. kalau di-enter, perlu koma stlh field terakhir
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get the ID from the URL
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)
}
