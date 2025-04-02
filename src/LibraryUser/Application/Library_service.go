package application

import (
	"time"
	"errors"
	"log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	entities "UserMac/src/LibraryUser/Domain/Entities"
	repositories "UserMac/src/LibraryUser/Domain/Repositories"
)

type LibraryUserService struct {
	repository repositories.LibraryRepository
}

func NewLibraryService(repo repositories.LibraryRepository) *LibraryUserService {
	return &LibraryUserService{repository: repo}
}

// HashPassword genera un hash seguro de la contraseña usando bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al hashear la contraseña:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword verifica si la contraseña ingresada coincide con el hash almacenado.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateJWT genera un token JWT con información básica del usuario.
func GenerateJWT(username, role string, folio int) (string, error) {
	secretKey := []byte("TuLlaveSecretaSuperSegura")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"folio":    folio,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Expira en 72 horas
	})
	return token.SignedString(secretKey)
}

// Crear un nuevo usuario de biblioteca con estado basado en el folio.
func (s *LibraryUserService) CreateLibraryUser(user *entities.LibraryUser) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	if user.Folio > 1000 {
		user.Status = "activo"
	} else {
		user.Status = "inactivo"
	}

	return s.repository.CreateLibraryUser(user)
}

// Autenticación de usuario, devuelve el JWT.
func (s *LibraryUserService) AuthenticateUser(id int16, password string) (string, error) {
	user, err := s.repository.GetLibraryUserByID(int64(id))
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	if err := CheckPassword(user.Password, password); err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	token, err := GenerateJWT(user.Username, user.Role, user.Folio)
	if err != nil {
		return "", err
	}

	return token, nil
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
