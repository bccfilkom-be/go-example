-- +goose Up
-- +goose StatementBegin
CREATE TABLE "tags" (id SERIAL PRIMARY KEY, name TEXT NOT NULL);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "tags";

-- +goose StatementEnd
