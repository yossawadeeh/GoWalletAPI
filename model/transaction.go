package model

import "gorm.io/gorm"

type TypeTransaction struct {
	gorm.Model

	TypeName string
}

type CategoryTransaction struct {
	gorm.Model

	CategoryName    string
	TypeTransaction TypeTransaction `gorm:"foreignKey:TypeID"`
	TypeID          uint
}

type Transaction struct {
	gorm.Model

	TransactionDate     string              `json:"transaction_date" binding:"required"`
	Total               int                 `json:"total" binding:"required"`
	Note                string              `json:"note" binding:"required"`
	CategoryTransaction CategoryTransaction `gorm:"foreignKey:CategoryID"`
	CategoryID          uint
	User                User `gorm:"foreignKey:UserID"`
	UserID              uint
}
