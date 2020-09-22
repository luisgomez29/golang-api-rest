package database

import (
	"fmt"
	"github.com/luisgomez29/golang-api-rest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	fmt.Println("DB=>", config.DBURL)
	db, err := gorm.Open(postgres.Open(config.DBURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
