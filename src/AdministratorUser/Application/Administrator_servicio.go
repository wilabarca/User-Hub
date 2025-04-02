package application

import (
	entities "UserMac/src/AdministratorUser/Domain/Entities"
	repositories "UserMac/src/AdministratorUser/Domain/Repositories"

	"github.com/golang-jwt/jwt"
)


type AdministratorUserService struct {
	repository          repositories.AdministratorRepository
	authenticationService *AdministratorUserService
}

func NewAdministratorUserService(repo repositories.AdministratorRepository, authService *AdministratorUserService) *AdministratorUserService {
	return &AdministratorUserService{repository: repo, authenticationService: authService}
}

func (s *AdministratorUserService) SaveAdministrator(administrator *entities.AdministratorUser) error {
	return s.repository.SaveAdministrator(administrator)
}

func (s *AdministratorUserService) GetAdministratorByID(id int64) (*entities.AdministratorUser, error) {
	return s.repository.GetAdministratorByID(id)
}

func (s *AdministratorUserService) DeleteAdministrator(id int64) error {
	return s.repository.DeleteAdministrator(id)
}

func (s *AdministratorUserService) GetAllAdministrators() ([]entities.AdministratorUser, error) {
	return s.repository.GetLAdminidtrator()
}

// Update an administrator
func (s *AdministratorUserService) UpdateAdministrator(administrator *entities.AdministratorUser) error {
	return s.repository.UpdateAdministrator(administrator)
}

// Generate JWT token for the administrator
func (s *AdministratorUserService) GenerateToken(administrator *entities.AdministratorUser) (string, error) {
	return s.authenticationService.GenerateToken(administrator)
}

// Validate the JWT token
func (s *AdministratorUserService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return s.authenticationService.ValidateToken(tokenString)
}