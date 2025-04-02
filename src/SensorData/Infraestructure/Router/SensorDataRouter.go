package routers

import (
	"github.com/gin-gonic/gin"
	controller "UserMac/src/SensorData/Infraestructure/Controller"
)

func RegisterSensorDataRoutes(router *gin.Engine, SensorDataController *controller.SensorDataController) {
	SensorDataGroup := router.Group("/SensorData")
	{
		SensorDataGroup.GET("/", SensorDataController.GetAllSensorData)           // Obtener todos los datos del sensor
		SensorDataGroup.GET("/:id", SensorDataController.GetSensorDataByID)      // Obtener datos del sensor por ID
		SensorDataGroup.POST("/", SensorDataController.CreateSensorData)         // Crear un nuevo dato del sensor
		SensorDataGroup.PUT("/:id", SensorDataController.UpdateSensorData)      // Actualizar datos del sensor por ID
		SensorDataGroup.DELETE("/:id", SensorDataController.DeleteSensorData)   // Eliminar datos del sensor por ID
	}
}
