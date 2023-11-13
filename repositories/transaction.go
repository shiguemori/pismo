package repositories

import (
	"gorm.io/gorm/clause"
	"pismo/models"

	"gorm.io/gorm"
)

type TransactionsRepository interface {
	ListAll() ([]models.Transaction, error)
	GetById(id uint) (*models.Transaction, error)
	GetDischargeTransactions(accountId uint) ([]models.Transaction, error)
	Create(transaction *models.Transaction) (*models.Transaction, error)
	Update(transaction *models.Transaction) (*models.Transaction, error)
	Delete(id uint) error
	UpdateBalance(transaction *models.Transaction) error
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
	result := r.db.Preload("Account").Preload("OperationType").Order("id").Find(&transactions)
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

func (r *transactionRepository) UpdateBalance(transaction *models.Transaction) error {
	if err := r.db.Model(transaction).UpdateColumn("balance", transaction.Balance).Error; err != nil {
		return err
	}

	if err := r.db.Preload("Account").Preload("OperationType").Find(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Transaction{}, id)
	return result.Error
}

func (r *transactionRepository) GetDischargeTransactions(accountId uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Where("account_id = ?", accountId).
		Where("operation_type_id != 4").
		Where("balance < ?", 0).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "event_date"}, Desc: false}).
		Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}
