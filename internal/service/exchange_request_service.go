package service

import (
	"errors"

	"github.com/Maksatus123/go-final-project/internal/models"
	"github.com/Maksatus123/go-final-project/internal/repository"
)

type ExchangeRequestService struct {
	bookRepo    *repository.BookRepository
	requestRepo *repository.ExchangeRequestRepository
}

func NewExchangeRequestService(bookRepo *repository.BookRepository, requestRepo *repository.ExchangeRequestRepository) *ExchangeRequestService {
	return &ExchangeRequestService{bookRepo: bookRepo, requestRepo: requestRepo}
}

func (s *ExchangeRequestService) CreateExchangeRequest(request *models.ExchangeRequest, userID int) error {
	offeredBook, err := s.bookRepo.GetByID(request.OfferedBookID)
	if err != nil {
		return err
	}
	if offeredBook.OwnerID != userID {
		return errors.New("offered book does not belong to requester")
	}

	requestedBook, err := s.bookRepo.GetByID(request.RequestedBookID)
	if err != nil {
		return err
	}
	if requestedBook.OwnerID == userID {
		return errors.New("cannot request your own book")
	}

	request.RequesterID = userID
	return s.requestRepo.Create(request)
}

func (s *ExchangeRequestService) UpdateExchangeRequestStatus(id int, status string, userID int) error {
	req, err := s.requestRepo.GetByID(id)
	if err != nil {
		return err
	}

	requestedBook, err := s.bookRepo.GetByID(req.RequestedBookID)
	if err != nil {
		return err
	}
	if requestedBook.OwnerID != userID {
		return errors.New("unauthorized: not the owner of the requested book")
	}

	if status == "accepted" {
		offeredBook, err := s.bookRepo.GetByID(req.OfferedBookID)
		if err != nil {
			return err
		}
		// Swap owners
		offeredBook.OwnerID = requestedBook.OwnerID
		requestedBook.OwnerID = req.RequesterID
		err = s.bookRepo.Update(offeredBook)
		if err != nil {
			return err
		}
		err = s.bookRepo.Update(requestedBook)
		if err != nil {
			return err
		}
	}

	req.Status = status
	return s.requestRepo.Update(req)
}

func (s *ExchangeRequestService) GetExchangeRequestByID(id int) (*models.ExchangeRequest, error) {
	return s.requestRepo.GetByID(id)
}

func (s *ExchangeRequestService) GetExchangeRequestsByRequester(requesterID int) ([]*models.ExchangeRequest, error) {
	return s.requestRepo.GetByRequester(requesterID)
}