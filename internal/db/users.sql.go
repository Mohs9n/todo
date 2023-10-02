// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: users.sql

package db

import (
	"context"
)

const checkUserExists = `-- name: CheckUserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE name = $1)
`

func (q *Queries) CheckUserExists(ctx context.Context, name string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUserExists, name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name FROM users WHERE name = $1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}