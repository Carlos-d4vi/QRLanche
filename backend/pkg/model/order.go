package model

type Order struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	Itens      []int   `json:"itens"`        // Lista de IDs dos itens do menu
	Total      float64 `json:"total"`
	TableID    int64   `json:"table_id"`     // ID da mesa do restaurante
}