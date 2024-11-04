-- +goose Up
-- +goose StatementBegin
CREATE TYPE "oauth_provider" AS ENUM ('google');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE "oauth_provider";

-- +goose StatementEnd
