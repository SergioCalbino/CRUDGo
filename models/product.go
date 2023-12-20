package models

import (
	"gormjwt/db"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	// ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	UserID      int64      `json:"user_id"`
	Categories  []Category `gorm:"many2many:product_categories"`
}

type Products []Product

func MigrateProduct() {
	db.Database.AutoMigrate(Product{})
}
