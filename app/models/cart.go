package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId      uint
	ProductName string
	BuyingPrice string
	Quantity    int
	isBought    bool
}
