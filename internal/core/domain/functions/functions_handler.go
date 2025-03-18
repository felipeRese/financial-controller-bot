package functions

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/felipeRese/financial-controller-bot/internal/core/domain"
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/entity"
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/usecase"
)

type FunctionHandler struct {
	ExpenseUseCase *usecase.ExpenseUseCase
}

var _ domain.FunctionHandler = (*FunctionHandler)(nil)

func NewFunctionHandler(expenseUC *usecase.ExpenseUseCase) *FunctionHandler {
	return &FunctionHandler{
		ExpenseUseCase: expenseUC,
	}
}

func (f *FunctionHandler) GetAvailableFunctions() []domain.Function {
	return []domain.Function{
		{
			Name:        "save_expense",
			Description: "Save an expense record in the financial controller",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"user_id": map[string]interface{}{
						"type":        "string",
						"description": "The ID of the user who made the expense",
					},
					"amount": map[string]interface{}{
						"type":        "number",
						"description": "The amount spent",
					},
					"category": map[string]interface{}{
						"type":        "string",
						"description": "The category of the expense",
            "enum": []string{
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
                    },
					  },
				},
				"required": []string{"user_id", "amount", "category"},
			},
		},
	}
}

func (f *FunctionHandler) ExecuteFunction(name string, args map[string]interface{}) (string, error) {
	switch name {
	case "save_expense":
		userID, userOk := args["user_id"].(string)
		amount, amountOk := args["amount"].(float64)
		category, categoryOk := args["category"].(string)

		if !userOk || !amountOk || !categoryOk {
			return "", errors.New("invalid parameters for saving expense")
		}

		expense := &entity.Expense{
			ID:       uuid.New(),
			UserID:   userID,
			Amount:   amount,
			Category: category,
		}

		err := f.ExpenseUseCase.SaveExpense(expense)
		if err != nil {
			return "", fmt.Errorf("failed to save expense: %v", err)
		}

		return fmt.Sprintf("Gasto de %.2fR$ em %s registrado com sucesso!", expense.Amount, expense.Category), nil

	default:
		return "", errors.New("unknown function: " + name)
	}
}
