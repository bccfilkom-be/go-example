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

-- name: GetPet :one
SELECT
  id,
  category_id,
  name,
  photoURL,
  sold
FROM
  pets
WHERE
  id = $1;

-- name: CreatePet :one
INSERT INTO
  pets (name, photoURL) VALUES ($1, $2) RETURNING id;

-- name: UpdatePet :exec
UPDATE pets
SET
  name = $2
WHERE
  id = $1;

-- name: DeletePet :exec
DELETE FROM pets
WHERE
  id = $1;
