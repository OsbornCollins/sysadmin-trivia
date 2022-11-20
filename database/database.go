// Filename: database/database.go

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/OsbornCollins/sysadmin-trivia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Create gorm struct
// GORM is an Object Relational Mapping (ORM) library for Golang. ORM converts data between
// incompatible type systems using object-oriented programming languages.
type Dbinstance struct {
	Db *gorm.DB
}

// Create a new package-level variable to hold our global database
var DB Dbinstance

// ConnectDb() function is used to connect our database to the app
func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Belize",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// Use GORM method open()
	// gorm.Open() takes two arguments.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// AutoMigrate creates tables that we need. We pass gorm Fact models to AutoMigrate so the table
	// can be created in DB.
	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	// Set the value of our global DB variable to the database we just set up.
	DB = Dbinstance{
		Db: db,
	}
}
