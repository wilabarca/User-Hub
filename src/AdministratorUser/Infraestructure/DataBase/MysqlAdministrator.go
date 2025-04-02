package database

import (
	"UserMac/src/AdministratorUser/Domain/Entities"
	"UserMac/src/AdministratorUser/Domain/Repositories"
	"database/sql"
	"fmt"
)

type MySQLAdministratorRepository struct {
	db *sql.DB
}

func NewMySQLAdministratorRepository(db *sql.DB) repositories.AdministratorRepository {
	return &MySQLAdministratorRepository{db: db}
}

// SaveAdministrator guarda un nuevo administrador en la base de datos.
func (m *MySQLAdministratorRepository) SaveAdministrator(administrator *entities.AdministratorUser) error {
	_, err := m.db.Exec("INSERT INTO administrators (username, password, email, nip) VALUES (?, ?, ?, ?)", 
		administrator.Username, administrator.Password, administrator.Email, administrator.NIP)
	return err
}

// GetAdministratorByID obtiene un administrador por su ID.
func (m *MySQLAdministratorRepository) GetAdministratorByID(id int64) (*entities.AdministratorUser, error) {
	var administrator entities.AdministratorUser
	err := m.db.QueryRow("SELECT id, username, password, email, nip FROM administrators WHERE id = ?", id).
		Scan(&administrator.ID, &administrator.Username, &administrator.Password, &administrator.Email, &administrator.NIP)
	if err != nil {
		return nil, err
	}
	return &administrator, nil
}

// UpdateAdministrator actualiza la información de un administrador en la base de datos.
func (m *MySQLAdministratorRepository) UpdateAdministrator(administrator *entities.AdministratorUser) error {
	result, err := m.db.Exec("UPDATE administrators SET username = ?, password = ?, email = ?, nip = ? WHERE id = ?",
		administrator.Username, administrator.Password, administrator.Email, administrator.NIP, administrator.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el administrador: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ningún administrador encontrado con el ID %d", administrator.ID)
	}

	return nil
}

// GetAllAdministrators obtiene todos los administradores desde la base de datos.
func (m *MySQLAdministratorRepository) GetLAdminidtrator() ([]entities.AdministratorUser, error) {
	rows, err := m.db.Query("SELECT id, username, password, email, nip FROM administrators")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var administrators []entities.AdministratorUser
	for rows.Next() {
		var administrator entities.AdministratorUser
		if err := rows.Scan(&administrator.ID, &administrator.Username, &administrator.Password, &administrator.Email, &administrator.NIP); err != nil {
			return nil, err
		}
		administrators = append(administrators, administrator)
	}
	return administrators, nil
}

// DeleteAdministrator elimina un administrador de la base de datos por su ID.
func (m *MySQLAdministratorRepository) DeleteAdministrator(id int64) error {
	_, err := m.db.Exec("DELETE FROM administrators WHERE id = ?", id)
	return err
}

// SaveHashedToken guarda el token hasheado del administrador en la base de datos.
func (m *MySQLAdministratorRepository) SaveHashedToken(userID int64, hashedToken string) error {
	_, err := m.db.Exec("UPDATE administrators SET hashed_token = ? WHERE id = ?", hashedToken, userID)
	if err != nil {
		return fmt.Errorf("error al guardar el token hasheado: %v", err)
	}
	return nil
}
