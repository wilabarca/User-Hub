package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando archivo .env: %v", err)
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")

	// Abrir la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar la base de datos con DSN '%v': %v", dsn, err)
	}

	// Crear tablas si no existen
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("error creando tablas: %v", err)
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	// Eliminar las tablas existentes si es necesario
	// (esto solo debe hacerse si deseas eliminar completamente las tablas anteriores)
	_, err := db.Exec("DROP TABLE IF EXISTS library_users, administrator_users")
	if err != nil {
		return fmt.Errorf("error eliminando tablas existentes: %v", err)
	}

	// Crear tabla para LibraryUser con el campo status agregado
	libraryUserTable := `CREATE TABLE IF NOT EXISTS library_users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		folio INT NOT NULL,
		status ENUM('activo', 'inactivo') NOT NULL DEFAULT 'inactivo'
	)`

	// Crear tabla para AdministratorUser, con campo para el token hasheado
	administratorUserTable := `CREATE TABLE IF NOT EXISTS administrator_users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		nip VARCHAR(255) NOT NULL,
		hashed_token VARCHAR(255) DEFAULT NULL  -- Agregar campo para el token hasheado
	)`

	// Ejecutar creación de tabla library_users
	if _, err := db.Exec(libraryUserTable); err != nil {
		return fmt.Errorf("error creando tabla library_users: %v", err)
	}

	// Ejecutar creación de tabla administrator_users
	if _, err := db.Exec(administratorUserTable); err != nil {
		return fmt.Errorf("error creando tabla administrator_users: %v", err)
	}

	return nil
}
