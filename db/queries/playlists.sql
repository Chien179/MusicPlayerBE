-- name: GetPlaylist :one
SELECT * FROM playlists
WHERE id = $1 LIMIT 1;

-- name: ListPlaylists :many
SELECT * FROM playlists
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreatePlaylist :one
INSERT INTO playlists (
  users_id,
  name,
  image
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeletePlaylist :exec
DELETE FROM playlists
WHERE id = $1;

-- name: UpdatePlaylist :one
UPDATE playlists
SET name = $2,
    image = $3
WHERE id = $1
RETURNING *;