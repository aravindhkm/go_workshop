package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GinExampleOne() {
	router := gin.Default()

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"response": "welcome to the backend",
		})
	})

	router.Run(":8080")

	fmt.Println("")
}