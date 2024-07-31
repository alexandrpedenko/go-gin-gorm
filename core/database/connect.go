package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	utils "gin-gorm-crud/core/utils"
)

// TODO: Move to .env
const (
	host     = "localhost"
	port     = 5432
	dbName   = "postgres"
	user     = "postgres"
	password = "admin"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	utils.ErrorLog("Database connection", err)

	return db
}
