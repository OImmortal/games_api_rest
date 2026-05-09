package repository

import (
	"errors"
	"games_api/models"

	"github.com/google/uuid"
)

func Login(infoLogin models.Login) (string, error) {

	email := infoLogin.Email
	password := infoLogin.Password

	if email == "" {
		return "", errors.New("E-mail é obrigatório")
	}

	if password == "" {
		return "", errors.New("Senha é obrigatório")
	}


	if email != "usuario@esoft.com" || password != "Abc123" {
		return "", errors.New("Usuário inválido")
	}

	u, _ := uuid.NewUUID()

	return u.String(), nil
	
}