package queries

import (
	"books-api/app/models"

	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books`

	// Send query to database.
	err := q.Get(&books, query)
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	// Return query result.
	return books, nil
}
