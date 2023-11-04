package repositories

import (
	"gorm.io/gorm"
	"pismo/models"
)

type TransactionsRepository interface {
	ListAll() ([]models.Transaction, error)
	GetById(id uint) (*models.Transaction, error)
	Create(transaction *models.Transaction) (*models.Transaction, error)
	Update(transaction *models.Transaction) (*models.Transaction, error)
	Delete(id uint) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) TransactionsRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) ListAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Preload("Account").Preload("OperationType").Find(&transactions).Order("id")
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (r *transactionRepository) GetById(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	result := r.db.Preload("Account").Preload("OperationType").First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *transactionRepository) Create(transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	if err := r.db.Preload("Account").Preload("OperationType").Find(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) Update(transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Updates(transaction).Error; err != nil {
		return nil, err
	}

	if err := r.db.Preload("Account").Preload("OperationType").Find(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Transaction{}, id)
	return result.Error
}
