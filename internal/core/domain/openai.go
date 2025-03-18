package domain

type OpenAIRequest struct {
    Messages  []Message   `json:"messages"`
    Model     string      `json:"model"`
    Functions []Function  `json:"functions"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content,omitempty"`
    Name    string `json:"name,omitempty"`
    FunctionCall *FunctionCall `json:"function_call,omitempty"`
}

type Function struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Parameters  map[string]interface{} `json:"parameters"`
}

type FunctionCall struct {
    Name string `json:"name"`
    Arguments map[string]interface{} `json:"arguments"`
}

type OpenAIResponse struct {
    Choices []struct {
        Message Message `json:"message"`
    } `json:"choices"`
}

type OpenAIClient interface {
    HandleUserInput(req OpenAIRequest) (*OpenAIResponse, error)
}
