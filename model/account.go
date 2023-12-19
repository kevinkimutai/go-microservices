package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Owner    string  `gorm:"not null"`
	Balance  float64 `gorm:"default:0"`
	Currency string  `gorm:"type:ENUM('usd', 'kes', 'eur','yen');default:'usd'"`
	Entries  []Entries
}
