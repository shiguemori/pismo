package controllers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"pismo/models"
	"pismo/services"
	"pismo/utils"
	"strconv"
)

type TransactionsController interface {
	Get(c *gin.Context)
	ListAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type transactionsController struct {
	TransactionService services.TransactionsService
}

func NewTransactionsController(service services.TransactionsService) TransactionsController {
	return &transactionsController{
		TransactionService: service,
	}
}

// ListAll Transaction godoc
// @Summary List all transactions
// @Description get transactions
// @Tags transactions
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Transaction
// @Failure 404 {object} utils.Response
// @Router /transactions [get]
func (ac *transactionsController) ListAll(c *gin.Context) {
	transaction, err := ac.TransactionService.ListAllTransaction()
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	if len(transaction) == 0 {
		slog.Error(utils.NotFound)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// Get Transaction godoc
// @Summary Get an transaction
// @Description get transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path uint true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 404 {object} utils.Response
// @Router /transactions/{id} [get]
func (ac *transactionsController) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		slog.Error(utils.IDCannotBeEmpty)
		c.JSON(http.StatusBadRequest, utils.Response{Message: utils.IDCannotBeEmpty})
		return
	}

	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error(utils.ErrorConvertingIDToUint, err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: utils.ErrorConvertingIDToUint})
		return
	}

	transaction, err := ac.TransactionService.GetTransactionByID(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// Create Transaction godoc
// @Summary Create an transaction
// @Description create new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body models.Transaction true "Create Transaction"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} utils.Response
// @Router /transactions [post]
func (ac *transactionsController) Create(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		return
	}

	createdTransaction, err := ac.TransactionService.CreateTransaction(&transaction)
	if err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTransaction)
}

// Update Transaction godoc
// @Summary Update an transaction
// @Description update transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path uint true "Transaction ID"
// @Param transaction body models.Transaction true "Update Transaction"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /transactions/{id} [put]
func (ac *transactionsController) Update(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		slog.Error(utils.ErrorBindingJSON, err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		return
	}

	id := c.Param("id")
	if id == "" {
		slog.Error(utils.IDCannotBeEmpty)
		c.JSON(http.StatusBadRequest, utils.Response{Message: utils.IDCannotBeEmpty})
		return
	}

	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error(utils.ErrorConvertingIDToUint, err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: utils.ErrorConvertingIDToUint})
		return
	}

	transaction.Id = uint(num)

	_, err = ac.TransactionService.GetTransactionByID(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	updatedTransaction, err := ac.TransactionService.UpdateTransaction(&transaction)
	if err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTransaction)
}

// Delete Transaction godoc
// @Summary Delete an transaction
// @Description delete transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path uint true "Transaction ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /transactions/{id} [delete]
func (ac *transactionsController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		slog.Error(utils.IDCannotBeEmpty)
		c.JSON(http.StatusBadRequest, utils.Response{Message: utils.IDCannotBeEmpty})
		return
	}

	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		slog.Error(utils.ErrorConvertingIDToUint, err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: utils.ErrorConvertingIDToUint})
		return
	}

	err = ac.TransactionService.DeleteTransaction(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, utils.DeletedSuccessfully)
}
