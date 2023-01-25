package database

import (
	"log"
	"uns/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.Migrator().DropTable("students")
	Instance.Migrator().DropTable("profesors")
	Instance.AutoMigrate(&models.Student{})
	Instance.AutoMigrate(&models.Profesor{})
	log.Println("Database Migration Completed!")
}
