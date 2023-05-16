-- name: GetUserPlaylists :many
SELECT * FROM playlists
WHERE users_id = $1
ORDER BY id;

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

-- name: GetUserPlaylistSongs :many
SELECT
  s.*
FROM
  playlists_songs ps
  JOIN songs s ON ps.songs_id = s.id
WHERE
  playlists_id = $1
ORDER BY 
  ps.id ASC;

-- name: GetUserPlaylist :one
SELECT * FROM playlists
WHERE id = $1
LIMIT 1;

-- name: AddSongToPlaylist :one
INSERT INTO playlists_songs (
  playlists_id,
  songs_id
) VALUES (
  $1, $2
) 
RETURNING *;

-- name: RemoveSongFromPlaylist :exec
DELETE FROM playlists_songs
WHERE playlists_id = $1 
AND songs_id = $2;

-- name: GetPlaylistSongWithOffset :one
SELECT
  s.*
FROM
  songs s
  JOIN playlists_songs ps ON ps.songs_id = s.id
WHERE
  playlists_id = $1
LIMIT 1
OFFSET $2;

-- name: GetRandomPlaylistSong :one
SELECT
  s.*
FROM
  songs s
  JOIN playlists_songs ps ON ps.songs_id = s.id
WHERE
  playlists_id = $1
AND
  s.id != $2
ORDER BY RANDOM()
LIMIT 1;