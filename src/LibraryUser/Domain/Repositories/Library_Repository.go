package repositories

import entities "UserMac/src/LibraryUser/Domain/Entities"

type LibraryRepository interface {
	CreateLibraryUser(user *entities.LibraryUser) error
	GetLibraryUser() ([]entities.LibraryUser, error)
	UpdateLibraryUser(user *entities.LibraryUser) error
	DeleteLibraryUser(id int64) error
	GetLibraryUserByID(id int64) (*entities.LibraryUser, error) 
}