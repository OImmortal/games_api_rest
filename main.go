package main

import (
	"games_api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"teste": "Pagina Inicial",
		})
	})

	service.LoginService(r)
	service.JogoService(r)

	r.Run(":8080")
}