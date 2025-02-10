package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Expense struct {
	ID       uuid.UUID
	UserID   string
	Amount   float64
	Category string
}

func NewExpense(amount float64, category, userID string) (*Expense, error) {
	expense := &Expense{
		ID:       uuid.New(),
		Amount:   amount,
		Category: category,
		UserID:   userID,
	}

	err := expense.IsValid()
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (e *Expense) IsValid() error {

	_, err := uuid.Parse(e.ID.String())
	if err != nil {
		return errors.New("invalid UUID format")
	}

	if e.Amount <= 0 {
		return errors.New("invalid amount")
	}

	categories := []string{
		"Alimentação",
		"Transporte",
		"Moradia",
		"Educação",
		"Saúde",
		"Lazer",
		"Vestuário",
		"Serviços",
		"Assinaturas",
		"Impostos",
		"Outros",
	}

	validCategory := false

	for _, category := range categories {
		if e.Category == category {
			validCategory = true
		}
	}
	if !validCategory {
		return errors.New("invalid category")
	}

	return nil
}
