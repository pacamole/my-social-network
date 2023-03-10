package services

import (
	"errors"
	"my-social-network/helper"
	m "my-social-network/models"
	repos "my-social-network/repositories"
	"time"
)

type (
	IUsers interface {
		CreateUser(email, name, password string) error

		GetAllUsers() ([]m.User, error)
		GetUserByID(uint16) (m.User, error)
		GetUserByEmail(string) (m.User, error)

		UpdateUserByID(uint16, string) error

		DeleteUserByID(uint16) error
	}

	UserConn struct {
		db *repos.Client
	}
)

func NewUserConn(db *repos.Client) (IUsers, error) {
	return &UserConn{db}, nil
}

func (u *UserConn) CreateUser(email, name, password string) error {
	us, err := u.GetUserByEmail(email)
	if us.ID == 0 {
		return errors.New("user already exists")
	}
	if err != nil {
		return err
	}

	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	var user m.User = m.User{
		Email:      email,
		Name:       name,
		Password:   hashedPassword,
		Created_at: time.Now(),
	}

	if err := u.db.InsertUser(user); err != nil {
		return err
	}

	return nil
}

func (u *UserConn) GetAllUsers() ([]m.User, error) {
	users, err := u.db.SelectAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserConn) GetUserByID(ID uint16) (m.User, error) {
	user, err := u.db.SelectUserByID(ID)
	if err != nil {
		return m.User{}, err
	}

	return user, nil
}

func (u *UserConn) GetUserByEmail(email string) (m.User, error) {
	user, err := u.db.SelectUserByEmail(email)
	if err != nil {
		return m.User{}, err
	}

	return user, nil
}

func (u *UserConn) UpdateUserByID(ID uint16, email string) error {
	if err := u.db.UpdateUserByID(ID, email); err != nil {
		return err
	}

	return nil
}

func (u *UserConn) DeleteUserByID(ID uint16) error {
	if err := u.db.DeleteUserByID(ID); err != nil {
		return err
	}

	return nil
}
