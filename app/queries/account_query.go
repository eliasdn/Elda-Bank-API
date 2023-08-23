package queries

import (
	"github.com/eliasdn/Elda-Bank-API/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AccountQueries struct for queries from account model.
type AccountQueries struct {
	*gorm.DB
}

// GetAccounts method for getting all accounts.
func (q *AccountQueries) GetAccounts() ([]models.Account, error) {
	accounts := []models.Account{}
	err := q.Find(&accounts)
	if err != nil {
		return accounts, err.Error
	}
	return accounts, nil
}

// GetAccountsByUser method for getting all accounts by given user.
func (q *AccountQueries) GetAccountsByUser(userId uuid.UUID) ([]models.Account, error) {
	accounts := []models.Account{}
	err := q.Where("userid = ?", userId).Find(&accounts)
	if err != nil {
		return accounts, err.Error
	}
	return accounts, nil
}

// GetAccount method for getting one account by given ID.
func (q *AccountQueries) GetAccount(id uuid.UUID) (models.Account, error) {
	account := models.Account{}
	err := q.First(&account, id)
	if err != nil {
		return account, err.Error
	}
	return account, nil
}

// Createaccount method for creating account by given account object.
func (q *AccountQueries) CreateAccount(b *models.Account) error {
	query := q.Create(b)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// UpdateAccount method for updating account by given Account object.
func (q *AccountQueries) UpdateAccount(id uuid.UUID, b *models.Account) error {
	query := q.Where("id = ?", id).Save(b)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

// DeleteAccount method for delete account by given ID.
func (q *AccountQueries) DeleteAccount(id uuid.UUID) error {
	account := models.Account{}
	query := q.Where("id = ?", id).Delete(&account)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
