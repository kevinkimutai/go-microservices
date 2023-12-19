package model

import "gorm.io/gorm"

type Entries struct {
	gorm.Model
	AccountID int     `json:"account_id"`
	Account   Account `gorm:"foreignKey:AccountID"`
	Amount    float64
}
