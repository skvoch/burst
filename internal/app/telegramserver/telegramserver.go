package telegramserver

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/telegramserver/conversations"
	tb "gopkg.in/tucnak/telebot.v2"
)

// TelegramServer ...
type TelegramServer struct {
	log    *logrus.Logger
	config *Config
	bot    *tb.Bot

	// Now supported only one conversation (for owner)
	conversation conversations.Conversation
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

	t.bot.Handle(&sourceBtn, func(m *tb.Message) {
		t.bot.Send(m.Sender, sourceCodeMessage)
	})

	t.bot.Handle(tb.OnPhoto, func(m *tb.Message) {
		if m.Sender.ID == t.config.OwnerID {

			if t.conversation != nil {
				reply := t.conversation.SetPhoto(m.Photo)
				t.bot.Send(m.Sender, reply.Text)

				if reply.IsEnd {
					t.conversation = nil
				}
			}
		}
	})

	t.bot.Handle(tb.OnDocument, func(m *tb.Message) {
		if m.Sender.ID == t.config.OwnerID {

			if t.conversation != nil {
				reply := t.conversation.SetDocument(m.Document)
				t.bot.Send(m.Sender, reply.Text)

				if reply.IsEnd {
					t.conversation = nil
				}
			}
		}
	})

	t.bot.Handle(tb.OnText, func(m *tb.Message) {

		if m.Sender.ID == t.config.OwnerID {

			if t.conversation != nil {
				reply := t.conversation.SetText(m.Text)
				t.bot.Send(m.Sender, reply.Text)

				if reply.IsEnd {
					t.conversation = nil
				}
			}
		}
	})

	t.bot.Handle(&createTypeButton, func(m *tb.Message) {
		if m.Sender.ID == t.config.OwnerID {
			t.conversation = conversations.NewCreateTypeConversation()

			t.bot.Send(m.Sender, t.conversation.CurrentText())
		}
	})
}
