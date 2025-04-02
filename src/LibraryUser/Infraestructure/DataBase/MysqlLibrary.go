package database

import (
	"UserMac/src/LibraryUser/Domain/Entities"
	"UserMac/src/LibraryUser/Domain/Repositories"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type MySQLLibraryUserRepository struct {
	db *sql.DB
}

func NewMySQLLibraryUserRepository(db *sql.DB) repositories.LibraryRepository {
	return &MySQLLibraryUserRepository{db: db}
}

// HashPassword hashes a plain text password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CreateLibraryUser saves a new library user in the database with hashed password and status.
func (m *MySQLLibraryUserRepository) CreateLibraryUser(user *entities.LibraryUser) error {
	// Hash the user password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Determine the user status based on the folio
	var status string
	if user.Folio > 1000 {
		status = "activo"
	} else {
		status = "inactivo"
	}

	// Insert the user with hashed password and status
	_, err = m.db.Exec("INSERT INTO library_users (username, password, email, folio, status) VALUES (?, ?, ?, ?, ?)",
		user.Username, hashedPassword, user.Email, user.Folio, status)
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}

	return nil
}

// ComparePassword compares the provided password with the hashed password stored in the database.
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GetLibraryUserByID retrieves a library user by their ID.
func (m *MySQLLibraryUserRepository) GetLibraryUserByID(id int64) (*entities.LibraryUser, error) {
	var user entities.LibraryUser
	err := m.db.QueryRow("SELECT id, username, password, email, folio, status FROM library_users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Folio, &user.Status)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}
	return &user, nil
}

// GetLibraryUser retrieves all library users from the database.
func (m *MySQLLibraryUserRepository) GetLibraryUser() ([]entities.LibraryUser, error) {
	rows, err := m.db.Query("SELECT id, username, email, folio, status FROM library_users")
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %v", err)
	}
	defer rows.Close()

	var users []entities.LibraryUser
	for rows.Next() {
		var user entities.LibraryUser
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Folio, &user.Status); err != nil {
			return nil, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateLibraryUser updates a library user in the database.
func (m *MySQLLibraryUserRepository) UpdateLibraryUser(user *entities.LibraryUser) error {
	_, err := m.db.Exec("UPDATE library_users SET username = ?, email = ?, folio = ?, status = ? WHERE id = ?",
		user.Username, user.Email, user.Folio, user.Status, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

// DeleteLibraryUser deletes a library user from the database.
func (m *MySQLLibraryUserRepository) DeleteLibraryUser(id int64) error {
	_, err := m.db.Exec("DELETE FROM library_users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
