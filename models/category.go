package models

import "gormjwt/db"

type Category struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Products []Product `gorm:"many2many:product_categories"`
}

func MigrateCategory() {
	db.Database.AutoMigrate(Category{})
}
