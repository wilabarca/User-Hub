package controller

import (
	application "UserMac/src/AdministratorUser/Application"
	entities "UserMac/src/AdministratorUser/Domain/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdministratorUserController struct {
	service *application.AdministratorUserService
}

func NewAdministratorUserController(service *application.AdministratorUserService) *AdministratorUserController {
	return &AdministratorUserController{service: service}
}

// Crear un administrador
func (c *AdministratorUserController) CreateAdministrator(ctx *gin.Context) {
	var administrator entities.AdministratorUser
	if err := ctx.ShouldBindJSON(&administrator); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	// Token enviado en el header
	token := ctx.GetHeader("Authorization")

	// Validación del token antes de crear el administrador
	_, err := c.service.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido o no autorizado"})
		return
	}

	// Guardar el administrador
	err = c.service.SaveAdministrator(&administrator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generar el hash del token
	hashedToken, err := c.service.HashToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el hash del token"})
		return
	}

	// Guardar el hash del token en la base de datos
	err = c.service.SaveHashedToken(administrator.ID, hashedToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el hash del token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Administrador creado"})
}

// Obtener todos los administradores
func (c *AdministratorUserController) GetAllAdministrators(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	_, err := c.service.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido o no autorizado"})
		return
	}

	administrators, err := c.service.GetAllAdministrators()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, administrators)
}

// Obtener un administrador por ID
func (c *AdministratorUserController) GetAdministratorByID(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Extraer token del encabezado de la solicitud
	token := ctx.GetHeader("Authorization")
	administrator, err := c.service.GetAdministratorByID(num)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el token es válido
	_, err = c.service.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido o no autorizado"})
		return
	}

	ctx.JSON(http.StatusOK, administrator)
}

// Actualizar un administrador
func (c *AdministratorUserController) UpdateAdministrator(ctx *gin.Context) {
	id := ctx.Param("id")
	administratorID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var administrator entities.AdministratorUser
	if err := ctx.ShouldBindJSON(&administrator); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	administrator.ID = administratorID

	err = c.service.UpdateAdministrator(&administrator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador actualizado"})
}

// Eliminar un administrador
func (c *AdministratorUserController) DeleteAdministrator(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.service.DeleteAdministrator(num)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Administrador eliminado"})
}

// Login endpoint para autenticar y devolver el token JWT
func (c *AdministratorUserController) Authenticate(ctx *gin.Context) {
	var administrator entities.AdministratorUser
	if err := ctx.ShouldBindJSON(&administrator); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	// Validar las credenciales del usuario aquí (comprobar la contraseña y el nombre de usuario en la DB)
	// Si son válidas, generar el token
	token, hashedToken, err := c.service.GenerateToken(administrator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Aquí puedes guardar el hash en la base de datos si es necesario
	// Ejemplo: c.service.SaveHashedToken(administrator.ID, hashedToken)

	ctx.JSON(http.StatusOK, gin.H{
		"token":        token,
		"hashed_token": hashedToken,
	})
}
