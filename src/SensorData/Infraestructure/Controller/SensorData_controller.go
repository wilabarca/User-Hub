package controller

import (
	application "UserMac/src/SensorData/Application"
	entities "UserMac/src/SensorData/Domain/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SensorDataController struct {
	service *application.SensorDataService
}

func NewSensorDataController(service *application.SensorDataService) *SensorDataController {
	return &SensorDataController{service: service}
}

// Crear un nuevo registro de SensorData
func (c *SensorDataController) CreateSensorData(ctx *gin.Context) {
	var sensorData entities.SensorData
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	err := c.service.CreateSensorData(&sensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos del sensor creados"})
}

// Obtener todos los registros de SensorData
func (c *SensorDataController) GetAllSensorData(ctx *gin.Context) {
	sensorDataList, err := c.service.GetAllSensorData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sensorDataList)
}

// Obtener un registro de SensorData por ID
func (c *SensorDataController) GetSensorDataByID(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de SensorData inválido"})
		return
	}

	sensorData, err := c.service.GetSensorDataByID(int64(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sensorData)
}

// Actualizar un registro de SensorData
func (c *SensorDataController) UpdateSensorData(ctx *gin.Context) {
	id := ctx.Param("id")
	sensorDataID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sensorData entities.SensorData
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	sensorData.ID = int64(sensorDataID)

	err = c.service.UpdateSensorData(&sensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Datos del sensor actualizados"})
}

// Eliminar un registro de SensorData
func (c *SensorDataController) DeleteSensorData(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de SensorData inválido"})
		return
	}

	err = c.service.DeleteSensorData(int64(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Datos del sensor eliminados"})
}
