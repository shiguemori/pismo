package services

import (
	"pismo/models"
	"pismo/repositories"
)

type AccountsService interface {
	ListAllAccount() ([]models.Account, error)
	GetAccountByID(id uint) (*models.Account, error)
	CreateAccount(account *models.Account) (*models.Account, error)
	UpdateAccount(account *models.Account) (*models.Account, error)
	DeleteAccount(id uint) error
}

type accountsService struct {
	repo repositories.AccountsRepository
}

func NewAccountsService(accountRepo repositories.AccountsRepository) AccountsService {
	return &accountsService{
		repo: accountRepo,
	}
}

func (s *accountsService) ListAllAccount() ([]models.Account, error) {
	return s.repo.ListAll()
}

func (s *accountsService) GetAccountByID(id uint) (*models.Account, error) {
	return s.repo.GetById(id)
}

func (s *accountsService) CreateAccount(account *models.Account) (*models.Account, error) {
	return s.repo.Create(account)
}

func (s *accountsService) UpdateAccount(account *models.Account) (*models.Account, error) {
	return s.repo.Update(account)
}

func (s *accountsService) DeleteAccount(id uint) error {
	return s.repo.Delete(id)
}
