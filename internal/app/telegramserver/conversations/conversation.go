package conversations

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// ContentType - type of content for conversation
type ContentType int

// Enum
const (
	Text ContentType = iota
	Photo
	Document
)

// Reply - contains text reply for customer and end flag
type Reply struct {
	Text  string
	IsEnd bool
}

// Conversation providing interface for "sending" data, and getting replies
type Conversation interface {
	SetText(text string) *Reply
	SetDocument(doc *tb.Document) *Reply
	SetPhoto(photo *tb.Photo) *Reply

	CurrentText() string
}

// ConversationPart providing interface for part of conversation
type ConversationPart struct {
	ReplyText string
	Text      string
	Want      ContentType
	Set       func(interface{}, interface{}) bool
}

// ConversationSequence sequence of conversation parts
type ConversationSequence []*ConversationPart
