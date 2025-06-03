package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Products []Product `gorm:"constraint:OnDelete:CASCADE;"`
}
