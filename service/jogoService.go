package service

import (
	"errors"
	"games_api/models"
	"games_api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var errInvalidID = errors.New("ID inválido")

type jogoInput struct {
	Nome   *string `json:"nome" binding:"required"`
	Tipo   *string `json:"tipo" binding:"required"`
	Nota   *int    `json:"nota" binding:"required"`
	Review *string `json:"review" binding:"required"`
}

func JogoService(r *gin.Engine) {
	r.GET("/jogos", func(ctx *gin.Context) {
		jogos, err := repository.GetAllJogos()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao buscar jogos: " + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, jogos)
	})

	r.GET("/jogos/:id", func(ctx *gin.Context) {
		id, err := parseID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		jogo, err := repository.GetJogoById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, jogo)
	})

	r.POST("/jogos", func(ctx *gin.Context) {
		var jogo models.Jogo

		if err := ctx.ShouldBindJSON(&jogo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "JSON inválido: " + err.Error(),
			})
			return
		}

		jogoCriado, err := repository.CreateJogo(jogo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, jogoCriado)
	})

	r.PUT("/jogos/:id", func(ctx *gin.Context) {
		id, err := parseID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var input jogoInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "JSON inválido: " + err.Error(),
			})
			return
		}

		jogo := models.Jogo{
			Nome:   *input.Nome,
			Tipo:   *input.Tipo,
			Nota:   *input.Nota,
			Review: *input.Review,
		}

		jogoAtualizado, err := repository.PutJogoById(jogo, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, jogoAtualizado)
	})

	r.DELETE("/jogos/:id", func(ctx *gin.Context) {
		id, err := parseID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := repository.DeleteJogo(id); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	})
}

func parseID(param string) (int, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, errInvalidID
	}

	if id <= 0 {
		return 0, errInvalidID
	}

	return id, nil
}
