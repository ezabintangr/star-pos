package repository

import (
	"star-pos/features/transaction/repository"
	"time"
)

type TransactionDetail struct {
	ID            string `gorm:"type:char(36);primarykey"`
	TransactionID string `gorm:"type:char(36)"`
	ProductName   string
	Quantity      int64
	PricePerUnit  float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Transaction   repository.Transaction `gorm:"ForeignKey:TransactionID"`
}
