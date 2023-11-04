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

type MockOperationTypesService struct {
	mock.Mock
}

func (m *MockOperationTypesService) ListAllOperationType() ([]models.OperationType, error) {
	args := m.Called()
	return args.Get(0).([]models.OperationType), args.Error(1)
}

func (m *MockOperationTypesService) GetOperationTypeByID(id uint) (*models.OperationType, error) {
	args := m.Called(id)
	return args.Get(0).(*models.OperationType), args.Error(1)
}

func (m *MockOperationTypesService) CreateOperationType(operationType *models.OperationType) (*models.OperationType, error) {
	args := m.Called(operationType)
	return args.Get(0).(*models.OperationType), args.Error(1)
}

func (m *MockOperationTypesService) UpdateOperationType(operationType *models.OperationType) (*models.OperationType, error) {
	args := m.Called(operationType)
	return args.Get(0).(*models.OperationType), args.Error(1)
}

func (m *MockOperationTypesService) DeleteOperationType(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateOperationType(t *testing.T) {
	mockService := new(MockOperationTypesService)
	operationType := &models.OperationType{
		Id:          1,
		Description: "test",
	}

	mockService.On("CreateOperationType", operationType).Return(operationType, nil)

	controller := NewOperationTypeController(mockService)

	type OperationTypeCreateRequest struct {
		ID          uint   `json:"id"`
		Description string `json:"description"`
	}

	operationTypeRequestBody := OperationTypeCreateRequest{
		ID:          1,
		Description: "test",
	}

	requestBodyBytes, err := json.Marshal(operationTypeRequestBody)
	if err != nil {
		t.Fatal(err)
	}

	requestBodyBuffer := bytes.NewBuffer(requestBodyBytes)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("POST", "/operation-types", requestBodyBuffer)

	controller.Create(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestListAllOperationTypes(t *testing.T) {
	mockService := new(MockOperationTypesService)
	expectedOperationTypes := []models.OperationType{{Id: 1, Description: "Test Type"}}
	mockService.On("ListAllOperationType").Return(expectedOperationTypes, nil)

	controller := NewOperationTypeController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("GET", "/operation-types", nil)

	controller.ListAll(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
func TestGetByID(t *testing.T) {
	mockService := new(MockAccountsService)
	account := &models.Account{Id: 1, DocumentNumber: "1234567890"}
	mockService.On("GetAccountByID", uint(1)).Return(account, nil)

	controller := NewAccountsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("GET", "/accounts/1", nil)
	context.Params = gin.Params{{Key: "id", Value: "1"}}

	controller.Get(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUpdateOperationType(t *testing.T) {
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
	context.Params = gin.Params{{Key: "id", Value: "1"}}

	controller.Update(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteOperationType(t *testing.T) {
	mockService := new(MockAccountsService)
	mockService.On("DeleteAccount", uint(1)).Return(nil)

	controller := NewAccountsController(mockService)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request, _ = http.NewRequest("DELETE", "/accounts/1", nil)
	context.Params = gin.Params{{Key: "id", Value: "1"}}

	controller.Delete(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
