-- name: GetFizzBuzz :one
SELECT *
FROM fizzbuzz
WHERE id = ?
LIMIT 1;

-- name: ListFizzBuzzes :many
SELECT *
FROM fizzbuzz
ORDER BY id DESC;

-- name: CreateFizzBuzz :execresult
INSERT INTO fizzbuzz (useragent, message, created_at)
VALUES (?, ?, ?);