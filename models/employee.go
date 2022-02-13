package models

type Employee struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
