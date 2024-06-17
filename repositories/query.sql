-- name: GetUserBySocialProviderInfo :one
SELECT *
FROM users
WHERE social_id = $1
  AND social_provider = $2
LIMIT 1;


-- name: GetLatestSheetByOwnerId :one
SELECT *
FROM sheets
WHERE owner_id = $1
ORDER BY id DESC
LIMIT 1;

-- name: GetMainCellsBySheetId :many
SELECT *
FROM cells
WHERE sheet_id = $1
  AND step = 1;

-- name: GetTodosByCellId :many
SELECT *
FROM todos
WHERE cell_id = $1;

-- name: GetLatestSheetWithMainCellsByOwnerId :many
SELECT sheets.id, sheets.name, cells.id "cell_id", cells.color, cells.goal, cells.is_completed
FROM sheets
         JOIN cells ON sheets.id = cells.sheet_id AND cells.step = 2
WHERE sheets.id = (SELECT id
                   FROM sheets
                   WHERE sheets.owner_id = $1
                   ORDER BY id DESC
                   LIMIT 1);

-- name: CreateUser :one
INSERT INTO users(social_id, social_provider)
VALUES ($1, $2)
RETURNING id;

-- name: CreateSheet :one
INSERT INTO sheets(owner_id, name)
VALUES ($1, $2)
RETURNING id;

-- name: GetChildrenCellsByParentId :many
SELECT *
FROM cells
WHERE parent_id = $1
ORDER BY step;

-- name: GetCellById :one
SELECT *
FROM cells
WHERE id = $1;

-- name: CreateCell :one
INSERT INTO cells(owner_id, sheet_id, goal, color, step, "order", parent_id, is_completed)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
returning id;

-- name: UpdateCell :exec
UPDATE cells
SET goal         = $2,
    color        = $3,
    is_completed = $4
WHERE id = $1;

-- name: DeleteTodosByCellID :exec
DELETE
FROM todos
WHERE cell_id = $1;

-- name: InsertTodosByCellID :copyfrom
INSERT INTO todos (owner_id, cell_id, content)
VALUES ($1, $2, $3);

-- name: GetTodosByCellID :many
SELECT * FROM todos WHERE cell_id = $1;
