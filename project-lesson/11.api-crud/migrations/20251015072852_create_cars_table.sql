-- +goose Up
CREATE TABLE cars (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE cars;
