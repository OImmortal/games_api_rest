package models

type Jogo struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Nota   int    `json:"nota"`
	Review string `json:"review"`
}