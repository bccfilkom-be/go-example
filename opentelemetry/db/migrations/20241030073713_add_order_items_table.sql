-- +goose Up
-- +goose StatementBegin
CREATE TABLE "order_items" (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  photo_url TEXT NOT NULL,
  category TEXT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "oder_items";

-- +goose StatementEnd
