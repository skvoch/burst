package telegramserver

import (
	"time"

	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

// TelegramServer ...
type TelegramServer struct {
	log    *logrus.Logger
	config *Config
	bot    *tb.Bot

	conversations map[int]*Conversation
}

// New ...
func New(config *Config) (*TelegramServer, error) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  config.ApplicationToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	return &TelegramServer{
		config: config,
		log:    logrus.New(),
		bot:    bot,
	}, nil
}

// Start ...
func (t *TelegramServer) Start() {
	t.bot.Start()
}

// SetupHandlers ...
func (t *TelegramServer) SetupHandlers() {
	t.bot.Handle("/start", func(m *tb.Message) {
		var keys [][]tb.ReplyButton

		if m.Sender.ID == t.config.OwnerID {
			keys = menuWithEdit
		} else {
			keys = menu
		}

		t.bot.Send(m.Sender, helloMessage, &tb.ReplyMarkup{
			ReplyKeyboard: keys,
		})
	})

	t.bot.Handle(tb.OnText, func(m *tb.Message) {
		t.log.Info("text: ", m.Text)
	})

	t.bot.Handle(tb.OnPhoto, func(m *tb.Message) {

		t.log.Info("photo: ", m.Sender.ID)
	})

	t.bot.Handle(tb.OnDocument, func(m *tb.Message) {
		t.log.Info("document: ", m.Sender.ID)
	})

	t.bot.Handle(&sourceBtn, func(m *tb.Message) {
		t.log.Info(m.Sender.ID)

		t.bot.Send(m.Sender, sourceCodeMessage)
	})
}
