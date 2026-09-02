package models

type Jogo struct {
	ID             int     `json:"id"`
	Nome           string  `json:"nome"`
	Tipo           string  `json:"tipo"`
	Nota           int     `json:"nota"`
	Review         string  `json:"review"`
	Preco          float64 `json:"preco"`
	DataLancamento string  `json:"dataLancamento,omitempty"`
}

type ListaJogos struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	TotalItems int    `json:"totalItems"`
	TotalPages int    `json:"totalPages"`
	Items      []Jogo `json:"items"`
}
