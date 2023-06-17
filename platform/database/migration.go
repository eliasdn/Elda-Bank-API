package database

import (
	"log"

	"github.com/eliasdn/Elda-Bank-API/app/models"
)

func checkDB() error {
	db, err := PostgreSQLConnection()
	if err != nil {
		log.Print(err)
		return err
	}

	if db.Migrator().HasTable(&models.User{}) {
		db.Migrator().CreateTable(&models.User{})
	}

	if db.Migrator().HasTable(&models.Account{}) {
		db.Migrator().CreateTable(&models.Account{})
	}

	if db.Migrator().HasTable(&models.Customer{}) {
		db.Migrator().CreateTable(&models.Customer{})
	}
	return nil
}
