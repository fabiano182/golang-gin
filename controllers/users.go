package controllers

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from server :)",
	})
	return
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": "[fabiano]",
	})
}
