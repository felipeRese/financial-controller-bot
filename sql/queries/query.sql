-- name: GetExpensesByUserId :many
SELECT * FROM expenses 
WHERE user_id = ?;

-- name:GetExpensesByCategoryAndUserId :many
SELECT * FROM expenses WHERE user_id = ? AND category = ?;

-- name: CreateExpense :exec
INSERT INTO expenses (id, user_id, amount, category)
VALUES (?, ?, ?, ?);
