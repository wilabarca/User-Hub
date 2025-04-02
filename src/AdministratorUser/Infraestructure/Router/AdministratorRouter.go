package routers

import (
	controller "UserMac/src/AdministratorUser/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterAdministratorRoutes(router *gin.Engine, administratorController *controller.AdministratorUserController) {
	administratorGroup := router.Group("/administrator")
	{ 
		administratorGroup.POST("/login", administratorController.Authenticate)
		administratorGroup.GET("/", administratorController.GetAllAdministrators)
		administratorGroup.GET("/:id", administratorController.GetAdministratorByID)
		administratorGroup.POST("/", administratorController.CreateAdministrator)
		administratorGroup.PUT("/:id", administratorController.UpdateAdministrator)
		administratorGroup.DELETE("/:id", administratorController.DeleteAdministrator)
	}
}
