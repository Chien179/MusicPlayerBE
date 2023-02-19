-- name: GetSong :one
SELECT * FROM songs
WHERE id = $1 LIMIT 1;

-- name: ListSongs :many
SELECT * FROM songs
ORDER BY id;

-- name: CreateSong :one
INSERT INTO songs (
  name,
  singer,
  image,
  file_url,
  duration
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteSong :exec
DELETE FROM songs
WHERE id = $1;

-- name: UpdateSong :one
UPDATE songs
SET name = $2,
    singer = $3,
    image = $4,
    file_url = $5,
    duration = $6
WHERE id = $1
RETURNING *;