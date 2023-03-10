// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: genres.sql

package db

import (
	"context"
)

const createGenre = `-- name: CreateGenre :one
INSERT INTO genres (
  name,
  image
) VALUES (
  $1, $2
)
RETURNING id, name, image, created_at
`

type CreateGenreParams struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (q *Queries) CreateGenre(ctx context.Context, arg CreateGenreParams) (Genre, error) {
	row := q.db.QueryRowContext(ctx, createGenre, arg.Name, arg.Image)
	var i Genre
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}

const deleteGenre = `-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = $1
`

func (q *Queries) DeleteGenre(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGenre, id)
	return err
}

const getGenre = `-- name: GetGenre :one
SELECT id, name, image, created_at FROM genres
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGenre(ctx context.Context, id int64) (Genre, error) {
	row := q.db.QueryRowContext(ctx, getGenre, id)
	var i Genre
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}

const listGenres = `-- name: ListGenres :many
SELECT id, name, image, created_at FROM genres
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListGenresParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListGenres(ctx context.Context, arg ListGenresParams) ([]Genre, error) {
	rows, err := q.db.QueryContext(ctx, listGenres, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Genre{}
	for rows.Next() {
		var i Genre
		if err := rows.Scan(
			&i.ID,
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

const updategenre = `-- name: Updategenre :one
UPDATE genres
SET name = $2,
    image = $3
WHERE id = $1
RETURNING id, name, image, created_at
`

type UpdategenreParams struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (q *Queries) Updategenre(ctx context.Context, arg UpdategenreParams) (Genre, error) {
	row := q.db.QueryRowContext(ctx, updategenre, arg.ID, arg.Name, arg.Image)
	var i Genre
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.CreatedAt,
	)
	return i, err
}
