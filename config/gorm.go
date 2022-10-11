package config

import (
	"fmt"
	"os"
	"sesi7/gin/repository"

	ormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBORM *gorm.DB

func ConnectGorm() error {
	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	user = os.Getenv("POSTGRES_USER")
	pass = os.Getenv("POSTGRES_PASS")
	dbname = os.Getenv("POSTGRES_DBNAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := gorm.Open(ormPostgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if !db.Migrator().HasTable(repository.Product{}) {
		db.Debug().AutoMigrate(repository.Product{})
	}

	if !db.Migrator().HasColumn(repository.Product{}, "name") {
		db.Debug().AutoMigrate(repository.Product{})

	}

	db.Debug().AutoMigrate(repository.User{})

	// db.Debug().AutoMigrate(r, repository.Person{})

	DBORM = db
	return nil
}

func GetGORM() *gorm.DB {
	return DBORM
}
