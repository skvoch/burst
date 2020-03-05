package conversations

import (
	"strconv"

	"github.com/skvoch/burst/internal/app/apiclient"
	"github.com/skvoch/burst/internal/app/model"
)

// CreateBookConversation conversation for creaing books
type CreateBookConversation struct {
	sequence ConversationSequence
	client   *apiclient.BurstClient

	book model.Book

	index int
}

// NewCreateBookConversation helper function
func NewCreateBookConversation(client *apiclient.BurstClient) *CreateBookConversation {

	sequence := ConversationSequence{
		&ConversationPart{
			Text: "Let's create a new type book, enter book name:",
			Want: Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				book, typeState := object.(*model.Book)
				text, textState := value.(string)

				if (typeState == false) || (textState == false) {
					return false
				}
				book.Name = text

				return true
			},
		},
		&ConversationPart{
			Text: "Please enter name of type:",
			Want: Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				book, typeState := object.(*model.Book)
				text, textState := value.(string)

				if (typeState == false) || (textState == false) {
					return false
				}
				bookType := t.findTypeByName(text)

				return true
			},
		},
		&ConversationPart{
			Text: "Please enter decsription of the book:",
			Want: Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				book, typeState := object.(*model.Book)
				text, textState := value.(string)

				if (typeState == false) || (textState == false) {
					return false
				}
				book.Description = text

				return true
			},
		},
		&ConversationPart{
			Text: "Please enter review of the book:",
			Want: Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				book, typeState := object.(*model.Book)
				text, textState := value.(string)

				if (typeState == false) || (textState == false) {
					return false
				}
				book.Review = text

				return true
			},
		},
		&ConversationPart{
			Text: "Please enter rating of the book (0-5)",
			Want: Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				book, typeState := object.(*model.Book)
				text, textState := value.(string)
				rating, _ := strconv.Atoi(text)
				book.Rating = rating

				if (typeState == false) || (textState == false) {
					return false
				}
				return true
			},
		},
		&ConversationPart{
			Text: "Please send preview file:",
			Want: Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				//book, typeState := object.(*model.Book)
				//photo, textState := value.(*tb.Photo)

				//if (typeState == false) || (textState == false) {
				//	return false
				//}
				// DO SOMETHING
				return true
			},
		},
		&ConversationPart{
			Text:      "Please send PDF file:",
			ReplyText: "The book has been created",
			Want:      Text,
			Set: func(object interface{}, value interface{}) bool {

				// This is looks ugly, I will refactor it later
				//book, typeState := object.(*model.Book)
				//doc, textState := value.(*tb.Document)

				//if (typeState == false) || (textState == false) {
				//	return false
				//}
				// DO SOMETHING

				book, err := t.client.Create
				return true
			},
		},
	}

	return &CreateBookConversation{
		sequence: sequence,
	}
}

func (c *CreateBookConversation) findTypeByName(name string) model.Type {
	return model.Type{}
}
