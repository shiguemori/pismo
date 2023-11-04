package models

type OperationType struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
