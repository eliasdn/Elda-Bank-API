package utils

import (
	"github.com/eliasdn/Elda-Bank-API/app/models"
)

func AutoMigrate() {
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.SignUp{}, &models.SignIn{}, &models.Customer{})
}
