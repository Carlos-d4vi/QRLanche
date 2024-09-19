package model

type RestaurantTable struct {
	ID        int64 `json:"id"`
	Number    int64 `json:"number"`
	Available bool  `json:"available"`
}