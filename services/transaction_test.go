package services

import (
	"github.com/stretchr/testify/assert"
	"pismo/models"
	"testing"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) ListAll() ([]models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m *TransactionRepositoryMock) GetById(id uint) (*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m *TransactionRepositoryMock) Update(transaction *models.Transaction) (*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m *TransactionRepositoryMock) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (m *TransactionRepositoryMock) GetDischargeTransactions(accountID uint) ([]models.Transaction, error) {
	args := m.Called(accountID)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) UpdateBalance(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Create(transaction *models.Transaction) (*models.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	mockRepo := new(TransactionRepositoryMock)
	service := transactionsService{repo: mockRepo}

	transactions := []models.Transaction{
		{
			AccountID:       2,
			OperationTypeID: 1,
			Amount:          -50,
			Balance:         -50,
		},
		{
			AccountID:       2,
			OperationTypeID: 2,
			Amount:          -23.5,
			Balance:         -23.5,
		},
		{
			AccountID:       2,
			OperationTypeID: 3,
			Amount:          -18.7,
			Balance:         -18.7,
		},
	}

	paymentTransactions := []models.Transaction{
		{
			AccountID:       2,
			OperationTypeID: models.Payment,
			Amount:          60,
			Balance:         60,
		},
		{
			AccountID:       2,
			OperationTypeID: models.Payment,
			Amount:          100,
			Balance:         100,
		},
	}

	expectedTransactions := []models.Transaction{
		{
			AccountID:       2,
			OperationTypeID: 1,
			Amount:          -50,
			Balance:         -50,
		},
		{
			AccountID:       2,
			OperationTypeID: 2,
			Amount:          -23.5,
			Balance:         -23.5,
		},
		{
			AccountID:       2,
			OperationTypeID: 3,
			Amount:          -18.7,
			Balance:         -18.7,
		},
	}

	expectedPaymentTransactions := []models.Transaction{
		{
			AccountID:       2,
			OperationTypeID: models.Payment,
			Amount:          60,
			Balance:         0,
		},
		{
			AccountID:       2,
			OperationTypeID: models.Payment,
			Amount:          100,
			Balance:         67.8,
		},
	}

	mockRepo.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&transactions[0], nil).Once()
	result, err := service.CreateTransaction(&transactions[0])
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedTransactions[0].Balance, result.Balance)

	mockRepo.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&transactions[1], nil).Once()
	result, err = service.CreateTransaction(&transactions[1])
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedTransactions[1].Balance, result.Balance)

	mockRepo.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&transactions[2], nil).Once()
	result, err = service.CreateTransaction(&transactions[2])
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedTransactions[2].Balance, result.Balance)

	mockRepo.On("GetDischargeTransactions", paymentTransactions[0].AccountID).Return(transactions, nil).Once()
	mockRepo.On("UpdateBalance", mock.AnythingOfType("*models.Transaction")).Return(nil).Twice()
	mockRepo.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&paymentTransactions[0], nil).Once()

	result, err = service.CreateTransaction(&paymentTransactions[0])
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedPaymentTransactions[0].Balance, result.Balance)

	mockRepo.On("GetDischargeTransactions", paymentTransactions[0].AccountID).Return(transactions[1:], nil).Once()
	mockRepo.On("UpdateBalance", mock.AnythingOfType("*models.Transaction")).Return(nil).Twice()
	mockRepo.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&paymentTransactions[1], nil).Once()

	result, err = service.CreateTransaction(&paymentTransactions[1])
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedPaymentTransactions[1].Balance, result.Balance)

	mockRepo.AssertExpectations(t)
}
