package model

type User struct {
	ID        int         `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Age       int         `json:"age"`
	Country   string      `json:"country"`
	TodoItems []*TodoItem `json:"todo_items,omitempty"`
}
