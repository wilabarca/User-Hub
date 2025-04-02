package application

import (
	services "UserMac/src/LibraryUser/Application/Services"
	entities "UserMac/src/LibraryUser/Domain/Entities"
	repositories "UserMac/src/LibraryUser/Domain/Repositories"

	"errors"
)

// LibraryUserService maneja las operaciones del usuario de la biblioteca.
type LibraryUserService struct {
	repository repositories.LibraryRepository
}

// NewLibraryService crea un nuevo servicio de usuario de biblioteca.
func NewLibraryService(repo repositories.LibraryRepository) *LibraryUserService {
	return &LibraryUserService{repository: repo}
}

// Crear un nuevo usuario de biblioteca con estado basado en el folio.
func (s *LibraryUserService) CreateLibraryUser(user *entities.LibraryUser) error {
	// Asignar el estado según el folio
	if user.Folio > 1000 {
		user.Status = "activo" // Folio mayor a 1000, el usuario está activo
	} else {
		user.Status = "inactivo" // Folio menor o igual a 1000, el usuario está inactivo
	}

	return s.repository.CreateLibraryUser(user)
}

// Obtener un usuario de biblioteca por ID.
func (s *LibraryUserService) GetLibraryUserByID(id int16) (*entities.LibraryUser, error) {
	return s.repository.GetLibraryUserByID(int64(id))
}

// Eliminar un usuario de biblioteca por ID.
func (s *LibraryUserService) DeleteLibraryUser(id int16) error {
	return s.repository.DeleteLibraryUser(int64(id))
}

// Obtener todos los usuarios de biblioteca.
func (s *LibraryUserService) GetLibraryUser() ([]entities.LibraryUser, error) {
	return s.repository.GetLibraryUser()
}

// Actualizar un usuario de biblioteca.
func (s *LibraryUserService) UpdateLibraryUser(user *entities.LibraryUser) error {
	return s.repository.UpdateLibraryUser(user)
}

// Autenticación de usuario, devuelve el JWT.
func (s *LibraryUserService) AuthenticateUser(id int16, password string) (string, error) {
	// Verificar usuario en la base de datos por ID.
	user, err := s.repository.GetLibraryUserByID(int64(id))
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	// Validar la contraseña (esto es un ejemplo, deberías usar un hash de contraseña).
	if user.Password != password {
		return "", errors.New("contraseña incorrecta")
	}

	// Generar el JWT.
	token, err := services.GenerateJWT(user.Username, user.Role, user.Folio)
	if err != nil {
		return "", err
	}

	return token, nil
}
