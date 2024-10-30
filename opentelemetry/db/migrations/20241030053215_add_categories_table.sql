-- +goose Up
-- +goose StatementBegin
CREATE TABLE "categories" (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "categories";

-- +goose StatementEnd
