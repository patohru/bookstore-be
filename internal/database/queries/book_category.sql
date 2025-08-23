-- name: CreateBookCategory :exec
INSERT INTO book_category(book_id, category_id)
VALUES (@book_id, @category_id);

-- name: GetByCategory :many
SELECT book_id, category_id
FROM book_category
WHERE category_id = @category_id;

-- name: GetByBook :many
SELECT book_id, category_id
FROM book_category
WHERE book_id = @book_id;
