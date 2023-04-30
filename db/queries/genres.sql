-- name: GetGenre :one
SELECT * FROM genres
WHERE id = $1 LIMIT 1;

-- name: GetGenres :many
SELECT * FROM genres
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateGenre :one
INSERT INTO genres (
  name,
  image
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = $1;

-- name: Updategenre :one
UPDATE genres
SET name = $2,
    image = $3
WHERE id = $1
RETURNING *;