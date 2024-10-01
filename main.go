package main

import (
	"net/http"

	"github.com/ayushkkkkkkk/mines/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(
			http.StatusOK, gin.H{
				"message": "working",
			},
		)
	})
	r.Run()
}
