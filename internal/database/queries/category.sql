-- name: CreateCategory :one
INSERT INTO categories(name, description)
VALUES (@name, @description)
RETURNING id;

-- name: GetCategory :one
SELECT name, description
FROM categories
WHERE id = @id;

-- name: UpdateCategory :exec
UPDATE categories
SET
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(slqc.narg('description'), description),
    updated_at = now()
WHERE id = @id;

-- name: DeleteCategory :exec
DELETE FROM books
WHERE id = @id;
