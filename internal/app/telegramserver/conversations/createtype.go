package conversations

import (
	"encoding/json"

	"github.com/skvoch/burst/internal/app/model"
	tb "gopkg.in/tucnak/telebot.v2"
)

// CreateTypeConversation - conversation for creating types of book, will ask custumer
// about the Type name
type CreateTypeConversation struct {
	sequence ConversationSequence
	_type    model.Type

	index int
}

// NewCreateTypeConversation helper function
func NewCreateTypeConversation() *CreateTypeConversation {

	sequence := ConversationSequence{
		&ConversationPart{
			Text:      "Let's create new type of books, enter type name:",
			ReplyText: "The type of book has been created",
			Want:      Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				_type, typeState := object.(*model.Type)
				text, textState := value.(string)

				if (typeState == false) || (textState == false) {
					return false
				}
				_type.Name = text

				return true
			},
		},
	}

	return &CreateTypeConversation{
		sequence: sequence,
	}
}

// SetText ...
func (c *CreateTypeConversation) SetText(text string) *Reply {
	current := c.currentPart()

	if current.Want == Text {
		current.Set(&c._type, text)
		bytes, _ := json.Marshal(c._type)
		text = string(bytes)

		c.index++

		return &Reply{
			IsEnd: c.isEnd(),
			Text:  current.ReplyText,
		}
	}

	return &Reply{}
}

// SetDocument unused in this conversation
func (c *CreateTypeConversation) SetDocument(text *tb.Document) *Reply {

	return &Reply{}
}

// SetPhoto unused in this conversation
func (c *CreateTypeConversation) SetPhoto(photo *tb.Photo) *Reply {

	return &Reply{}
}

// CurrentText providing text of current part of conversation
func (c *CreateTypeConversation) CurrentText() string {
	current := c.currentPart()

	return current.Text
}

func (c *CreateTypeConversation) currentPart() *ConversationPart {
	return c.sequence[c.index]
}

func (c *CreateTypeConversation) isEnd() bool {
	return c.index >= len(c.sequence)
}
