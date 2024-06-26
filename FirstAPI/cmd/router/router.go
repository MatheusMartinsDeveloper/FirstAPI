package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Initializer() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
