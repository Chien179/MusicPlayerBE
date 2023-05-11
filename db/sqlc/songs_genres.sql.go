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
WHERE genres_id = $1
AND songs_id = $2
`

type DeleteSongGenreParams struct {
	GenresID int64 `json:"genres_id"`
	SongsID  int64 `json:"songs_id"`
}

func (q *Queries) DeleteSongGenre(ctx context.Context, arg DeleteSongGenreParams) error {
	_, err := q.db.ExecContext(ctx, deleteSongGenre, arg.GenresID, arg.SongsID)
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
SELECT genres_id FROM songs_genres
WHERE songs_id = $1
ORDER BY id
`

func (q *Queries) ListSongsGenres(ctx context.Context, songsID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, listSongsGenres, songsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var genres_id int64
		if err := rows.Scan(&genres_id); err != nil {
			return nil, err
		}
		items = append(items, genres_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
