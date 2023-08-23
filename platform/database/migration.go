package database

import (
	"log"
	"github.com/eliasdn/Elda-Bank-API/app/models"
)

func CheckDB() error {
	db, err := PostgreSQLConnection()
	utils.CheckErr(err)
	//log.Print(db.Migrator().GetTables())

	if !db.Migrator().HasTable(&models.User{}) {
		err := db.Migrator().AutoMigrate(&models.User{})
		if err != nil {
			log.Print(err)
			return err
		}
	}

	if !db.Migrator().HasTable(&models.Account{}) {
		err := db.Migrator().AutoMigrate(&models.Account{})
		if err != nil {
			log.Print(err)
			return err
		}
	}

	if !db.Migrator().HasTable(&models.Customer{}) {
		err := db.Migrator().AutoMigrate(&models.Customer{})
		if err != nil {
			log.Print(err)
			return err
		}
	}

	return nil
}
