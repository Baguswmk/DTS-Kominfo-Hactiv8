package database

import (
	"dts/Project1/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host		= "localhost"
	user		= "postgres"
	password	= "admin"
	dbPort		= 5432
	dbName		= "db-go-gorm"
	db 			*gorm.DB
	err			 error	
)

func ConnectDB(){
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connect to database ", err)
	}

	db.Debug().AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}