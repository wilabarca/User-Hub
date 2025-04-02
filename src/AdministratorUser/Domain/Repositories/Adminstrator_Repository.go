package repositories

import entities "UserMac/src/AdministratorUser/Domain/Entities"

type AdministratorRepository interface{
	SaveAdministrator(administrator *entities.AdministratorUser) error
	GetLAdminidtrator() ([]entities.AdministratorUser, error)
	UpdateAdministrator(administrator *entities.AdministratorUser) error
	DeleteAdministrator(id int64) error
	GetAdministratorByID(id int64) (*entities.AdministratorUser, error) 
}
