-- name: ListPets :many
SELECT
  id,
  category_id,
  name,
  photoURL,
  sold
FROM
  pets
LIMIT
  "$1::integer"
OFFSET
  "$2::integer";

-- name: CountPets :one
SELECT
  COUNT(*)
FROM
  pets;
