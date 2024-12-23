package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/oziev02/url-shortener/internal/link"
	"github.com/oziev02/url-shortener/internal/stat"
	"github.com/oziev02/url-shortener/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
}
