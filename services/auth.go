package services

import (
	"my-social-network/helper"
)

type IAuth interface {
	JwtAuth(email, password string) error
}

func (u *UserConn) JwtAuth(email, password string) error {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	if err := u.db.LoginUser(email, hashedPassword); err != nil {
		return err
	}
	return nil
}
