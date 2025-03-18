package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/felipeRese/financial-controller-bot/internal/core/domain"
)


type OpenAIClient struct {
    APIKey string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
    return &OpenAIClient{APIKey: apiKey}
}

func (c *OpenAIClient) HandleUserInput(req domain.OpenAIRequest) (*domain.OpenAIResponse, error) {
    url := "https://api.openai.com/v1/chat/completions"

    requestBody, err := json.Marshal(req)
    if err != nil {
        return nil, err
    }

    reqBody := bytes.NewBuffer(requestBody)
    httpReq, err := http.NewRequest("POST", url, reqBody)
    if err != nil {
        return nil, err
    }

    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Authorization", "Bearer "+string(c.APIKey))

    client := &http.Client{}
    resp, err := client.Do(httpReq)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("failed to process user input: " + resp.Status)
    }

    var openAIResp domain.OpenAIResponse
    if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
        return nil, err
    }
    
  fmt.Print(openAIResp)

    return &openAIResp, nil
}
