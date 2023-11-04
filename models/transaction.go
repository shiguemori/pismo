package models

import (
	"time"
)

type Transaction struct {
	Id              uint          `gorm:"primaryKey" json:"id,omitempty"`
	AccountID       uint          `gorm:"not null" json:"accountID,omitempty"`
	OperationTypeID uint          `gorm:"not null" json:"operationTypeID,omitempty"`
	Amount          float64       `gorm:"not null" json:"amount,omitempty"`
	EventDate       time.Time     `gorm:"not null" json:"eventDate"`
	Account         Account       `gorm:"foreignKey:AccountID" json:"account"`
	OperationType   OperationType `gorm:"foreignKey:OperationTypeID" json:"operationType"`
}
