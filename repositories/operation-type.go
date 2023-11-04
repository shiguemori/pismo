package repositories

import (
	"gorm.io/gorm"
	"pismo/models"
)

type OperationTypesRepository interface {
	ListAll() ([]models.OperationType, error)
	GetById(id uint) (*models.OperationType, error)
	Create(operationType *models.OperationType) (*models.OperationType, error)
	Update(operationType *models.OperationType) (*models.OperationType, error)
	Delete(id uint) error
}

type operationTypeRepository struct {
	db *gorm.DB
}

func NewOperationTypesRepository(db *gorm.DB) OperationTypesRepository {
	return &operationTypeRepository{
		db: db,
	}
}

func (r *operationTypeRepository) ListAll() ([]models.OperationType, error) {
	var operationTypes []models.OperationType
	result := r.db.Find(&operationTypes).Order("id")
	if result.Error != nil {
		return nil, result.Error
	}
	return operationTypes, nil
}

func (r *operationTypeRepository) GetById(id uint) (*models.OperationType, error) {
	var operationType models.OperationType
	result := r.db.First(&operationType, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &operationType, nil
}

func (r *operationTypeRepository) Create(operationType *models.OperationType) (*models.OperationType, error) {
	if err := r.db.Create(operationType).Error; err != nil {
		return nil, err
	}
	return operationType, nil
}

func (r *operationTypeRepository) Update(operationType *models.OperationType) (*models.OperationType, error) {
	if err := r.db.Updates(operationType).Error; err != nil {
		return nil, err
	}
	return operationType, nil
}

func (r *operationTypeRepository) Delete(id uint) error {
	result := r.db.Delete(&models.OperationType{}, id)
	return result.Error
}
