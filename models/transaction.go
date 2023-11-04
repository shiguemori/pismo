package models

import (
	"time"
)

type Transaction struct {
	Id              uint          `gorm:"primaryKey" json:"id,omitempty"`
	AccountID       uint          `gorm:"not null" json:"account_id,omitempty"`
	OperationTypeID uint          `gorm:"not null" json:"operation_type_id,omitempty"`
	Amount          float64       `gorm:"not null" json:"amount,omitempty"`
	EventDate       time.Time     `gorm:"not null" json:"eventDate"`
	Account         Account       `gorm:"foreignKey:AccountID"`
	OperationType   OperationType `gorm:"foreignKey:OperationTypeID"`
}
