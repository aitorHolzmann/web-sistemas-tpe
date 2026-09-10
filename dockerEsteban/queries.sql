-- name: CreateTask :one
INSERT INTO tasks (title, description)
VALUES ($1, $2)
RETURNING id, title, description, completed, created_at;

-- name: GetTask :one
SELECT id, title, description, completed, created_at
FROM tasks
WHERE id = $1;

-- name: ListTasks :many
SELECT id, title, description, completed, created_at
FROM tasks
ORDER BY id;

-- name: UpdateTask :exec
UPDATE tasks
SET title = $2, description = $3, completed = $4
WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
