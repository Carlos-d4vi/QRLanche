package model

type OrderItem struct {
    OrderID    int64 `json:"order_id"`    // ID do pedido
    MenuItemID int64 `json:"menu_item_id"` // ID do item no menu
}