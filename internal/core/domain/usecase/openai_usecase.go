package usecase

import (
	"errors"

	"github.com/felipeRese/financial-controller-bot/internal/core/domain"
)

type OpenAIUseCase struct {
    OpenAIRepo domain.OpenAIClient
    FunctionHandler domain.FunctionHandler
}

func NewOpenAIUseCase(client domain.OpenAIClient, handler domain.FunctionHandler) *OpenAIUseCase {
  return &OpenAIUseCase{OpenAIRepo: client, FunctionHandler: handler}
}

func (u *OpenAIUseCase) ProcessUserMessage(userMessage string, userId string) (string, error) {
    request := domain.OpenAIRequest{
        Messages: []domain.Message{
            {Role: "system", Content: "You can call functions when needed."},
            {Role: "user", Content: userMessage},
        },
        Model: "gpt-4-turbo",
        Functions: u.FunctionHandler.GetAvailableFunctions(),
    }

    response, err := u.OpenAIRepo.HandleUserInput(request)
    if err != nil {
        return "", err
    }

    if len(response.Choices) > 0 && response.Choices[0].Message.FunctionCall != nil {
        functionCall := response.Choices[0].Message.FunctionCall
        return u.FunctionHandler.ExecuteFunction(functionCall.Name, functionCall.Arguments)
    }

    return "", errors.New("no function call detected")
}
