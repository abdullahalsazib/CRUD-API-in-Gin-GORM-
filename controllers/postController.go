package controllers

import (
	"github.com/abdullahalsazib/crud_app/initializers"
	"github.com/abdullahalsazib/crud_app/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
func PostsShow(c *gin.Context) {
	// get the post id
	id := c.Param("id")
	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id of the url
	id := c.Param("id")

	// Get the data of req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	// post.Body = body.Body
	// post.Title = body.Title
	// initializers.DB.Save(&body)
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}
func PostDelete(c *gin.Context) {
	// get the post id
	id := c.Param("id")
	// Get the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with them
	c.Status(200)
}
