-- name: CreateBook :one
INSERT INTO books(title, description, isbn, author)
VALUES (@title, @description, @isbn, @author)
RETURNING id;

-- name: GetBook :one
SELECT title, description, isbn, author
FROM books
WHERE id = @id;

-- name: UpdateBook :exec
UPDATE books
SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    author = COALESCE(sqlc.narg('author'), author),
    updated_at = now()
WHERE id = @id;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = @id;
