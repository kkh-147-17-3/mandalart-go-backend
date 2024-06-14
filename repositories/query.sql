-- name: GetUserBySocialProviderInfo :one
SELECT * FROM users WHERE social_id = $1 AND social_provider = $2 LIMIT 1;


-- name: GetLatestSheetByOwnerId :one
SELECT * FROM sheets WHERE owner_id = $1 ORDER BY id DESC LIMIT 1;

-- name: GetMainCellsBySheetId :many
SELECT * FROM cells WHERE sheet_id = $1 AND step = 1;

-- name: GetTodosByCellId :many
SELECT * FROM todos WHERE cell_id = $1;

-- name: GetLatestSheetWithMainCellsByOwnerId :many
SELECT sheets.id, sheets.name, cells.id "cell_id", cells.color, cells.goal, cells.is_completed FROM sheets
JOIN cells ON sheets.id = cells.sheet_id AND cells.step = 2
WHERE sheets.id = (
    SELECT id FROM sheets WHERE sheets.owner_id = $1 ORDER BY id DESC LIMIT 1
);

-- name: CreateUser :one
INSERT INTO users(social_id, social_provider)
VALUES($1,$2)
RETURNING id;