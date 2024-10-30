-- +goose Up
-- +goose StatementBegin
CREATE TYPE "order_status" AS ENUM ('Submitted', 'Paid', 'Shipped', 'Cancelled');

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE "order_status";

-- +goose StatementEnd
