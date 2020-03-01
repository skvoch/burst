package telegramserver

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// Conversation ...
type Conversation interface {
	SetText(text string)
	SetDocument(doc *tb.Document)
	SetPhoto(img *tb.Photo)

	Reply() string
}
