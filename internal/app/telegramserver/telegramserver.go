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
		t.bot.Send(m.Sender, helloMessage, &tb.ReplyMarkup{
			InlineKeyboard: menu,
		})
	})
	// Menu handling

	t.bot.Handle(&sourceBtn, func(c *tb.Callback) {
		t.bot.Respond(c, &tb.CallbackResponse{Text: "If you want to modify, or close this the bot \n"})
		t.bot.Send(c.Sender, "--")
	})
}
