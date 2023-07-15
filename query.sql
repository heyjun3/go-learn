-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: CreateAuthors :exec
INSERT INTO authors
SELECT unnest(@ids::bigint[]) AS id,
  unnest(@names::text[]) AS name,
  unnest(@bios::text[]) AS bio;
