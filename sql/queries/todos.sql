-- name: GetTodos :many
SELECT todos.id, todos.text, todos.done
FROM todos
JOIN users ON todos.user_id = users.id
WHERE users.id = $1;

-- name: CreateTodo :one
INSERT INTO todos
(text, done, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET done = NOT done
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;