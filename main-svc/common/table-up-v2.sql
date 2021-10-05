-- SQLite

CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE quotes(
    id INTEGER PRIMARY KEY,
    body TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    quote_source VARCHAR(255) NOT NULL
);

CREATE TABLE userfavoritesquotes(
    user_id INTEGER,
    quote_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(quote_id) REFERENCES quotes(id)
);