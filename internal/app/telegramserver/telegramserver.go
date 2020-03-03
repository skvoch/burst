package telegramserver

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/apiclient"
	"github.com/skvoch/burst/internal/app/telegramserver/conversations"
	tb "gopkg.in/tucnak/telebot.v2"
)

// TelegramServer ...
type TelegramServer struct {
	log    *logrus.Logger
	config *Config
	bot    *tb.Bot
	client *apiclient.BurstClient

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

	client, err := apiclient.New(config.BurstServerAddr)

	if err != nil {
		return nil, err
	}

	return &TelegramServer{
		config: config,
		log:    logrus.New(),
		bot:    bot,
		client: client,
	}, nil
}

// Start ...
func (t *TelegramServer) Start() {
	t.bot.Start()
}

// SetupHandlers ...
func (t *TelegramServer) SetupHandlers() {
	t.bot.Handle("/start", t.handleStart)

	t.bot.Handle(tb.OnPhoto, t.handlePhoto)
	t.bot.Handle(tb.OnDocument, t.handleDocument)
	t.bot.Handle(tb.OnText, t.handleText)

	t.bot.Handle(&sourceBtn, t.handleSourceCodeButton)
	t.bot.Handle(&createTypeButton, t.handleCreateTypeButton)
}

func (t *TelegramServer) handleStart(m *tb.Message) {
	var keys [][]tb.ReplyButton

	if m.Sender.ID == t.config.OwnerID {
		keys = menuWithEdit
	} else {
		keys = menu
	}

	t.bot.Send(m.Sender, helloMessage, &tb.ReplyMarkup{
		ReplyKeyboard: keys,
	})
}

func (t *TelegramServer) handleSourceCodeButton(m *tb.Message) {
	t.bot.Send(m.Sender, sourceCodeMessage)
}

func (t *TelegramServer) handleCreateTypeButton(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {
		t.conversation = conversations.NewCreateTypeConversation(t.client)

		t.bot.Send(m.Sender, t.conversation.CurrentText())
	}
}

func (t *TelegramServer) handlePhoto(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {

		if t.conversation != nil {
			reply := t.conversation.SetPhoto(m.Photo)
			t.bot.Send(m.Sender, reply.Text)

			if reply.IsEnd {
				t.conversation = nil
			}
		}
	}
}

func (t *TelegramServer) handleDocument(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {

		if t.conversation != nil {
			reply := t.conversation.SetDocument(m.Document)
			t.bot.Send(m.Sender, reply.Text)

			if reply.IsEnd {
				t.conversation = nil
			}
		}
	}
}

func (t *TelegramServer) handleText(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {

		if t.conversation != nil {
			reply := t.conversation.SetText(m.Text)
			t.bot.Send(m.Sender, reply.Text)

			if reply.IsEnd {
				t.conversation = nil
			}
		}
	}
}
