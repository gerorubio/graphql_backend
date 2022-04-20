package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=xdxd1234 dbname=sedemaTest port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&Tipos_Residuos{})

	return db
}
