-- +goose Up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name TEXT NOT NULL,
                       last_name TEXT NOT NULL,
                       email TEXT NOT NULL,
                       password TEXT NOT NULL,
                       phone TEXT NOT NULL,
                       user_status INT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
