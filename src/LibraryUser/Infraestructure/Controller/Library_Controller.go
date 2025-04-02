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

	if user.Folio > 1000 {
		user.Status = "activo"
	} else {
		user.Status = "inactivo"
	}

	err := c.service.CreateLibraryUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario de biblioteca"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Usuario de Biblioteca creado exitosamente"})
}

// Obtener todos los usuarios de biblioteca
func (c *LibraryUserController) GetAllLibraryUsers(ctx *gin.Context) {
	users, err := c.service.GetLibraryUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
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
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Autenticar un usuario de biblioteca y devolver un token JWT.
func (c *LibraryUserController) AuthenticateLibraryUser(ctx *gin.Context) {
	var loginData struct {
		ID       int16  `json:"id"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de autenticación inválidos"})
		return
	}

	token, err := c.service.AuthenticateUser(loginData.ID, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Autenticación exitosa", "token": token})
}
// Eliminar un usuario de biblioteca por ID
func (c *LibraryUserController) DeleteLibraryUser(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	err = c.service.DeleteLibraryUser(int16(num))
	if err != nil {
		// Si no se encuentra el usuario, respondemos con 404
		if err.Error() == "usuario no encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		}
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Usuario eliminado exitosamente"})
}
// Actualizar un usuario de biblioteca por ID
func (c *LibraryUserController) UpdateLibraryUser(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var updatedUser entities.LibraryUser
	// Intentar enlazar los datos JSON del cuerpo de la solicitud al struct LibraryUser
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos para actualizar el usuario"})
		return
	}

	// Asignar el ID del usuario que se va a actualizar
	updatedUser.ID = int64(num)

	// Llamar al servicio para actualizar el usuario
	err = c.service.UpdateLibraryUser(&updatedUser)
	if err != nil {
		// Si no se encuentra el usuario, respondemos con 404
		if err.Error() == "usuario no encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		}
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, gin.H{"mensaje": "Usuario actualizado exitosamente"})
}
