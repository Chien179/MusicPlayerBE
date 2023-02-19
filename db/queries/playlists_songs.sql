-- name: GetPlaylistSong :one
SELECT * FROM playlists_songs
WHERE id = $1 LIMIT 1;

-- name: ListPlaylistsSongs :many
SELECT * FROM playlists_songs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreatePlaylistSong :one
INSERT INTO playlists_songs (
  playlists_id,
  songs_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeletePlaylistSong :exec
DELETE FROM playlists_songs
WHERE id = $1;