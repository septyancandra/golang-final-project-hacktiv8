package config

import (
	"Mygram/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed load file env")
	} else {
		fmt.Println("successfully read file env")
	}

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Conn Failed!", err)
	}

	fmt.Println("DB Conn succeed")
	db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Comment{}, models.Photo{})

}

func GetDB() *gorm.DB {
	return db
}
