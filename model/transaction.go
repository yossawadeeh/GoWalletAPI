package model

import (
	"time"

	"gorm.io/gorm"
)

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

	TransactionDate     time.Time
	Total               float32             `json:"total" binding:"required"`
	Note                string              `json:"note" binding:"required"`
	CategoryTransaction CategoryTransaction `gorm:"foreignKey:CategoryID"`
	CategoryID          uint                `json:"category_id"`
	User                User                `gorm:"foreignKey:UserID"`
	UserID              float64             `json:"user_id"`
}
