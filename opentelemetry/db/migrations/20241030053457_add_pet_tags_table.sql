-- +goose Up
-- +goose StatementBegin
CREATE TABLE "pet_tags" (
  pet_id BIGSERIAL REFERENCES "pets" (id) ON DELETE CASCADE,
  tag_id BIGSERIAL REFERENCES "tags" (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE "pet_tags";

-- +goose StatementEnd
