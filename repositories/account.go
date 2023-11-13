package repositories

import (
	"pismo/models"

	"gorm.io/gorm"
)

type AccountsRepository interface {
	ListAll() ([]models.Account, error)
	GetById(id uint) (*models.Account, error)
	Create(account *models.Account) (*models.Account, error)
	Update(account *models.Account) (*models.Account, error)
	Delete(id uint) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountsRepository(db *gorm.DB) AccountsRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) ListAll() ([]models.Account, error) {
	var accounts []models.Account
	result := r.db.Order("id").Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func (r *accountRepository) GetById(id uint) (*models.Account, error) {
	var account models.Account
	result := r.db.First(&account, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (r *accountRepository) Create(account *models.Account) (*models.Account, error) {
	if err := r.db.Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *accountRepository) Update(account *models.Account) (*models.Account, error) {
	if err := r.db.Updates(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *accountRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Account{}, id)
	return result.Error
}
