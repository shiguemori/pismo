package services

import (
	"pismo/models"
	"pismo/repositories"
)

type TransactionsService interface {
	ListAllTransaction() ([]models.Transaction, error)
	GetTransactionByID(id uint) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error)
	DeleteTransaction(id uint) error
}

type transactionsService struct {
	repo repositories.TransactionsRepository
}

func NewTransactionsService(transactionRepo repositories.TransactionsRepository) TransactionsService {
	return &transactionsService{
		repo: transactionRepo,
	}
}

func (s *transactionsService) ListAllTransaction() ([]models.Transaction, error) {
	return s.repo.ListAll()
}

func (s *transactionsService) GetTransactionByID(id uint) (*models.Transaction, error) {
	return s.repo.GetById(id)
}

func (s *transactionsService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.repo.Create(transaction)
}

func (s *transactionsService) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.repo.Update(transaction)
}

func (s *transactionsService) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}
