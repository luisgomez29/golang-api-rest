package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	PORT      = "8000"
	DBURL     = ""
	SECRETKEY []byte
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT = os.Getenv("API_PORT")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
	DBURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		os.Getenv("HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
}
