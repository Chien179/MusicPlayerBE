// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: playlists.sql

package db

import (
	"context"
)

const addSongToPlaylist = `-- name: AddSongToPlaylist :one
INSERT INTO playlists_songs (
  playlists_id,
  songs_id
) VALUES (
  $1, $2
) 
RETURNING id, songs_id, playlists_id
`

type AddSongToPlaylistParams struct {
	PlaylistsID int64 `json:"playlists_id"`
	SongsID     int64 `json:"songs_id"`
}

func (q *Queries) AddSongToPlaylist(ctx context.Context, arg AddSongToPlaylistParams) (PlaylistsSong, error) {
	row := q.db.QueryRowContext(ctx, addSongToPlaylist, arg.PlaylistsID, arg.SongsID)
	var i PlaylistsSong
	err := row.Scan(&i.ID, &i.SongsID, &i.PlaylistsID)
	return i, err
}

const createPlaylist = `-- name: CreatePlaylist :one
INSERT INTO playlists (
  users_id,
  name,
  image
) VALUES (
  $1, $2, $3
)
RETURNING id, users_id, name, image, created_at
`

type CreatePlaylistParams struct {
	UsersID int64  `json:"users_id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
}

func (q *Queries) CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, createPlaylist, arg.UsersID, arg.Name, arg.Image)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.UsersID,
		&i.Name,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}

const deletePlaylist = `-- name: DeletePlaylist :exec
DELETE FROM playlists
WHERE id = $1
`

func (q *Queries) DeletePlaylist(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePlaylist, id)
	return err
}

const getPlaylistSongWithOffset = `-- name: GetPlaylistSongWithOffset :one
SELECT
  s.id,
  s.name,
  s.singer,
  s.image,
  s.file_url,
  s.duration,
  s.created_at
FROM
  songs s
  JOIN playlists_songs ps ON ps.songs_id = s.id
WHERE
  playlists_id = $1
LIMIT 1
OFFSET $2
`

type GetPlaylistSongWithOffsetParams struct {
	PlaylistsID int64 `json:"playlists_id"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) GetPlaylistSongWithOffset(ctx context.Context, arg GetPlaylistSongWithOffsetParams) (Song, error) {
	row := q.db.QueryRowContext(ctx, getPlaylistSongWithOffset, arg.PlaylistsID, arg.Offset)
	var i Song
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Singer,
		&i.Image,
		&i.FileUrl,
		&i.Duration,
		&i.CreatedAt,
	)
	return i, err
}

const getRandomPlaylistSong = `-- name: GetRandomPlaylistSong :one
SELECT
  s.id,
  s.name,
  s.singer,
  s.image,
  s.file_url,
  s.duration,
  s.created_at
FROM
  songs s
  JOIN playlists_songs ps ON ps.songs_id = s.id
WHERE
  playlists_id = $1
AND
  s.id != $2
ORDER BY RANDOM()
LIMIT 1
`

type GetRandomPlaylistSongParams struct {
	PlaylistsID int64 `json:"playlists_id"`
	ID          int64 `json:"id"`
}

func (q *Queries) GetRandomPlaylistSong(ctx context.Context, arg GetRandomPlaylistSongParams) (Song, error) {
	row := q.db.QueryRowContext(ctx, getRandomPlaylistSong, arg.PlaylistsID, arg.ID)
	var i Song
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Singer,
		&i.Image,
		&i.FileUrl,
		&i.Duration,
		&i.CreatedAt,
	)
	return i, err
}

const getUserPlaylist = `-- name: GetUserPlaylist :one
SELECT id, users_id, name, image, created_at FROM playlists
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUserPlaylist(ctx context.Context, id int64) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, getUserPlaylist, id)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.UsersID,
		&i.Name,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}

const getUserPlaylistSongs = `-- name: GetUserPlaylistSongs :many
SELECT
  s.id,
  s.name,
  s.singer,
  s.image,
  s.file_url,
  s.duration,
  s.created_at
FROM
  songs s
  JOIN playlists_songs ps ON ps.songs_id = s.id
WHERE
    playlists_id = $1
ORDER BY s.created_at
`

func (q *Queries) GetUserPlaylistSongs(ctx context.Context, playlistsID int64) ([]Song, error) {
	rows, err := q.db.QueryContext(ctx, getUserPlaylistSongs, playlistsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Song{}
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Singer,
			&i.Image,
			&i.FileUrl,
			&i.Duration,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPlaylists = `-- name: GetUserPlaylists :many
SELECT id, users_id, name, image, created_at FROM playlists
WHERE users_id = $1
ORDER BY id
`

func (q *Queries) GetUserPlaylists(ctx context.Context, usersID int64) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, getUserPlaylists, usersID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Playlist{}
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.ID,
			&i.UsersID,
			&i.Name,
			&i.Image,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeSongFromPlaylist = `-- name: RemoveSongFromPlaylist :exec
DELETE FROM playlists_songs
WHERE playlists_id = $1 
AND songs_id = $2
`

type RemoveSongFromPlaylistParams struct {
	PlaylistsID int64 `json:"playlists_id"`
	SongsID     int64 `json:"songs_id"`
}

func (q *Queries) RemoveSongFromPlaylist(ctx context.Context, arg RemoveSongFromPlaylistParams) error {
	_, err := q.db.ExecContext(ctx, removeSongFromPlaylist, arg.PlaylistsID, arg.SongsID)
	return err
}

const updatePlaylist = `-- name: UpdatePlaylist :one
UPDATE playlists
SET name = $2,
    image = $3
WHERE id = $1
RETURNING id, users_id, name, image, created_at
`

type UpdatePlaylistParams struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (q *Queries) UpdatePlaylist(ctx context.Context, arg UpdatePlaylistParams) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, updatePlaylist, arg.ID, arg.Name, arg.Image)
	var i Playlist
	err := row.Scan(
		&i.ID,
		&i.UsersID,
		&i.Name,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}
