package model

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	FromAccount   Account `gorm:"foreignKey:FromAccountID"`
	FromAccountID int     `json:"from_account_id"`
	ToAccount     Account `gorm:"foreignKey:FromAccountID"`
	ToAccountID   int     `json:"to_account_id"`
	Amount        float64
}
