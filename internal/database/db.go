package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"project-hive/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("hive.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.Project{},
		&models.Sprint{},
		&models.Issue{},
	)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
