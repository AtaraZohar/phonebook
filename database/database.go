package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "host=db user=atara password=mypassword dbname=phonebook port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}
