package db

import (
	"os"

	"github.com/jotace1/simple-authentication/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	db.AutoMigrate(&entity.Account{})

	if err != nil {
		panic(err)
	}

	return db
}
