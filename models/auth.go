package models

type LoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterStruct struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
