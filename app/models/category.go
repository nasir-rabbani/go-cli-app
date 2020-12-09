package models

import (
	"gorm.io/gorm"
)

// Category - To map the product details
type Category struct {
	gorm.Model
	Name string
	// Products []Category
}
