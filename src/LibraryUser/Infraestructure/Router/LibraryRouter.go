package routers

import (
	"UserMac/src/LibraryUser/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

// RegisterLibraryUserRoutes registra las rutas de los usuarios de biblioteca en el router.
func RegisterLibraryUserRoutes(router *gin.Engine, libraryUserController *controller.LibraryUserController) {
	// Agrupamos las rutas relacionadas con LibraryUser
	LibraryUserGroup := router.Group("/LibraryUser")
	{
		// Rutas públicas (no requieren autenticación)
		LibraryUserGroup.POST("/login", libraryUserController.AuthenticateLibraryUser) // Login
		LibraryUserGroup.GET("/", libraryUserController.GetAllLibraryUsers)           // Obtener todos los usuarios
		
		// Ruta con ID (también no requiere autenticación)
		LibraryUserGroup.GET("/:id", libraryUserController.GetLibraryUserByID) // Obtener usuario por ID

		// Rutas que no requieren autenticación para crear, actualizar y eliminar
		LibraryUserGroup.POST("/", libraryUserController.CreateLibraryUser)     // Crear usuario
		LibraryUserGroup.PUT("/:id", libraryUserController.UpdateLibraryUser)   // Actualizar usuario
		LibraryUserGroup.DELETE("/:id", libraryUserController.DeleteLibraryUser) // Eliminar usuario
	}
}
