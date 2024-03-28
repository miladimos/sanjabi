package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgresDB() *gorm.DB {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	// dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
		panic("Failed to connect to postgres")
	}

	// db.AutoMigrate(&User{})

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
