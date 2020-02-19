package telegramserver

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type TelegramServer struct {
	log    *logrus.Logger
	config *Config
	bot    *tgbotapi.BotAPI
}

func New(config *Config) *TelegramServer {
	return &TelegramServer{
		config: config,
		log:    logrus.New(),
	}
}

func (t *TelegramServer) Start() {
	tgbot, err := tgbotapi.NewBotAPI(t.config.ApplicationToken)

	if err != nil {
		t.log.Fatal(err)
	}

	t.bot = tgbot
	t.log.Info("starting telegram server, token:", t.config.ApplicationToken)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		t.log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		t.bot.Send(msg)
	}
}
