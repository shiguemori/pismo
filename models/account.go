package models

type Account struct {
	Id             uint   `gorm:"primaryKey" json:"id,omitempty"`
	DocumentNumber string `gorm:"type:varchar(255);unique_index" json:"document_number"`
}
