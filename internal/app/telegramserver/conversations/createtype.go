package conversations

import (
	"github.com/skvoch/burst/internal/app/apiclient"
	"github.com/skvoch/burst/internal/app/model"
	tb "gopkg.in/tucnak/telebot.v2"
)

// CreateTypeConversation - conversation for creating types of book, will ask custumer
// about the Type name
type CreateTypeConversation struct {
	handler SequenceHandler

	client *apiclient.BurstClient
	_type  model.Type
}

// NewCreateTypeConversation helper function
func NewCreateTypeConversation(client *apiclient.BurstClient) *CreateTypeConversation {

	sequence := ConversationSequence{
		&ConversationPart{
			Text:      "Let's create a new type of books, enter type name:",
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
		handler: SequenceHandler{
			sequence: sequence,
		},
		client: client,
	}
}

// SetText ...
func (c *CreateTypeConversation) SetText(text string) *Reply {
	current := c.handler.CurrentPart()

	if current.Want == Text {
		current.Set(&c._type, text)
		c.handler.Next()

		isEnd := c.handler.IsEnd()

		if isEnd {
			c.client.CreateType(&c._type)
		}

		return &Reply{
			IsEnd: c.handler.IsEnd(),
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
	current := c.handler.CurrentPart()

	return current.Text
}
