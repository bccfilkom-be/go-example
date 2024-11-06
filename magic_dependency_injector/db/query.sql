-- name: ListBooks :many
SELECT
  *
FROM
  books;

-- name: CreatePet :one
INSERT INTO
  books (title) VALUES ($1) RETURNING id;
