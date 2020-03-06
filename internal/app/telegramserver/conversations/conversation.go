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

// SequenceHandler providing functional for handing conversation sequence
type SequenceHandler struct {
	sequence ConversationSequence

	index int
}

// Next - incrase index counter
func (c *SequenceHandler) Next() {
	c.index++
}

// CurrentPart return current part of sequence
func (c *SequenceHandler) CurrentPart() *ConversationPart {
	return c.sequence[c.index]
}

// IsEnd ...
func (c *SequenceHandler) IsEnd() bool {
	return c.index >= len(c.sequence)
}

// CurrentText providing text of current part of conversation
func (c *SequenceHandler) CurrentText() string {
	current := c.CurrentPart()

	return current.Text
}
