package services

import (
	"pismo/models"
	"pismo/repositories"
)

type OperationTypesService interface {
	ListAllOperationType() ([]models.OperationType, error)
	GetOperationTypeByID(id uint) (*models.OperationType, error)
	CreateOperationType(operationType *models.OperationType) (*models.OperationType, error)
	UpdateOperationType(operationType *models.OperationType) (*models.OperationType, error)
	DeleteOperationType(id uint) error
}

type operationTypesService struct {
	repo repositories.OperationTypesRepository
}

func NewOperationTypesService(operationTypeRepo repositories.OperationTypesRepository) OperationTypesService {
	return &operationTypesService{
		repo: operationTypeRepo,
	}
}

func (s *operationTypesService) ListAllOperationType() ([]models.OperationType, error) {
	return s.repo.ListAll()
}

func (s *operationTypesService) GetOperationTypeByID(id uint) (*models.OperationType, error) {
	return s.repo.GetById(id)
}

func (s *operationTypesService) CreateOperationType(operationType *models.OperationType) (*models.OperationType, error) {
	return s.repo.Create(operationType)
}

func (s *operationTypesService) UpdateOperationType(operationType *models.OperationType) (*models.OperationType, error) {
	return s.repo.Update(operationType)
}

func (s *operationTypesService) DeleteOperationType(id uint) error {
	return s.repo.Delete(id)
}
