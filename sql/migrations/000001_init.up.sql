CREATE TABLE expenses (
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    category TEXT NOT NULL
);

CREATE TABLE users (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    expiration_date TIMESTAMP NOT NULL
);
