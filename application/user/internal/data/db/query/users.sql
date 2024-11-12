-- name: CreateUser :one
-- 不存在则创建, 存在则跳过并返回空记录
INSERT INTO users(email, password)
VALUES ($1, $2)
ON CONFLICT ($1) DO NOTHING
RETURNING *;

-- name: LoginUser :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUser :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;
