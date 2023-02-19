// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: playlists.sql

package db

import (
	"context"
)

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

const getPlaylist = `-- name: GetPlaylist :one
SELECT id, users_id, name, image, created_at FROM playlists
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPlaylist(ctx context.Context, id int64) (Playlist, error) {
	row := q.db.QueryRowContext(ctx, getPlaylist, id)
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

const listPlaylists = `-- name: ListPlaylists :many
SELECT id, users_id, name, image, created_at FROM playlists
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPlaylistsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPlaylists(ctx context.Context, arg ListPlaylistsParams) ([]Playlist, error) {
	rows, err := q.db.QueryContext(ctx, listPlaylists, arg.Limit, arg.Offset)
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
