-- name: GetSongGenre :one
SELECT * FROM songs_genres
WHERE id = $1 LIMIT 1;

-- name: ListSongsGenres :many
SELECT genres_id FROM songs_genres
WHERE songs_id = $1
ORDER BY id;

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
WHERE genres_id = $1
AND songs_id = $2;