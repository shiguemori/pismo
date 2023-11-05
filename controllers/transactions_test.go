package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pismo/models"
	"pismo/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionsService struct {
	mock.Mock
}

func (m *MockTransactionsService) GetTransactionByID(id uint) (*models.Transaction, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTransactionsService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	args := m.Called(transaction)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTransactionsService) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	args := m.Called(transaction)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTransactionsService) DeleteTransaction(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTransactionsService) ListAllTransaction() ([]models.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func TestListAllTransactions(t *testing.T) {
	mockService := new(MockTransactionsService)
	expectedTransactions := []models.Transaction{
		{Id: 1, Amount: 100.00, AccountID: 1, OperationTypeID: 1},
		{Id: 2, Amount: 200.00, AccountID: 2, OperationTypeID: 2},
	}
	mockService.On("ListAllTransaction").Return(expectedTransactions, nil)

	controller := NewTransactionsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("GET", "/transactions", nil)

	controller.ListAll(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetTransaction(t *testing.T) {
	mockService := new(MockTransactionsService)
	transaction := &models.Transaction{Id: 1, Amount: 100.00, AccountID: 1, OperationTypeID: 1}
	mockService.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	controller := NewTransactionsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("GET", "/transactions/1", nil)
	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	controller.Get(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, utils.ToJSON(transaction), recorder.Body.String())
}

func TestCreateTransaction(t *testing.T) {
	mockService := new(MockTransactionsService)
	transaction := &models.Transaction{
		Amount:          100.00,
		AccountID:       1,
		OperationTypeID: 4,
	}

	mockService.On("CreateTransaction", mock.AnythingOfType("*models.Transaction")).Return(transaction, nil)

	controller := NewTransactionsController(mockService)

	transactionRequestBody := models.Transaction{
		Amount:          100.00,
		AccountID:       1,
		OperationTypeID: 4,
	}

	requestBodyBytes, err := json.Marshal(transactionRequestBody)
	if err != nil {
		t.Fatal(err)
	}

	requestBodyBuffer := bytes.NewBuffer(requestBodyBytes)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("POST", "/transactions", requestBodyBuffer)

	controller.Create(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.JSONEq(t, utils.ToJSON(transaction), recorder.Body.String())
}

func TestUpdateTransaction(t *testing.T) {
	mockService := new(MockTransactionsService)
	existingTransaction := &models.Transaction{
		Id:              1,
		Amount:          100.00,
		AccountID:       1,
		OperationTypeID: 4,
	}
	updatedTransactionData := &models.Transaction{
		Amount:          150.00,
		AccountID:       1,
		OperationTypeID: 3,
	}
	updatedTransaction := &models.Transaction{
		Id:              1,
		Amount:          150.00,
		AccountID:       1,
		OperationTypeID: 3,
	}

	mockService.On("GetTransactionByID", uint(1)).Return(existingTransaction, nil)
	mockService.On("UpdateTransaction", mock.AnythingOfType("*models.Transaction")).Return(updatedTransaction, nil)

	controller := NewTransactionsController(mockService)

	requestBodyBytes, err := json.Marshal(updatedTransactionData)
	if err != nil {
		t.Fatal(err)
	}

	requestBodyBuffer := bytes.NewBuffer(requestBodyBytes)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("PUT", "/transactions/1", requestBodyBuffer)
	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	controller.Update(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, utils.ToJSON(updatedTransaction), recorder.Body.String())
}

func TestDeleteTransaction(t *testing.T) {
	mockService := new(MockTransactionsService)
	mockService.On("DeleteTransaction", uint(1)).Return(nil)

	controller := NewTransactionsController(mockService)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("DELETE", "/transactions/1", nil)
	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	controller.Delete(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "\"deleted successfully\"", recorder.Body.String())
}
