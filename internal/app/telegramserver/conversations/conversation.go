package conversations

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// ContentType - type of contect for conversation
type ContentType int

// ...
const (
	Text ContentType = iota
	Photo
	Document
)

// Reply ...
type Reply struct {
	Text  string
	IsEnd bool
}

// Conversation ...
type Conversation interface {
	SetText(text string) *Reply
	SetDocument(doc *tb.Document) *Reply
	SetPhoto(photo *tb.Photo) *Reply

	CurrentText() string
}

// ConversationPart ...
type ConversationPart struct {
	ReplyText string
	Text      string
	Want      ContentType
	Set       func(interface{}, interface{}) bool
}

// ConversationSequence ...
type ConversationSequence []*ConversationPart
