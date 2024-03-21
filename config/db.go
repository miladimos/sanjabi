package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgresDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
		panic("Failed to connect to postgres")
	}

	// db.AutoMigrate()

	fmt.Println("Connected to postgres")

	return db
}

func ClosePostgresDB(db *gorm.DB) {
	dbPspl, err := db.DB()
	if err != nil {
		log.Fatal("Failed. postgres connection Error: ", err)
		panic("Failed. postgres connection Error:")
	}

	err = dbPspl.Close()
	if err != nil {
		log.Fatal("Failed: unable to close postgres ", err)
		panic("Failed: unable to close postgres")
	}
}
