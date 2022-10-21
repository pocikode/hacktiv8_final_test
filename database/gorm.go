package database

import (
	"MyGram/server/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%v/%v", os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_NAME"))

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Printf("Error connecting database\n%v", err)
		panic(err)
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.Photo{})
	db.Debug().AutoMigrate(models.SocialMedia{})
	db.Debug().AutoMigrate(models.Comment{})
}
