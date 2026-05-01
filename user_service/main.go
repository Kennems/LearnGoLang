package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/register", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "User registered"})
	})

	r.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "User logged in"})
	})

	r.GET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "User info"})
	})

	r.Run(":8080")
}
