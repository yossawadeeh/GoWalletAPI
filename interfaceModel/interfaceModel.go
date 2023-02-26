package interfaceModel

import (
	"time"

	"gorm.io/gorm"
)

// Transactions
type TransactionInterface struct {
	gorm.Model

	TransactionDate time.Time
	Total           float32 `json:"total" binding:"required"`
	Note            string  `json:"note" binding:"required"`
	CategoryID      uint    `json:"category_id" binding:"required"`
}

type UpdateTransactionInterface struct {
	gorm.Model

	TransactionDate string  `json:"transaction_date"`
	Total           float32 `json:"total" binding:"required"`
	Note            string  `json:"note"`
	CategoryID      uint    `json:"category_id" binding:"required"`
}

// Reports
type TransactionByDateInterface struct {
	gorm.Model

	StartDate int64 `json:"start_date" binding:"required"`
	EndDate   int64 `json:"end_date" binding:"required"`
}
