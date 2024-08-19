-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS card_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
