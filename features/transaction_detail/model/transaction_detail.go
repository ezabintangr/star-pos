package transactionDetailModel

import (
	transactionModel "star-pos/features/transaction/model"
	"time"
)

type TransactionDetail struct {
	ID            string                       `gorm:"type:char(36);primarykey" json:"id" form:"id"`
	TransactionID string                       `gorm:"type:char(36)" json:"transaction_id" form:"transaction_id"`
	ProductName   string                       `json:"product_name" form:"product_name"`
	Quantity      int64                        `json:"quantity" form:"quantity"`
	PricePerUnit  float64                      `json:"price_per_unit" form:"price_per_unit"`
	CreatedAt     time.Time                    `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time                    `json:"updated_at" form:"updated_at"`
	Transaction   transactionModel.Transaction `gorm:"ForeignKey:TransactionID"`
}
