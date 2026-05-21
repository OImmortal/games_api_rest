package main

import (
	"games_api/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://gamesapirest-production.up.railway.app"}, // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},       // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},       // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                                // Headers exposed to the browser
		AllowCredentials: true,                                                      // Allow cookies/auth headers
		MaxAge:           12 * time.Hour,                                            // Cache preflight response
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"teste": "Pagina Inicial",
		})
	})

	service.LoginService(r)
	service.JogoService(r)

	r.Run(":8080")
}