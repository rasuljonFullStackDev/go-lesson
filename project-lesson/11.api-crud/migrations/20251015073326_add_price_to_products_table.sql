-- +goose Up
ALTER TABLE products
ADD COLUMN new_column_name VARCHAR(255);

-- +goose Down
ALTER TABLE products
DROP COLUMN new_column_name;
