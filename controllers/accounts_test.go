package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"pismo/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock service
type MockAccountsService struct {
	mock.Mock
}

func (m *MockAccountsService) ListAllAccount() ([]models.Account, error) {
	args := m.Called()
	return args.Get(0).([]models.Account), args.Error(1)
}

func (m *MockAccountsService) GetAccountByID(id uint) (*models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockAccountsService) UpdateAccount(account *models.Account) (*models.Account, error) {
	args := m.Called(account)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockAccountsService) DeleteAccount(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAccountsService) CreateAccount(account *models.Account) (*models.Account, error) {
	args := m.Called(account)
	return args.Get(0).(*models.Account), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	mockService := new(MockAccountsService)
	account := &models.Account{
		Id:             1,
		DocumentNumber: "test",
	}
	mockService.On("CreateAccount", account).Return(account, nil)

	controller := NewAccountsController(mockService)

	type AccountCreateRequest struct {
		Id             uint
		DocumentNumber string `json:"document_number"`
	}

	accountRequestBody := AccountCreateRequest{
		Id:             1,
		DocumentNumber: "test",
	}

	requestBodyBytes, err := json.Marshal(accountRequestBody)
	if err != nil {
		t.Fatal(err)
	}

	requestBodyBuffer := bytes.NewBuffer(requestBodyBytes)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("POST", "/accounts", requestBodyBuffer)

	controller.Create(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestListAllAccounts(t *testing.T) {
	mockService := new(MockAccountsService)
	expectedAccounts := []models.Account{{Id: 1, DocumentNumber: "1234567890"}}
	mockService.On("ListAllAccount").Return(expectedAccounts, nil)

	controller := NewAccountsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("GET", "/accounts", nil)

	controller.ListAll(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAccountByID(t *testing.T) {
	mockService := new(MockAccountsService)
	account := &models.Account{Id: 1, DocumentNumber: "1234567890"}
	mockService.On("GetAccountByID", uint(1)).Return(account, nil)

	controller := NewAccountsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("GET", "/accounts/1", nil)
	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	controller.Get(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateAccount(t *testing.T) {
	mockService := new(MockAccountsService)
	account := &models.Account{Id: 1, DocumentNumber: "1234567890"}
	updatedAccount := &models.Account{Id: 1, DocumentNumber: "0987654321"}
	mockService.On("GetAccountByID", uint(1)).Return(account, nil)
	mockService.On("UpdateAccount", updatedAccount).Return(updatedAccount, nil)

	controller := NewAccountsController(mockService)

	updateRequestBody := models.Account{DocumentNumber: updatedAccount.DocumentNumber}
	requestBodyBytes, _ := json.Marshal(updateRequestBody)
	requestBodyBuffer := bytes.NewBuffer(requestBodyBytes)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("PATCH", "/accounts/1", requestBodyBuffer)
	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	controller.Update(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteAccount(t *testing.T) {
	mockService := new(MockAccountsService)
	mockService.On("DeleteAccount", uint(1)).Return(nil)

	controller := NewAccountsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("DELETE", "/accounts/1", nil)
	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}

	controller.Delete(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
