// internal/database/db.go
package database

import (
	"log"
	"task-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=go_user password=123456 dbname=taskdb port=5432 sslmode=disable TimeZone=Asia/Dhaka"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("PostgreSQL not cannecting: ", err)
	}

	DB = db
	log.Println("PostgreSQL is connected")

	if err := DB.AutoMigrate(&models.Task{}, &models.Product{}); err != nil {
		log.Fatal("table not insert: ", err)
	}
	log.Println("tables are created")
}
