package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// TableName :
// func (User) TableName() string {
// 	return "users"
// }
