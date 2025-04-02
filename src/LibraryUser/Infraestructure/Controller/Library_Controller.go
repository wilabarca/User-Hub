package controller

import (
	application "UserMac/src/LibraryUser/Application"
	entities "UserMac/src/LibraryUser/Domain/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LibraryUserController struct {
	service *application.LibraryUserService
}

func NewLibraryUserController(service *application.LibraryUserService) *LibraryUserController {
	return &LibraryUserController{service: service}
}

// Crear un usuario de biblioteca
func (c *LibraryUserController) CreateLibraryUser(ctx *gin.Context) {
	var user entities.LibraryUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	// Establecer el status basado en el folio
	if user.Folio > 1000 {  // Suponiendo que el folio mayor a 1000 activa al usuario
		user.Status = "activo"
	} else {
		user.Status = "inactivo"  // Si el folio es menor o igual a 1000, el usuario es inactivo
	}

	err := c.service.CreateLibraryUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario de Biblioteca creado"})
}

// Obtener todos los usuarios de biblioteca
func (c *LibraryUserController) GetAllLibraryUsers(ctx *gin.Context) {
	users, err := c.service.GetLibraryUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// Obtener un usuario de biblioteca por ID
func (c *LibraryUserController) GetLibraryUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	user, err := c.service.GetLibraryUserByID(int16(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Actualizar un usuario de biblioteca
func (c *LibraryUserController) UpdateLibraryUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user entities.LibraryUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	user.ID = int64(userID)

	// Establecer el status basado en el folio al actualizar
	if user.Folio > 1000 {  // Folio mayor a 1000: status activo
		user.Status = "activo"
	} else {
		user.Status = "inactivo"  // Folio menor o igual a 1000: status inactivo
	}

	err = c.service.UpdateLibraryUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario de Biblioteca actualizado"})
}

// Eliminar un usuario de biblioteca
func (c *LibraryUserController) DeleteLibraryUser(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	err = c.service.DeleteLibraryUser(int16(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario de Biblioteca eliminado"})
}
// Autenticar un usuario de biblioteca y devolver un token JWT.
func (c *LibraryUserController) AuthenticateLibraryUser(ctx *gin.Context) {
	var loginData struct {
		ID       int16  `json:"id"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	token, err := c.service.AuthenticateUser(loginData.ID, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
