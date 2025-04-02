package routers

import (
	controller "UserMac/src/LibraryUser/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterLibraryUserRoutes(router *gin.Engine, LibraryUserController *controller.LibraryUserController) {
    LibraryUserGroup := router.Group("/LibraryUser")
    {
        LibraryUserGroup.POST("/login", LibraryUserController.AuthenticateLibraryUser)
        LibraryUserGroup.GET("/", LibraryUserController.GetAllLibraryUsers)
        LibraryUserGroup.GET("/:id", LibraryUserController.GetLibraryUserByID)
        LibraryUserGroup.POST("/", LibraryUserController.CreateLibraryUser)
        LibraryUserGroup.PUT("/:id", LibraryUserController.UpdateLibraryUser)
        LibraryUserGroup.DELETE("/:id", LibraryUserController.DeleteLibraryUser)
    }
}
