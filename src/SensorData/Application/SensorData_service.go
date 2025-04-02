package application

import (
	entities "UserMac/src/SensorData/Domain/Entities"
	repositories "UserMac/src/SensorData/Domain/Repositories"
)

type SensorDataService struct {
	repository repositories.SensorDataRepository
}

func NewSensorDataService(repo repositories.SensorDataRepository) *SensorDataService {
	return &SensorDataService{repository: repo}
}

func (s *SensorDataService) CreateSensorData(sensorData *entities.SensorData) error {
	return s.repository.CreateSensorData(sensorData)
}

func (s *SensorDataService) GetSensorDataByID(id int64) (*entities.SensorData, error) {
	return s.repository.GetSensorDataByID(id)
}

func (s *SensorDataService) DeleteSensorData(id int64) error {
	return s.repository.DeleteSensorData(id)
}

func (s *SensorDataService) GetAllSensorData() ([]entities.SensorData, error) {
	return s.repository.GetAllSensorData()
}

// Función para actualizar un registro de SensorData
func (s *SensorDataService) UpdateSensorData(sensorData *entities.SensorData) error {
	return s.repository.UpdateSensorData(sensorData)
}
