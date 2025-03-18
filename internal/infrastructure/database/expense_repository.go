package database

import (
	"context"
	"database/sql"

	"github.com/felipeRese/financial-controller-bot/internal/core/db"
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/entity"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db}
}

func (r *ExpenseRepository) Save(expense *entity.Expense) error {
  dbExpense := db.CreateExpenseParams{
    ID: expense.ID.String(),
    UserID: expense.UserID,
    Amount: expense.Amount,
    Category: expense.Category,
  }
  queries := db.New(r.db)
  return queries.CreateExpense(context.Background(), dbExpense)
}
