-- name: GetMovie :one
SELECT id, name, description, release_date, runtime, genre, director
FROM movies
WHERE id = $1;
-- name: ListMovies :many
SELECT id, name, description, release_date, runtime, genre, director
FROM movies
ORDER BY name;
-- name: CreateMovie :one
INSERT INTO movies (name, description, release_date, runtime, genre, director)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, description, release_date, runtime, genre, director;
-- name: UpdateMovie :exec
UPDATE movies
SET name = $2, description = $3, release_date = $4, runtime = $5, genre = $6, director = $7
WHERE id = $1;
-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1;
