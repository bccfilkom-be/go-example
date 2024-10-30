-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  username TEXT UNIQUE NOT NULL,
  provider oauth_provider NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";

-- +goose StatementEnd
