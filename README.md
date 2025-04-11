# Book Service

Microservice for managing books and exchange requests in the BookSwap Platform.

## Prerequisites

- Go 1.21
- Docker & Docker Compose
- `migrate` CLI (install via `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/Maksatus123/go-final-project
   cd book-service
   ```

1. Install dependencies:

   ```bash
   go mod download
   ```

1. Run locally with Docker Compose (from parent folder):
   ```bash
   cd ..
   make up
   ```

## API Endpoints

- POST /books: Create a book (authenticated)
- GET /books/:id: Get a book by ID
- GET /books: List all books
- PUT /books/:id: Update a book (authenticated, owner only)
- DELETE /books/:id: Delete a book (authenticated, owner only)
- GET /my-books?owner_id=<id>: Get books by owner ID
- POST /exchange-requests: Create an exchange request (authenticated)
- GET /exchange-requests/:id: Get an exchange request by ID
- GET /exchange-requests: List user's exchange requests (authenticated)
- PUT /exchange-requests/:id: Update exchange request status (authenticated, owner of requested book only)
