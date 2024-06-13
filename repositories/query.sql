-- name: GetUserBySocialProviderInfo :one
SELECT *
FROM users
WHERE social_id = $1
  AND social_provider = $2;


-- name: GetLatestSheetByOwnerId :one
SELECT * FROM sheets WHERE owner_id = $1 ORDER BY id DESC LIMIT 1;

-- name: GetMainCellsBySheetId :many
SELECT * FROM cells WHERE sheet_id = $1 AND step = 1;

-- name: GetTodosByCellId :many
SELECT * FROM todos WHERE cell_id = $1;