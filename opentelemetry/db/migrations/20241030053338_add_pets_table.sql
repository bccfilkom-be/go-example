-- +goose Up
-- +goose StatementBegin
CREATE TABLE "pets" (
  id BIGSERIAL PRIMARY KEY,
  -- category_id BIGSERIAL REFERENCES "categories" (id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  photo_url TEXT NOT NULL,
  sold BOOLEAN DEFAULT FALSE NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "pets";

-- +goose StatementEnd
