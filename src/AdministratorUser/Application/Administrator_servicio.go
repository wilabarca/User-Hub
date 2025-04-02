package application

import (
	services "UserMac/src/AdministratorUser/Application/Services"
	entities "UserMac/src/AdministratorUser/Domain/Entities"
	repositories "UserMac/src/AdministratorUser/Domain/Repositories"

	"github.com/golang-jwt/jwt"
)

type AdministratorUserService struct {
	repository          repositories.AdministratorRepository
	authenticationService *services.AuthenticationService
}

func NewAdministratorUserService(repo repositories.AdministratorRepository, authService *services.AuthenticationService) *AdministratorUserService {
	return &AdministratorUserService{repository: repo, authenticationService: authService}
}

func (s *AdministratorUserService) SaveAdministrator(administrator *entities.AdministratorUser) error {
	return s.repository.SaveAdministrator(administrator)
}

func (s *AdministratorUserService) GetAllAdministrators() ([]entities.AdministratorUser, error) {
	return s.repository.GetLAdminidtrator()
}

func (s *AdministratorUserService) GetAdministratorByID(id int64) (*entities.AdministratorUser, error) {
	return s.repository.GetAdministratorByID(id)
}

func (s *AdministratorUserService) UpdateAdministrator(administrator *entities.AdministratorUser) error {
	return s.repository.UpdateAdministrator(administrator)
}

func (s *AdministratorUserService) DeleteAdministrator(id int64) error {
	return s.repository.DeleteAdministrator(id)
}

func (s *AdministratorUserService) GenerateToken(administrator entities.AdministratorUser) (string, string, error) {
	return s.authenticationService.GenerateToken(administrator.ID, administrator.Username)
}

func (s *AdministratorUserService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return s.authenticationService.ValidateToken(tokenString)
}

func (s *AdministratorUserService) HashToken(token string) (string, error) {
	return s.authenticationService.HashToken(token)
}

func (s *AdministratorUserService) SaveHashedToken(userID int64, hashedToken string) error {
	return s.repository.SaveHashedToken(userID, hashedToken)
}
