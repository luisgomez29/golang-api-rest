package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DBURL = ""
	PORT  = ":8000"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT = os.Getenv("API_PORT")
	DBURL = fmt.Sprintf("host=%v user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Bogota",
		os.Getenv("HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"))
}
