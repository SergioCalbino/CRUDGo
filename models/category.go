package models

import "gormjwt/db"

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func MigrateCategory() {
	db.Database.AutoMigrate(Category{})
}
