package main

import (
	"log"

	"UserMac/core"
	adminApp "UserMac/src/AdministratorUser/Application"
	services "UserMac/src/AdministratorUser/Application/Services"
	adminController "UserMac/src/AdministratorUser/Infraestructure/Controller"
	adminRouters "UserMac/src/AdministratorUser/Infraestructure/Router"
	adminRepo "UserMac/src/AdministratorUser/Infraestructure/database"

	libraryApp "UserMac/src/LibraryUser/Application"
	libraryController "UserMac/src/LibraryUser/Infraestructure/Controller"
	libraryRouters "UserMac/src/LibraryUser/Infraestructure/Router"
	libraryRepo "UserMac/src/LibraryUser/Infraestructure/database"

	sensorApp "UserMac/src/SensorData/Application"
	sensorController "UserMac/src/SensorData/Infraestructure/Controller"
	sensorRouters "UserMac/src/SensorData/Infraestructure/Router"
	sensorRepo "UserMac/src/SensorData/Infraestructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect")
		return
	}
	defer db.Close()

	// Inicializar repositorios
	adminRepository := adminRepo.NewMySQLAdministratorRepository(db)
	libraryRepository := libraryRepo.NewMySQLLibraryUserRepository(db)
	sensorRepository := sensorRepo.NewMySQLSensorDataRepository(db)

	// Configurar autenticación JWT
	authService, err := services.NewAuthenticationService()
	if err != nil {
		log.Fatal("Error al configurar el servicio de autenticación:", err)
	}

	// Inicializar servicios
	adminService := adminApp.NewAdministratorUserService(adminRepository, authService)
	libraryService := libraryApp.NewLibraryService(libraryRepository)
	sensorService := sensorApp.NewSensorDataService(sensorRepository)

	// Crear controladores
	adminCtrl := adminController.NewAdministratorUserController(adminService)
	libraryCtrl := libraryController.NewLibraryUserController(libraryService)
	sensorCtrl := sensorController.NewSensorDataController(sensorService)

	// Configurar el router Gin
	router := gin.Default()

	// Registrar rutas
	adminRouters.RegisterAdministratorRoutes(router, adminCtrl)
	libraryRouters.RegisterLibraryUserRoutes(router, libraryCtrl)
	sensorRouters.RegisterSensorDataRoutes(router, sensorCtrl)

	// Iniciar servidor
	port := ":5000"
	log.Printf("Servidor iniciado en el puerto %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
