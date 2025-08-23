-- name: CreateReview :one
INSERT INTO reviews(rate, book_id, user_id, content)
VALUES (@rate, @book_id, @user_id, @content)
RETURNING id;

-- name: GetReview :one
SELECT rate, book_id, user_id, content
FROM reviews
WHERE id = @id;

-- name: GetReviewByBook :many
SELECT rate, book_id, user_id, content
FROM reviews
WHERE book_id = @book_id;

-- name: UpdateReview :exec
UPDATE reviews
SET
    rate = COALESCE(sqlc.narg('rate'), rate),
    content = COALESCE(sqlc.narg('content'), content),
    is_edited = true,
    updated_at = now()
WHERE id = @id;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE id = @id AND user_id = @user_id;
