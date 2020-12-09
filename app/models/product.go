package models

import (
	"gorm.io/gorm"
)

// Product - To map the product details
type Product struct {
	gorm.Model
	Name       string
	Price      float32
	CategoryID uint
}
