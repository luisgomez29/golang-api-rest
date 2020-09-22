package auto

import (
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/utils"
	"gorm.io/gorm"
	"log"
)

func Load(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// Insertar datos de prueba
	db.Create(&users)

	// Mostrar en consola datos insertados
	for _, user := range users {
		utils.Pretty(user)
	}
}
