package main

import (
	"games_api/service"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}


func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"teste": "Pagina Inicial",
		})
	})

	service.LoginService(r)
	service.JogoService(r)

	r.Run(":8080")
}