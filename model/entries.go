package model

import "gorm.io/gorm"

type Entries struct {
	gorm.Model
	AccountID int     `json:"account_id" gorm:"default:not null"`
	Account   Account `gorm:"foreignKey:AccountID"`
	Amount    float64 `gorm:"not null"`
}
