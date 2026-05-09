package service

import (
	"games_api/models"
	"games_api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginService(r *gin.Engine) {
	r.POST("/login", func(ctx *gin.Context) {

		var login models.Login

		if err := ctx.ShouldBindBodyWithJSON(&login); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": "Erro ao localizar login: " + err.Error(),
				},
			)
			return
		}

		token, err := repository.Login(login)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})


	})
}