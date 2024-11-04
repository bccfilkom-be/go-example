-- name: ListPets :many
SELECT
  id,
  name,
  photo_url,
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
  name,
  photo_url,
  sold
FROM
  pets
WHERE
  id = $1;

-- name: CreatePet :one
INSERT INTO
  pets (name, photo_url) VALUES ($1, $2) RETURNING id;

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
