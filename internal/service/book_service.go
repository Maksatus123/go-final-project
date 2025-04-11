package service

import (
	"errors"

	"github.com/Maksatus123/go-final-project/internal/models"
	"github.com/Maksatus123/go-final-project/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *models.Book, userID int) error {
	book.OwnerID = userID
	return s.repo.Create(book)
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) UpdateBook(book *models.Book, userID int) error {
	existing, err := s.repo.GetByID(book.ID)
	if err != nil {
		return err
	}
	if existing.OwnerID != userID {
		return errors.New("unauthorized: not the owner")
	}
	book.OwnerID = userID
	return s.repo.Update(book)
}

func (s *BookService) DeleteBook(id int, userID int) error {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if book.OwnerID != userID {
		return errors.New("unauthorized: not the owner")
	}
	return s.repo.Delete(id, userID)
}

func (s *BookService) GetBooksByOwner(ownerID int) ([]*models.Book, error) {
	return s.repo.GetByOwner(ownerID)
}