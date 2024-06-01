-- name: CreateUser :one
INSERT INTO users (email, password_hash, avatar_url)
VALUES ($1, $2, $3)
RETURNING id, email, created_at;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, google_id, avatar_url, created_at
FROM users
WHERE email = $1;

-- name: CreateUserWithGoogle :one
INSERT INTO users (email, google_id, avatar_url)
VALUES ($1, $2)
RETURNING id, email, created_at;

-- name: CreateUserWithApple :one
INSERT INTO users (email, apple_id, avatar_url)
VALUES ($1, $2, $3)
RETURNING id, email, created_at;

-- name: CreateUserWithMicrosoft :one
INSERT INTO users (email, microsoft_id, avatar_url)
VALUES ($1, $2, $3)
RETURNING id, email, created_at;

-- name: GetUserByAppleID :one
SELECT id, email, password_hash, apple_id, google_id, microsoft_id, avatar_url, created_at
FROM users
WHERE apple_id = $1;

-- name: GetUserByMicrosoftID :one
SELECT id, email, password_hash, apple_id, google_id, microsoft_id, avatar_url, created_at
FROM users
WHERE microsoft_id = $1;

-- name: GetUserByGoogleID :one
SELECT id, email, password_hash, google_id, created_at
FROM users
WHERE google_id = $1;
