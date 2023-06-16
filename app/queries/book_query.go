package queries

import (
	"github.com/eliasdn/fiberAPI-template/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*gorm.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	//query := `SELECT * FROM books`

	// Send query to database.
	err := q.Find(&books)
	if err != nil {
		// Return empty object and error.
		return books, err.Error
	}

	// Return query result.
	return books, nil
}

// GetBooksByAuthor method for getting all books by given author.
func (q *BookQueries) GetBooksByAuthor(author string) ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	//query := `SELECT * FROM books WHERE author = $1`

	// Send query to database.
	err := q.Where("author = ?", author).Find(&books)
	if err != nil {
		// Return empty object and error.
		return books, err.Error
	}

	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Define query string.
	//query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	err := q.First(&book, id)
	if err != nil {
		// Return empty object and error.
		return book, err.Error
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error {
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

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	// Define query string.
	//query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	// Send query to database.
	query := q.Where("id = ?", id).Save(b)
	if query.Error != nil {
		// Return only error.
		return query.Error
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	book := models.Book{}
	// Define query string.
	//query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	query := q.Where("id = ?", id).Delete(&book)
	if query.Error != nil {
		// Return only error.
		return query.Error
	}

	// This query returns nothing.
	return nil
}
