-- +goose Up
-- +goose StatementBegin
CREATE TABLE "orders" (
  id BIGSERIAL PRIMARY KEY,
  buyer_id BIGSERIAL REFERENCES "buyers" (id),
  status order_status NOT NULL,
  description TEXT,
  shipping_address TEXT NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "orders";

-- +goose StatementEnd
