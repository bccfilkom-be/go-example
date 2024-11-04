-- +goose Up
-- +goose StatementBegin
CREATE TABLE "baskets" (
  id BIGSERIAL PRIMARY KEY,
  buyer_id BIGSERIAL REFERENCES "buyers" (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "baskets";

-- +goose StatementEnd
