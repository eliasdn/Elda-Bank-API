package queries

import (
	"github.com/eliasdn/Elda-Bank-API/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*gorm.DB
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	//query := `SELECT * FROM users WHERE id = $1`

	// Send query to database.
	query := q.First(&user, id)
	if query.Error != nil {
		// Return empty object and error.
		return user, query.Error
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Send query to database.
	// `SELECT * FROM users WHERE email = $1`
	query := q.Where("email = ?", email).First(&user)
	if query.Error != nil {
		// Return empty object and error.
		return user, query.Error
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByUsername(username string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Send query to database.
	// `SELECT * FROM users WHERE email = $1`
	query := q.Where("username = ?", username).First(&user)
	if query.Error != nil {
		// Return empty object and error.
		return user, query.Error
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(user *models.User) error {
	// Define query string.
	//query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	query := q.Create(&user)
	if query.Error != nil {
		// Return only error.
		return query.Error
	}

	// This query returns nothing.
	return nil
}
