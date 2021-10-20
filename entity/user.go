package entity

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
