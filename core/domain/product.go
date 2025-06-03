package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string  `gorm:"not null"`
	Price  float64 `gorm:"not null"`
	UserID uint    `gorm:"not null"`
	User   User    `gorm:"foreignKey:UserID"`
}
