package database

import (
	"UserMac/src/SensorData/Domain/Entities"
	"UserMac/src/SensorData/Domain/Repositories"
	"database/sql"
	"fmt"
)

type MySQLSensorDataRepository struct {
	db *sql.DB
}

func NewMySQLSensorDataRepository(db *sql.DB) repositories.SensorDataRepository {
	return &MySQLSensorDataRepository{db: db}
}

// CreateSensorData guarda un nuevo registro de datos de sensor en la base de datos.
func (m *MySQLSensorDataRepository) CreateSensorData(sensorData *entities.SensorData) error {
	_, err := m.db.Exec("INSERT INTO sensor_data (direccion_mac, ubicacion) VALUES (?, ?)", 
		sensorData.DireccionMac, sensorData.Ubicacion)
	return err
}

// GetSensorDataByID obtiene un registro de datos de sensor por su ID.
func (m *MySQLSensorDataRepository) GetSensorDataByID(id int64) (*entities.SensorData, error) {
	var sensorData entities.SensorData
	err := m.db.QueryRow("SELECT id, direccion_mac, ubicacion FROM sensor_data WHERE id = ?", id).
		Scan(&sensorData.ID, &sensorData.DireccionMac, &sensorData.Ubicacion)
	if err != nil {
		return nil, err
	}
	return &sensorData, nil
}

// UpdateSensorData actualiza la información de un registro de datos de sensor en la base de datos.
func (m *MySQLSensorDataRepository) UpdateSensorData(sensorData *entities.SensorData) error {
	result, err := m.db.Exec("UPDATE sensor_data SET direccion_mac = ?, ubicacion = ? WHERE id = ?",
		sensorData.DireccionMac, sensorData.Ubicacion, sensorData.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar los datos del sensor: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ningún dato encontrado con el ID %d", sensorData.ID)
	}

	return nil
}

// GetAllSensorData obtiene todos los registros de datos de sensores desde la base de datos.
func (m *MySQLSensorDataRepository) GetAllSensorData() ([]entities.SensorData, error) {
	rows, err := m.db.Query("SELECT id, direccion_mac, ubicacion FROM sensor_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensorDataList []entities.SensorData
	for rows.Next() {
		var sensorData entities.SensorData
		if err := rows.Scan(&sensorData.ID, &sensorData.DireccionMac, &sensorData.Ubicacion); err != nil {
			return nil, err
		}
		sensorDataList = append(sensorDataList, sensorData)
	}
	return sensorDataList, nil
}

// DeleteSensorData elimina un registro de datos de sensor de la base de datos por su ID.
func (m *MySQLSensorDataRepository) DeleteSensorData(id int64) error {
	_, err := m.db.Exec("DELETE FROM sensor_data WHERE id = ?", id)
	return err
}
