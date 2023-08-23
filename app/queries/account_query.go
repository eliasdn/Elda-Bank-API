package queries

import (
	"github.com/eliasdn/Elda-Bank-API/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AccountQueries struct for queries from Book model.
type AccountQueries struct {
	*gorm.DB
}

// GetAccounts method for getting all accounts.
func (q *AccountQueries) GetAccounts() ([]models.Account, error) {
	// Define books variable.
	accounts := []models.Account{}

	// Define query string.
	//query := `SELECT * FROM books`

	// Send query to database.
	err := q.Find(&accounts)
	if err != nil {
		// Return empty object and error.
		return accounts, err.Error
	}

	// Return query result.
	return accounts, nil
}

// GetAccountsByUser method for getting all accounts by given user.
func (q *AccountQueries) GetAccountsByUser(userId uuid.UUID) ([]models.Account, error) {
	// Define books variable.
	accounts := []models.Account{}

	// Define query string.
	//query := `SELECT * FROM books WHERE author = $1`

	// Send query to database.
	err := q.Where("userid = ?", userId).Find(&accounts)
	if err != nil {
		// Return empty object and error.
		return accounts, err.Error
	}

	// Return query result.
	return accounts, nil
}

// GetAccount method for getting one account by given ID.
func (q *AccountQueries) GetAccount(id uuid.UUID) (models.Account, error) {
	// Define book variable.
	account := models.Account{}

	// Define query string.
	//query := `SELECT * FROM accounts WHERE id = $1`

	// Send query to database.
	err := q.First(&account, id)
	if err != nil {
		// Return empty object and error.
		return account, err.Error
	}

	// Return query result.
	return account, nil
}

// CreateBook method for creating book by given Book object.
func (q *AccountQueries) CreateAccount(b *models.Account) error {
	// Define query string.
	//query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	query := q.Create(b)
	if query.Error != nil {
		// Return only error.
		return query.Error
	}

	// This query returns nothing.
	return nil
}

// UpdateAccount method for updating account by given Account object.
func (q *AccountQueries) UpdateAccount(id uuid.UUID, b *models.Account) error {
	// Define query string.
	//query := `UPDATE accounts SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	// Send query to database.
	query := q.Where("id = ?", id).Save(b)
	if query.Error != nil {
		// Return only error.
		return query.Error
	}

	// This query returns nothing.
	return nil
}

// DeleteAccount method for delete account by given ID.
func (q *AccountQueries) DeleteAccount(id uuid.UUID) error {
	account := models.Account{}
	// Define query string.
	//query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	query := q.Where("id = ?", id).Delete(&account)
	if query.Error != nil {
		// Return only error.
		return query.Error
	}

	// This query returns nothing.
	return nil
}
