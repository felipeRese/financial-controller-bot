package repository

import "github.com/felipeRese/financial-controller-bot/internal/core/domain/entity"

type ExpenseRepositoryInterface interface {
	Save(expense *entity.Expense) error
}

type UserRepositoryInterface interface {
	Save(user *entity.User) error
}
