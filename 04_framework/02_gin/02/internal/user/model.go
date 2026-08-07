package user

type User struct{
	Id int `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"Passwordhash"`
	Role string `json:"role"`

}

type AuthResult struct{
	Token string `json:"token"`
	User User	`json:"user"`
}

type UserCreationRequest struct{
	Email string `json:"email"`
	Password string 	`json:"password"`
}

type UserCreationResponse struct{
	Token string `json:"token"`
	Email string `json:"email"`
}
