package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/luisgomez29/golang-api-rest/config"
	"log"
	"os"
)

func Connect() *pg.DB {
	fmt.Println("DB=>", config.DBURL)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     fmt.Sprintf("%s", os.Getenv("DB_USER")),
		Password: fmt.Sprintf("%s", os.Getenv("DB_PWD")),
		Database: fmt.Sprintf("%s", os.Getenv("DB_NAME")),
	})
}
