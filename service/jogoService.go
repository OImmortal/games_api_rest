package service

import (
	"errors"
	"games_api/models"
	"games_api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	errInvalidID        = errors.New("ID inválido")
	errPageInvalida     = errors.New("O parâmetro page deve ser maior que zero.")
	errPageSizeInvalido = errors.New("O parâmetro pageSize deve ser maior que zero.")
)

type jogoInput struct {
	Nome           *string  `json:"nome" binding:"required"`
	Tipo           *string  `json:"tipo" binding:"required"`
	Nota           *int     `json:"nota" binding:"required"`
	Review         *string  `json:"review" binding:"required"`
	Preco          *float64 `json:"preco"`
	DataLancamento *string  `json:"dataLancamento"`
}

func JogoService(r *gin.Engine) {
	r.GET("/jogos", func(ctx *gin.Context) {
		page, pageSize, err := parsePaginacao(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		resultado, err := repository.ListJogos(repository.FiltroJogos{
			Nome:     ctx.Query("nome"),
			Genero:   ctx.Query("genero"),
			Page:     page,
			PageSize: pageSize,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao buscar jogos: " + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, resultado)
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
			respondJogoError(ctx, err)
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

		if input.Preco != nil {
			jogo.Preco = *input.Preco
		}

		if input.DataLancamento != nil {
			jogo.DataLancamento = *input.DataLancamento
		}

		jogoAtualizado, err := repository.PutJogoById(jogo, id)
		if err != nil {
			respondJogoError(ctx, err)
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
			respondJogoError(ctx, err)
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	})
}

func parsePaginacao(ctx *gin.Context) (int, int, error) {
	page := 1
	pageSize := 10

	if pageParam := ctx.Query("page"); pageParam != "" {
		valor, err := strconv.Atoi(pageParam)
		if err != nil || valor <= 0 {
			return 0, 0, errPageInvalida
		}
		page = valor
	}

	if pageSizeParam := ctx.Query("pageSize"); pageSizeParam != "" {
		valor, err := strconv.Atoi(pageSizeParam)
		if err != nil || valor <= 0 {
			return 0, 0, errPageSizeInvalido
		}
		pageSize = valor
	}

	return page, pageSize, nil
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

func respondJogoError(ctx *gin.Context, err error) {
	status := http.StatusBadRequest
	if errors.Is(err, repository.ErrJogoNaoEncontrado) {
		status = http.StatusNotFound
	}

	ctx.JSON(status, gin.H{
		"error": err.Error(),
	})
}
