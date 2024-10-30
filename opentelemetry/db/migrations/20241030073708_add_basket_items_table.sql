-- +goose Up
-- +goose StatementBegin
CREATE TABLE "basket_items" (
  id BIGSERIAL PRIMARY KEY,
  basket_id BIGSERIAL REFERENCES "baskets" (id) ON DELETE CASCADE,
  pet_id BIGSERIAL REFERENCES "pets" (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "basket_items";

-- +goose StatementEnd
