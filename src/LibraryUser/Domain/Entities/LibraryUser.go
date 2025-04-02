package entities

type LibraryUser struct{
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Folio int `json:"Folio"`
	Status string `json:"status"`
	Role string `json:"role"`	
}