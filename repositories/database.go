package repositories

import (
	"fmt"
	"my-social-network/config"
	m "my-social-network/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() (*gorm.DB, error) {
	var dns string = fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		config.DbUser, config.DbPassword, config.DbHost, config.DbName,
	)

	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Client struct {
	db *gorm.DB
}

func NewGormClient(db *gorm.DB) *Client {
	return &Client{db}
}

func MigrateGormDB(db *gorm.DB) error {
	return db.AutoMigrate(&m.User{})
}
