-- +goose Up
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        pet_id INT NOT NULL,
                        quantity INT NOT NULL,
                        ship_date TIMESTAMP NOT NULL,
                        status TEXT NOT NULL,
                        complete BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS orders;
