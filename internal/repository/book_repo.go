package repository

import (
	"database/sql"

	"github.com/Maksatus123/go-final-project/internal/models"
	_ "github.com/lib/pq"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `INSERT INTO books (title, author, genre, owner_id) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, book.Title, book.Author, book.Genre, book.OwnerID).Scan(&book.ID)
}

func (r *BookRepository) GetByID(id int) (*models.Book, error) {
	book := &models.Book{}
	query := `SELECT id, title, author, genre, owner_id FROM books WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.OwnerID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) GetAll() ([]*models.Book, error) {
	query := `SELECT id, title, author, genre, owner_id FROM books`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.OwnerID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookRepository) Update(book *models.Book) error {
	query := `UPDATE books SET title = $1, author = $2, genre = $3, owner_id = $4 WHERE id = $5`
	_, err := r.db.Exec(query, book.Title, book.Author, book.Genre, book.OwnerID, book.ID)
	return err
}

func (r *BookRepository) Delete(id int, ownerID int) error {
	query := `DELETE FROM books WHERE id = $1 AND owner_id = $2`
	_, err := r.db.Exec(query, id, ownerID)
	return err
}

func (r *BookRepository) GetByOwner(ownerID int) ([]*models.Book, error) {
	query := `SELECT id, title, author, genre, owner_id FROM books WHERE owner_id = $1`
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.OwnerID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}