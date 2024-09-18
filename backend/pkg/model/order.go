package model

type Order struct {
	ID int64
	Customer Customer
	Itens MenuItem
	Total float64
}