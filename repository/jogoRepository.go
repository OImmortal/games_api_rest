package repository

import (
	"errors"
	"games_api/models"
	"strings"
)

var (
	db        = make(map[int]models.Jogo)
	idCounter = 0
)

func validarJogo(jogo models.Jogo) error {
	if strings.TrimSpace(jogo.Nome) == "" {
		return errors.New("Nome é obrigatório")
	}

	if strings.TrimSpace(jogo.Tipo) == "" {
		return errors.New("Tipo é obrigatório")
	}

	if jogo.Nota < 0 || jogo.Nota > 10 {
		return errors.New("Nota deve estar entre 0 e 10")
	}

	return nil
}

func GetAllJogos() ([]models.Jogo, error) {
	jogos := make([]models.Jogo, 0, len(db))

	for _, jogo := range db {
		jogos = append(jogos, jogo)
	}

	return jogos, nil
}

func GetJogoById(id int) (models.Jogo, error) {
	if id <= 0 {
		return models.Jogo{}, errors.New("ID inválido")
	}

	jogo, ok := db[id]
	if !ok {
		return models.Jogo{}, errors.New("Jogo não encontrado")
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
		return models.Jogo{}, errors.New("ID inválido")
	}

	if _, ok := db[id]; !ok {
		return models.Jogo{}, errors.New("Jogo não encontrado")
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
		return errors.New("ID inválido")
	}

	if _, ok := db[id]; !ok {
		return errors.New("Jogo não encontrado")
	}

	delete(db, id)

	return nil
}
