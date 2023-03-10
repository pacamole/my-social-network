package repositories

import (
	"errors"
	m "my-social-network/models"
)

func (c *Client) InsertUser(user m.User) error {
	c.db.Create(user)
	if c.db.Error != nil {
		return c.db.Error
	}

	return nil
}

func (c *Client) SelectAllUsers() ([]m.User, error) {
	var users []m.User
	c.db.Table("users").Take(&users)
	if c.db.Error != nil {
		return nil, c.db.Error
	}

	return users, nil
}

func (c *Client) SelectUserByID(ID uint16) (m.User, error) {
	var user m.User
	c.db.Table("users").First(&user, ID)
	if c.db.Error != nil {
		return m.User{}, c.db.Error
	}

	if user.Email == "" {
		return m.User{}, errors.New("user not found")
	}

	return user, nil
}

func (c *Client) SelectUserByEmail(email string) (m.User, error) {
	var user m.User
	c.db.Table("users").Where("email = ?", email).Find(&user)
	if c.db.Error != nil {
		return m.User{}, c.db.Error
	}
	if user.Email == "" {
		return m.User{}, errors.New("user not found")
	}

	return user, nil
}

func (c *Client) UpdateUserByID(ID uint16, name string) error {
	c.db.Table("users").Where(m.User{}, ID).Update("name", name)
	if c.db.Error != nil {
		return c.db.Error
	}

	return nil
}

func (c *Client) DeleteUserByID(ID uint16) error {

	c.db.Delete(m.User{}, ID)
	if c.db.Error != nil {
		return c.db.Error
	}

	return nil
}

func (c *Client) LoginUser(email, password string) error {
	// var result map[string]interface{}
	c.db.Table("users").Where(map[string]interface{}{
		"email":    email,
		"password": password,
	})
	if c.db.Error != nil {
		return c.db.Error
	}

	return nil
}
