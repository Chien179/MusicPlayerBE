// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: songs_genres.sql

package db

import (
	"context"
)

const createSongGenre = `-- name: CreateSongGenre :one
INSERT INTO songs_genres (
  songs_id,
  genres_id
) VALUES (
  $1, $2
)
RETURNING id, songs_id, genres_id
`

type CreateSongGenreParams struct {
	SongsID  int64 `json:"songs_id"`
	GenresID int64 `json:"genres_id"`
}

func (q *Queries) CreateSongGenre(ctx context.Context, arg CreateSongGenreParams) (SongsGenre, error) {
	row := q.db.QueryRowContext(ctx, createSongGenre, arg.SongsID, arg.GenresID)
	var i SongsGenre
	err := row.Scan(&i.ID, &i.SongsID, &i.GenresID)
	return i, err
}

const deleteSongGenre = `-- name: DeleteSongGenre :exec
DELETE FROM songs_genres
WHERE id = $1
`

func (q *Queries) DeleteSongGenre(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSongGenre, id)
	return err
}

const getSongGenre = `-- name: GetSongGenre :one
SELECT id, songs_id, genres_id FROM songs_genres
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSongGenre(ctx context.Context, id int64) (SongsGenre, error) {
	row := q.db.QueryRowContext(ctx, getSongGenre, id)
	var i SongsGenre
	err := row.Scan(&i.ID, &i.SongsID, &i.GenresID)
	return i, err
}

const listSongsGenres = `-- name: ListSongsGenres :many
SELECT id, songs_id, genres_id FROM songs_genres
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListSongsGenresParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSongsGenres(ctx context.Context, arg ListSongsGenresParams) ([]SongsGenre, error) {
	rows, err := q.db.QueryContext(ctx, listSongsGenres, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SongsGenre{}
	for rows.Next() {
		var i SongsGenre
		if err := rows.Scan(&i.ID, &i.SongsID, &i.GenresID); err != nil {
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
