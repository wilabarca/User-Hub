package entities

type AdministratorUser struct{
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	NIP string `json:"nip"`
}