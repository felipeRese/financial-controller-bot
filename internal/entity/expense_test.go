package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExpense_IsValid_InvalidAmount(t *testing.T) {
	// Given an expense with a non-positive amount and a valid category
	expense := Expense{
		ID:       uuid.New(),
		Amount:   0, // Invalid amount
		Category: "Alimentação", // Valid category
	}

	// When calling IsValid
	err := expense.IsValid()

	// Then an error should be returned
	assert.Error(t, err)
	assert.Equal(t, "invalid amount", err.Error())
}

func TestExpense_IsValid_InvalidCategory(t *testing.T) {
	// Given an expense with a valid amount but an invalid category
	expense := Expense{
		ID:       uuid.New(),
		Amount:   100.0,
		Category: "InvalidCategory", // Not in the valid category list
	}

	// When calling IsValid
	err := expense.IsValid()

	// Then an error should be returned
	assert.Error(t, err)
	assert.Equal(t, "invalid category", err.Error())
}

func TestNewExpense_WithInvalidAmount(t *testing.T) {
	// When creating a new expense with an invalid amount
	expense, err := NewExpense(0, "Alimentação")

	// Then an error should be returned and no expense created
	assert.Nil(t, expense)
	assert.Error(t, err)
	assert.Equal(t, "invalid amount", err.Error())
}

func TestNewExpense_WithInvalidCategory(t *testing.T) {
	// When creating a new expense with an invalid category
	expense, err := NewExpense(100.0, "InvalidCategory")

	// Then an error should be returned and no expense created
	assert.Nil(t, expense)
	assert.Error(t, err)
	assert.Equal(t, "invalid category", err.Error())
}

func TestNewExpense_WithValidParams(t *testing.T) {
	// When creating a new expense with valid parameters
	expense, err := NewExpense(100.0, "Alimentação")

	// Then no error should be returned and the expense fields should be correctly set
	assert.NotNil(t, expense)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, expense.Amount)
	assert.Equal(t, "Alimentação", expense.Category)
	// Verify that a valid UUID was generated
	assert.NotEqual(t, uuid.Nil, expense.ID)
}
