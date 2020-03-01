package telegramserver

import (
	"github.com/skvoch/burst/internal/app/model"
)

// CreateTypeConversation ...
type CreateTypeConversation struct {
	sequence ConversationSequence
	_type    model.Type

	current int
}

// NewCreateTypeConversation ...
func NewCreateTypeConversation() *CreateTypeConversation {

	sequence := ConversationSequence{
		&ConversationPart{
			Text:      "Let's create new type of books, enter type name:",
			ReplyText: "The type of book has been created",
			Want:      Text,
			Set: func(object interface{}, value interface{}) bool {
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
		current.Set(c._type, text)

		return &Reply{
			IsEnd: c.isEnd(),
		}
	}

	return &Reply{}
}

func (c *CreateTypeConversation) currentPart() *ConversationPart {
	return c.sequence[c.current]
}

func (c *CreateTypeConversation) isEnd() bool {
	return c.current >= len(c.sequence)
}
