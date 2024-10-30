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
  @_limit::integer
OFFSET
  @_offset::integer;

-- name: CountPets :one
SELECT
  COUNT(*)
FROM
  pets;
