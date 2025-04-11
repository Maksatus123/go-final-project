package models

type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Genre   string `json:"genre"`
	OwnerID int    `json:"owner_id"`
}

type ExchangeRequest struct {
	ID              int    `json:"id"`
	RequesterID     int    `json:"requester_id"`
	RequestedBookID int    `json:"requested_book_id"`
	OfferedBookID   int    `json:"offered_book_id"`
	Status          string `json:"status"` // "pending", "accepted", "rejected"
}