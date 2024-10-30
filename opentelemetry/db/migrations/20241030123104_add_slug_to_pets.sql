-- +goose Up
-- +goose StatementBegin
ALTER TABLE pets
ADD COLUMN slug TEXT UNIQUE NOT NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE pets
DROP COLUMN slug;

-- +goose StatementEnd
