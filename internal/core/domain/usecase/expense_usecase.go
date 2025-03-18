package usecase

import (
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/entity"
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/repository"
)

type ExpenseUseCase struct {
  repository repository.ExpenseRepositoryInterface
}

func NewExpenseUseCase(repo repository.ExpenseRepositoryInterface) *ExpenseUseCase {
  return &ExpenseUseCase{repository: repo}
}

func (u *ExpenseUseCase) SaveExpense(expense *entity.Expense) error {
  return u.repository.Save(expense)
}
