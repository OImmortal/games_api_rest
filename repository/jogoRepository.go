package repository

import (
	"errors"
	"games_api/models"
	"sort"
	"strings"
	"time"
)

var (
	db        = make(map[int]models.Jogo)
	idCounter = 0

	ErrJogoNaoEncontrado = errors.New("Jogo não encontrado.")
	ErrIDInvalido        = errors.New("ID inválido")
)

type FiltroJogos struct {
	Nome     string
	Genero   string
	Page     int
	PageSize int
}

func validarJogo(jogo models.Jogo) error {
	if strings.TrimSpace(jogo.Nome) == "" {
		return errors.New("O nome do jogo é obrigatório.")
	}

	if strings.TrimSpace(jogo.Tipo) == "" {
		return errors.New("O tipo do jogo é obrigatório.")
	}

	if jogo.Preco < 0 {
		return errors.New("O preço do jogo não pode ser negativo.")
	}

	if jogo.Nota < 0 || jogo.Nota > 10 {
		return errors.New("A nota do jogo deve estar entre 0 e 10.")
	}

	if err := validarDataLancamento(jogo.DataLancamento); err != nil {
		return err
	}

	return nil
}

func validarDataLancamento(data string) error {
	data = strings.TrimSpace(data)
	if data == "" {
		return nil
	}

	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"02/01/2006",
	}

	for _, layout := range layouts {
		if _, err := time.Parse(layout, data); err == nil {
			return nil
		}
	}

	return errors.New("A data de lançamento deve possuir um valor válido.")
}

func GetAllJogos() ([]models.Jogo, error) {
	jogos := make([]models.Jogo, 0, len(db))

	for _, jogo := range db {
		jogos = append(jogos, jogo)
	}

	sort.Slice(jogos, func(i, j int) bool {
		return jogos[i].ID < jogos[j].ID
	})

	return jogos, nil
}

func ListJogos(filtro FiltroJogos) (models.ListaJogos, error) {
	jogos, err := GetAllJogos()
	if err != nil {
		return models.ListaJogos{}, err
	}

	filtrados := filtrarJogos(jogos, filtro.Nome, filtro.Genero)
	return paginarJogos(filtrados, filtro.Page, filtro.PageSize), nil
}

func filtrarJogos(jogos []models.Jogo, nome, genero string) []models.Jogo {
	nomeFiltro := strings.ToLower(strings.TrimSpace(nome))
	generoFiltro := strings.ToLower(strings.TrimSpace(genero))

	filtrados := make([]models.Jogo, 0, len(jogos))
	for _, jogo := range jogos {
		if nomeFiltro != "" && !strings.Contains(strings.ToLower(jogo.Nome), nomeFiltro) {
			continue
		}

		if generoFiltro != "" && strings.ToLower(strings.TrimSpace(jogo.Tipo)) != generoFiltro {
			continue
		}

		filtrados = append(filtrados, jogo)
	}

	return filtrados
}

func paginarJogos(jogos []models.Jogo, page, pageSize int) models.ListaJogos {
	totalItems := len(jogos)
	totalPages := 0
	if pageSize > 0 {
		totalPages = (totalItems + pageSize - 1) / pageSize
	}

	inicio := (page - 1) * pageSize
	if inicio < 0 {
		inicio = 0
	}
	if inicio > totalItems {
		inicio = totalItems
	}

	fim := inicio + pageSize
	if fim > totalItems {
		fim = totalItems
	}

	items := jogos[inicio:fim]
	if items == nil {
		items = []models.Jogo{}
	}

	return models.ListaJogos{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Items:      items,
	}
}

func GetJogoById(id int) (models.Jogo, error) {
	if id <= 0 {
		return models.Jogo{}, ErrIDInvalido
	}

	jogo, ok := db[id]
	if !ok {
		return models.Jogo{}, ErrJogoNaoEncontrado
	}

	return jogo, nil
}

func CreateJogo(jogo models.Jogo) (models.Jogo, error) {
	if err := validarJogo(jogo); err != nil {
		return models.Jogo{}, err
	}

	idCounter++
	jogo.ID = idCounter
	db[jogo.ID] = jogo

	return jogo, nil
}

func PutJogoById(jogo models.Jogo, id int) (models.Jogo, error) {
	if id <= 0 {
		return models.Jogo{}, ErrIDInvalido
	}

	if _, ok := db[id]; !ok {
		return models.Jogo{}, ErrJogoNaoEncontrado
	}

	if err := validarJogo(jogo); err != nil {
		return models.Jogo{}, err
	}

	jogo.ID = id
	db[id] = jogo

	return jogo, nil
}

func DeleteJogo(id int) error {
	if id <= 0 {
		return ErrIDInvalido
	}

	if _, ok := db[id]; !ok {
		return ErrJogoNaoEncontrado
	}

	delete(db, id)

	return nil
}
