package models

const (
	Payment = 4
)

type OperationType struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
