package model

type TodoItem struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
