package database

import (
	"UserMac/src/LibraryUser/Domain/Entities"
	"UserMac/src/LibraryUser/Domain/Repositories"
	"database/sql"
	"fmt"
)

type MySQLLibraryUserRepository struct {
	db *sql.DB
}

func NewMySQLLibraryUserRepository(db *sql.DB) repositories.LibraryRepository {
	return &MySQLLibraryUserRepository{db: db}
}

// CreateLibraryUser guarda un nuevo usuario de biblioteca en la base de datos y asigna su estado.
func (m *MySQLLibraryUserRepository) CreateLibraryUser(user *entities.LibraryUser) error {
	// Lógica para determinar el estado del usuario basado en el folio
	var status string
	if user.Folio > 1000 { // Ejemplo de condición: si el folio es mayor que 1000, el usuario es "activo"
		status = "activo"
	} else {
		status = "inactivo"
	}

	// Insertar el nuevo usuario con su estado
	_, err := m.db.Exec("INSERT INTO library_users (username, password, email, folio, status) VALUES (?, ?, ?, ?, ?)", 
		user.Username, user.Password, user.Email, user.Folio, status)
	return err
}

// GetLibraryUserByID obtiene un usuario de biblioteca por su ID.
func (m *MySQLLibraryUserRepository) GetLibraryUserByID(id int64) (*entities.LibraryUser, error) {
	var user entities.LibraryUser
	err := m.db.QueryRow("SELECT id, username, password, email, folio FROM library_users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Folio)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateLibraryUser actualiza la información de un usuario de biblioteca en la base de datos.
func (m *MySQLLibraryUserRepository) UpdateLibraryUser(user *entities.LibraryUser) error {
	result, err := m.db.Exec("UPDATE library_users SET username = ?, password = ?, email = ?, folio = ? WHERE id = ?",
		user.Username, user.Password, user.Email, user.Folio, user.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el usuario de biblioteca: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ningún usuario encontrado con el ID %d", user.ID)
	}

	return nil
}

// GetAllLibraryUser obtiene todos los usuarios de biblioteca desde la base de datos.
func (m *MySQLLibraryUserRepository) GetLibraryUser() ([]entities.LibraryUser, error) {
	rows, err := m.db.Query("SELECT id, username, password, email, folio FROM library_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.LibraryUser
	for rows.Next() {
		var user entities.LibraryUser
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Folio); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// DeleteLibraryUser elimina un usuario de biblioteca de la base de datos por su ID.
func (m *MySQLLibraryUserRepository) DeleteLibraryUser(id int64) error {
	_, err := m.db.Exec("DELETE FROM library_users WHERE id = ?", id)
	return err
}
