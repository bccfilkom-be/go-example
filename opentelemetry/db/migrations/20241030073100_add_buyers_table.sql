-- +goose Up
-- +goose StatementBegin
CREATE TABLE "buyers" (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL REFERENCES "users" (id),
  username TEXT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "buyers";

-- +goose StatementEnd
