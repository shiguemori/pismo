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

type AccountsController interface {
	GetAccount(c *gin.Context)
	ListAllAccount(c *gin.Context)
	CreateAccount(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type accountsController struct {
	AccountService services.AccountsService
}

func NewAccountsController(service services.AccountsService) AccountsController {
	return &accountsController{
		AccountService: service,
	}
}

// ListAllAccount godoc
// @Summary List all accounts
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Account
// @Failure 404 {object} utils.Response
// @Router /accounts [get]
func (ac *accountsController) ListAllAccount(c *gin.Context) {
	account, err := ac.AccountService.ListAllAccount()
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, account)
}

// GetAccount godoc
// @Summary Get an account
// @Description get account by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path uint true "Account ID"
// @Success 200 {object} models.Account
// @Failure 404 {object} utils.Response
// @Router /accounts/{id} [get]
func (ac *accountsController) GetAccount(c *gin.Context) {
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

	account, err := ac.AccountService.GetAccountByID(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, account)
}

// CreateAccount godoc
// @Summary Create an account
// @Description create new account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body models.Account true "Create Account"
// @Success 201 {object} models.Account
// @Failure 400 {object} utils.Response
// @Router /accounts [post]
func (ac *accountsController) CreateAccount(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		return
	}

	createdAccount, err := ac.AccountService.CreateAccount(&account)
	if err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAccount)
}

// UpdateAccount godoc
// @Summary Update an account
// @Description update account by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path uint true "Account ID"
// @Param account body models.Account true "Update Account"
// @Success 200 {object} models.Account
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /accounts/{id} [put]
func (ac *accountsController) UpdateAccount(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
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

	account.Id = uint(num)

	_, err = ac.AccountService.GetAccountByID(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	updatedAccount, err := ac.AccountService.UpdateAccount(&account)
	if err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAccount)
}

// DeleteAccount godoc
// @Summary Delete an account
// @Description delete account by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path uint true "Account ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /accounts/{id} [delete]
func (ac *accountsController) DeleteAccount(c *gin.Context) {
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

	err = ac.AccountService.DeleteAccount(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, utils.DeletedSuccessfully)
}
