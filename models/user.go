package models

import "gormjwt/db"

type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
	Products []Product `gorm:"foreignkey:UserID"`
}

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrateUser() {
	db.Database.AutoMigrate(User{})
}
