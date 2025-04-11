package repository

import (
	"database/sql"

	"github.com/Maksatus123/go-final-project/internal/models"
)

type ExchangeRequestRepository struct {
	db *sql.DB
}

func NewExchangeRequestRepository(db *sql.DB) *ExchangeRequestRepository {
	return &ExchangeRequestRepository{db: db}
}

func (r *ExchangeRequestRepository) Create(request *models.ExchangeRequest) error {
	query := `INSERT INTO exchange_requests (requester_id, requested_book_id, offered_book_id, status) 
              VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, request.RequesterID, request.RequestedBookID, request.OfferedBookID, "pending").Scan(&request.ID)
}

func (r *ExchangeRequestRepository) GetByID(id int) (*models.ExchangeRequest, error) {
	request := &models.ExchangeRequest{}
	query := `SELECT id, requester_id, requested_book_id, offered_book_id, status FROM exchange_requests WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&request.ID, &request.RequesterID, &request.RequestedBookID, &request.OfferedBookID, &request.Status)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *ExchangeRequestRepository) GetByRequester(requesterID int) ([]*models.ExchangeRequest, error) {
	query := `SELECT id, requester_id, requested_book_id, offered_book_id, status FROM exchange_requests WHERE requester_id = $1`
	rows, err := r.db.Query(query, requesterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*models.ExchangeRequest
	for rows.Next() {
		req := &models.ExchangeRequest{}
		err := rows.Scan(&req.ID, &req.RequesterID, &req.RequestedBookID, &req.OfferedBookID, &req.Status)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}

func (r *ExchangeRequestRepository) Update(request *models.ExchangeRequest) error {
	query := `UPDATE exchange_requests SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, request.Status, request.ID)
	return err
}