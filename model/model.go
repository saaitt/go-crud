package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	IsActive    bool `gorm:"index"`
	Description string
}

func (Product) TableName() string {
	return "products"
}
