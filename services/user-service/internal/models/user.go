package models

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}
