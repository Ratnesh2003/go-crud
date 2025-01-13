package controllers

import (
	"ratnesh/crud/initializers"
	"ratnesh/crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Hello", Body: "World"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400);
	}

	c.JSON(200, gin.H{
		"message": post,
	})
}