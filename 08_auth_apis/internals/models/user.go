package models

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
