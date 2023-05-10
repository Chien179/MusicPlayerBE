-- name: GetSong :one
SELECT * FROM songs
WHERE id = $1 LIMIT 1;

-- name: GetSongs :many
SELECT * FROM songs
ORDER BY id
LIMIT $1
OFFSET $2;

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
SET name = COALESCE($2, name),
    singer = COALESCE($3, singer),
    image = COALESCE($4, image),
    file_url = COALESCE($5, file_url),
    duration = COALESCE($6, duration)
WHERE id = $1
RETURNING *;

-- name: UpdateSongFile :one
UPDATE songs
SET file_url = $2,
    image = $3
WHERE id = $1
RETURNING *;