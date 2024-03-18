-- +goose Up
CREATE TABLE pets (
                      id SERIAL PRIMARY KEY,
                      category_id INT NOT NULL,
                      photo_urls TEXT[] NOT NULL,
                      name TEXT NOT NULL,
                      status TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS pets;
