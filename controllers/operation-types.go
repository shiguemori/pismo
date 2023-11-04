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

type OperationTypesController interface {
	Get(c *gin.Context)
	ListAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type operationTypesController struct {
	OperationTypeService services.OperationTypesService
}

func NewOperationTypeController(service services.OperationTypesService) OperationTypesController {
	return &operationTypesController{
		OperationTypeService: service,
	}
}

// ListAll OperationType godoc
// @Summary List all operationTypes
// @Description get operationTypes
// @Tags operationTypes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.OperationType
// @Failure 404 {object} utils.Response
// @Router /operation-types [get]
func (ac *operationTypesController) ListAll(c *gin.Context) {
	operationType, err := ac.OperationTypeService.ListAllOperationType()
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	if len(operationType) == 0 {
		slog.Error(utils.NotFound)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, operationType)
}

// Get OperationType godoc
// @Summary Get an operationType
// @Description get operationType by ID
// @Tags operationTypes
// @Accept  json
// @Produce  json
// @Param id path uint true "OperationType ID"
// @Success 200 {object} models.OperationType
// @Failure 404 {object} utils.Response
// @Router /operation-types/{id} [get]
func (ac *operationTypesController) Get(c *gin.Context) {
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

	operationType, err := ac.OperationTypeService.GetOperationTypeByID(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, operationType)
}

// Create OperationType godoc
// @Summary Create an operationType
// @Description create new operationType
// @Tags operationTypes
// @Accept  json
// @Produce  json
// @Param operationType body models.OperationType true "Create OperationType"
// @Success 201 {object} models.OperationType
// @Failure 400 {object} utils.Response
// @Router /operation-types [post]
func (ac *operationTypesController) Create(c *gin.Context) {
	var operationType models.OperationType

	if err := c.ShouldBindJSON(&operationType); err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		return
	}

	createdOperationType, err := ac.OperationTypeService.CreateOperationType(&operationType)
	if err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdOperationType)
}

// Update OperationType godoc
// @Summary Update an operationType
// @Description update operationType by ID
// @Tags operationTypes
// @Accept  json
// @Produce  json
// @Param id path uint true "OperationType ID"
// @Param operationType body models.OperationType true "Update OperationType"
// @Success 200 {object} models.OperationType
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /operation-types/{id} [put]
func (ac *operationTypesController) Update(c *gin.Context) {
	var operationType models.OperationType

	if err := c.ShouldBindJSON(&operationType); err != nil {
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

	operationType.Id = uint(num)

	_, err = ac.OperationTypeService.GetOperationTypeByID(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	updatedOperationType, err := ac.OperationTypeService.UpdateOperationType(&operationType)
	if err != nil {
		slog.Error(err.Error(), err)
		c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedOperationType)
}

// Delete OperationType godoc
// @Summary Delete an operationType
// @Description delete operationType by ID
// @Tags operationTypes
// @Accept  json
// @Produce  json
// @Param id path uint true "OperationType ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /operation-types/{id} [delete]
func (ac *operationTypesController) Delete(c *gin.Context) {
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

	err = ac.OperationTypeService.DeleteOperationType(uint(num))
	if err != nil {
		slog.Error(utils.NotFound, err)
		c.JSON(http.StatusNotFound, utils.Response{Message: utils.NotFound})
		return
	}

	c.JSON(http.StatusOK, utils.DeletedSuccessfully)
}
