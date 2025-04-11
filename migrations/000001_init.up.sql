CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    owner_id INTEGER NOT NULL
);

CREATE TABLE exchange_requests (
    id SERIAL PRIMARY KEY,
    requester_id INTEGER NOT NULL,
    requested_book_id INTEGER NOT NULL,
    offered_book_id INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    FOREIGN KEY (requested_book_id) REFERENCES books(id),
    FOREIGN KEY (offered_book_id) REFERENCES books(id)
);