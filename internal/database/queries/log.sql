-- name: CreateLog :exec
INSERT INTO "logs" ("data") VALUES ($1);
