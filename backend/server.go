package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DB_Conection() {

	var err error
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Failed to load env file")
	}

	var DSN = "host=db user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	// In case of error, print it and exit
	if err != nil {

		log.Fatal(err)

	} else {

		log.Println("Database connected")
	}

}
