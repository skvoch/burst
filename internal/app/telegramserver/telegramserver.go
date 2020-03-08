package telegramserver

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/apiclient"
	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/telegramserver/conversations"
	tb "gopkg.in/tucnak/telebot.v2"
)

// TelegramServer ...
type TelegramServer struct {
	log    *logrus.Logger
	config *Config
	bot    *tb.Bot
	client *apiclient.BurstClient

	handlers []func(m *tb.Message)

	// For each button for selecting types of books
	typesCache map[string]*model.Type

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
		config:     config,
		log:        logrus.New(),
		bot:        bot,
		client:     client,
		typesCache: make(map[string]*model.Type),
	}, nil
}

// Start ...
func (t *TelegramServer) Start() {
	t.bot.Start()
}

// SetupHandlers ...
func (t *TelegramServer) SetupHandlers() {
	t.bot.Handle("/start", t.handleStart)

	t.bot.Handle(tb.OnDocument, t.handleDocument)
	t.bot.Handle(tb.OnText, t.handleText)

	t.bot.Handle(&sourceBtn, t.handleSourceCodeButton)
	t.bot.Handle(&createTypeButton, t.handleCreateTypeButton)
	t.bot.Handle(&createBookButton, t.handleCreateBookButton)

	t.bot.Handle(&typesBtn, t.handleTypesButton)
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

func (t *TelegramServer) handleCreateBookButton(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {
		t.conversation = conversations.NewCreateBookConversation(t.client, t.log)

		t.bot.Send(m.Sender, t.conversation.CurrentText())
	}
}

func (t *TelegramServer) handleTypesButton(m *tb.Message) {
	types, err := t.client.GetAllTypes()

	if err != nil {
		t.log.Error(err)
	}

	replyKeys := make([][]tb.InlineButton, 0)
	keysRow := make([]tb.InlineButton, 0)

	// Add handling of "dynamic" buttons
	for index, _type := range types {
		typeBtn := tb.InlineButton{
			Text:   _type.Name,
			Unique: "_" + strconv.Itoa(index),
		}

		t.typesCache[typeBtn.Text] = _type

		t.bot.Handle(&typeBtn, func(c *tb.Callback) {
			btnName := typeBtn.Text
			typeID := t.typesCache[btnName].ID

			if err != nil {
				return
			}
			t.sendBookList(m.Sender, typeID)
			t.bot.Respond(c, &tb.CallbackResponse{})
		})

		keysRow = append(keysRow, typeBtn)
	}
	replyKeys = append(replyKeys, keysRow)

	_, err = t.bot.Send(m.Sender, "📖",
		&tb.ReplyMarkup{InlineKeyboard: replyKeys})

	if err != nil {
		t.log.Error(err)
	}
}

func (t *TelegramServer) sendBookList(user *tb.User, typeID int) {

	books, err := t.client.GetBookIDs(typeID)

	if err != nil {
		t.log.Error(err)
	}

	for _, id := range books {
		book, err := t.client.GetBookByID(id)

		if err != nil {
			t.log.Error("Cannot get book by ID:", err)
			continue
		}
		textMessage := ""

		textMessage += "Name: " + book.Name
		textMessage += "\n"
		textMessage += "Description: " + book.Description
		textMessage += "\n"
		textMessage += "Review: " + book.Review
		textMessage += "\n"
		textMessage += "Rating: " + renderRatingToEmoji(book.Rating)

		textMessage += "\n"

		t.bot.Send(user, textMessage)
	}
}

func (t *TelegramServer) handleDocument(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {
		if t.conversation != nil {

			if m.Document.OnDisk() == false {
				if err := t.bot.Download(&m.Document.File, m.Document.FileName); err != nil {
					t.log.Error("Cannot download file from Telegram:", err)
				}
			}

			reply := t.conversation.SetDocument(m.Document)

			if reply.Text != "" {
				t.bot.Send(m.Sender, reply.Text)
			}

			if reply.IsEnd {
				t.conversation = nil
				return
			}

			if t.conversation.CurrentText() != "" {
				t.bot.Send(m.Sender, t.conversation.CurrentText())
			}
		}
	}
}

func (t *TelegramServer) handleText(m *tb.Message) {
	if m.Sender.ID == t.config.OwnerID {

		if t.conversation != nil {

			reply := t.conversation.SetText(m.Text)

			if reply.Text != "" {
				t.bot.Send(m.Sender, reply.Text)
			}

			if reply.IsEnd {
				t.conversation = nil
				return
			}

			if t.conversation.CurrentText() != "" {
				t.bot.Send(m.Sender, t.conversation.CurrentText())
			}
		}
	}
}
