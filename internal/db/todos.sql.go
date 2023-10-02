// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: todos.sql

package db

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos
(text, done, user_id)
VALUES ($1, $2, $3)
RETURNING id, text, done, user_id
`

type CreateTodoParams struct {
	Text   string
	Done   bool
	UserID int32
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.Text, arg.Done, arg.UserID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Done,
		&i.UserID,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodos = `-- name: GetTodos :many
SELECT todos.id, todos.text, todos.done
FROM todos
JOIN users ON todos.user_id = users.id
WHERE users.id = $1
`

type GetTodosRow struct {
	ID   int32
	Text string
	Done bool
}

func (q *Queries) GetTodos(ctx context.Context, id int32) ([]GetTodosRow, error) {
	rows, err := q.db.QueryContext(ctx, getTodos, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTodosRow
	for rows.Next() {
		var i GetTodosRow
		if err := rows.Scan(&i.ID, &i.Text, &i.Done); err != nil {
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

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET done = NOT done
WHERE id = $1
RETURNING id, text, done, user_id
`

func (q *Queries) UpdateTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Done,
		&i.UserID,
	)
	return i, err
}