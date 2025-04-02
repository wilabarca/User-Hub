package repositories

import (
	entities "UserMac/src/SensorData/Domain/Entities"
)

type SensorDataRepository interface {
	CreateSensorData(sensorData *entities.SensorData) error
	GetSensorDataByID(id int64) (*entities.SensorData, error)
	UpdateSensorData(sensorData *entities.SensorData) error
	DeleteSensorData(id int64) error
	GetAllSensorData() ([]entities.SensorData, error)
}
