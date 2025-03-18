package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/felipeRese/financial-controller-bot/configs"
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/functions"
	"github.com/felipeRese/financial-controller-bot/internal/core/domain/usecase"
	"github.com/felipeRese/financial-controller-bot/internal/infrastructure/database"
	"github.com/felipeRese/financial-controller-bot/internal/infrastructure/openai"
	"github.com/felipeRese/financial-controller-bot/internal/infrastructure/telegram"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  cfg, err:= configs.LoadConfig(".")
  if err != nil {
    log.Fatal(err)
  }
  
	apiKey := cfg.OpenAiKey
	telegramToken := cfg.TelegramBotToken

	if apiKey == "" || telegramToken == "" {
		log.Fatal("Missing API keys: Set OPENAI_API_KEY and TELEGRAM_BOT_TOKEN")
	}

  data, err := sql.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,))
  if err != nil {
    log.Fatal(err)
  }

  expenseRepository := database.NewExpenseRepository(data)
  openaiClient := openai.NewOpenAIClient(apiKey)
  expenseUseCase := usecase.NewExpenseUseCase(expenseRepository)
  functionHandler := functions.NewFunctionHandler(expenseUseCase)
  openaiUseCase := usecase.NewOpenAIUseCase(openaiClient, functionHandler)
  telegramBot, err := telegram.NewTelegramBot(telegramToken, openaiUseCase)
  if err != nil {
    log.Fatal(err)
  }

  go telegramBot.Start()


	log.Println("Server running...")
	select {} // Keep the main function running
}
