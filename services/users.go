package services

import (
	m "my-social-network/models"
)

type IUsers interface {
	CreateUser(email, name, password string) error

	GetAllUsers() ([]m.User, error)
	GetUserByID(string) (m.User, error)
	GetUserByEmail(string) (m.User, error)

	UpdateUserByID(string) error

	DeleteUserByID(int) error
}
