package services

import (
	"math"
	"pismo/models"
	"pismo/repositories"
	"time"
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
	transaction.EventDate = time.Now()
	if models.Payment == transaction.OperationTypeID {
		transactions, err := s.repo.GetDischargeTransactions(transaction.AccountID)
		if err != nil {
			return nil, err
		}
		balance := transaction.Balance
		balanceAux := 0.0
		for i, _ := range transactions {
			if balance > 0 {
				balanceAux = balance + transactions[i].Balance
				if balance > math.Abs(transactions[i].Balance) {
					transactions[i].Balance = 0
					balance = balanceAux
				} else {
					transactions[i].Balance = balanceAux
					balance = 0
				}
				err = s.repo.UpdateBalance(&transactions[i])
				if err != nil {
					return nil, err
				}
			} else {
				break
			}
		}
		transaction.Balance = balance
	}
	return s.repo.Create(transaction)
}

func (s *transactionsService) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.repo.Update(transaction)
}

func (s *transactionsService) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}
