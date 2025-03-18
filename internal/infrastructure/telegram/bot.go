package telegram

import (
	"log"

	"github.com/felipeRese/financial-controller-bot/internal/core/domain/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
    Bot         *tgbotapi.BotAPI
    OpenAIUseCase *usecase.OpenAIUseCase
}


func NewTelegramBot(token string, aiUseCase *usecase.OpenAIUseCase) (*TelegramBot, error) {
    bot, err := tgbotapi.NewBotAPI(string(token))
    if err != nil {
        return nil, err
    }

    bot.Debug = true 

    return &TelegramBot{
        Bot:          bot,
        OpenAIUseCase: aiUseCase,
    }, nil
}

func (t *TelegramBot) Start() {
    log.Printf("Authorized on account %s", t.Bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    updates, err := t.Bot.GetUpdatesChan(updateConfig)
    if err != nil {
      log.Fatal(err)
    }

    for update := range updates {
        if update.Message != nil { // If we received a message
            go t.handleMessage(update.Message)
        }
    }
}

func (t *TelegramBot) handleMessage(message *tgbotapi.Message) {
    log.Printf("[%s] %s", message.From.UserName, message.Text)

    response, err := t.OpenAIUseCase.ProcessUserMessage(message.Text)
    if err != nil {
        response = err.Error()
    }

    msg := tgbotapi.NewMessage(message.Chat.ID, response)
    t.Bot.Send(msg)
}
