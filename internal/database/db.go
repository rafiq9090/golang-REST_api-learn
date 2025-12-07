// internal/database/db.go
package database

import (
	"fmt"
	"log"
	"task-api/internal/config"
	"task-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "host=localhost user=go_user password=123456 dbname=taskdb port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dhaka",
		config.App.DBHost,
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBName,
		config.App.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("PostgreSQL not cannecting: ", err)
	}

	DB = db
	log.Println("PostgreSQL is connected")

	if err := DB.AutoMigrate(&models.Task{}, &models.Product{}, &models.User{}); err != nil {
		log.Fatal("table not insert: ", err)
	}
	log.Println("tables are created")
}
