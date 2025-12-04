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
		log.Fatal("PostgreSQL কানেক্ট করতে পারিনি: ", err)
	}

	DB = db
	log.Println("PostgreSQL সফলভাবে কানেক্ট হয়েছে!")

	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("টেবিল মাইগ্রেট করতে পারিনি: ", err)
	}
	log.Println("tasks টেবিল তৈরি/আপডেট হয়েছে")
}
