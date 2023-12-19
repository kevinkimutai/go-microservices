package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Owner    string  `gorm:"not null"`
	Balance  float64 `gorm:"not null;default:0"`
	Currency string  `gorm:"type:ENUM('usd', 'kes', 'eur','yen');default:'usd'"`
	Entries  []Entries
}
