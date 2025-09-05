package model

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact int64  `json:"phno"`
}
