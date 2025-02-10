package entity

type ExpenseRepositoryInterface interface {
  Save(expense *Expense) error
}

type UserRepositoryInterface interface {
  Save(user *User) error
}
