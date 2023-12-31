package models

import (
	"time"
)

type Transaction struct {
	Id              uint          `gorm:"primaryKey" json:"id,omitempty"`
	AccountID       uint          `gorm:"not null" json:"account_id"`
	OperationTypeID uint          `gorm:"not null" json:"operation_type_id"`
	Amount          float64       `gorm:"not null" json:"amount"`
	Balance         float64       `gorm:"not null" json:"balance"`
	EventDate       time.Time     `gorm:"not null" json:"event_date,omitempty"`
	Account         Account       `gorm:"foreignKey:AccountID"`
	OperationType   OperationType `gorm:"foreignKey:OperationTypeID"`
}
