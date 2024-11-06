package db

import (
	"log"
	"miniProject/models"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{
		{
			Title:       "The Lord of the Rings",
			Author:      "J. R. R. Tolkien",
			Description: "The Lord of the Rings is a series of three epic fantasy novels by English author and scholar J. R. R. Tolkien.",
			Stock:       10,
		},
		{
			Title:       "Harry Potter",
			Author:      "J. K. Rowling",
			Description: "Harry Potter is a series of seven fantasy novels written by British author J. K. Rowling.",
			Stock:       5,
		},
		{
			Title:       "To Kill a Mockingbird",
			Author:      "Harper Lee",
			Description: "To Kill a Mockingbird is a novel by American writer Harper Lee.",
			Stock:       20,
		},
		{
			Title:       "Pride and Prejudice",
			Author:      "Jane Austen",
			Description: "Pride and Prejudice is a novel by English Author: Firstname Lastname.",
			Stock:       15,
		},
		{
			Title:       "The Catcher in the Rye",
			Author:      "J. D. Salinger",
			Description: "The Catcher in the Rye is a novel by J. D. Salinger.",
			Stock:       10,
		},
	}

	for _, book := range data {
		db.Create(&book)
	}
}