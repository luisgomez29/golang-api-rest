package auto

import (
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/utils"
	"gorm.io/gorm"
	"log"
)

func Load(db *gorm.DB) {
	mdl := []interface{}{&models.User{}, &models.Product{}}
	err := db.Migrator().DropTable(mdl...)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(mdl...)
	if err != nil {
		log.Fatal(err)
	}

	// Insertar datos de prueba
	db.Create(&users)

	// Mostrar en consola datos insertados
	for _, user := range users {
		utils.Pretty(user)
	}

	db.Create(&products)
	// Mostrar en consola datos insertados
	for _, product := range products {
		utils.Pretty(product)
	}
}
