package database

import (
	"golang-simple-rest-api/helpers"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ClientPostgres *gorm.DB

func PGOpen() error {
	dsn := helpers.GetDSN()

	dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	ClientPostgres = dbClient

	sqlDB, err := ClientPostgres.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return nil
}
