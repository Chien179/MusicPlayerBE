-- name: GetSongGenre :one
SELECT * FROM songs_genres
WHERE id = $1 LIMIT 1;

-- name: ListSongsGenres :many
SELECT * FROM songs_genres
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateSongGenre :one
INSERT INTO songs_genres (
  songs_id,
  genres_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteSongGenre :exec
DELETE FROM songs_genres
WHERE id = $1;