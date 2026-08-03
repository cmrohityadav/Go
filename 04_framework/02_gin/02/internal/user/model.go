package user

type User struct{
	Id int `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"Passwordhash"`
	Role string `json:"role"`

}