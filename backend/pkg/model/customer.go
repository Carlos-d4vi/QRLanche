package model

type Customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Cpf string `json:"cpf"`
}