package conversations

import (
	"strconv"

	"github.com/skvoch/burst/internal/app/apiclient"
	"github.com/skvoch/burst/internal/app/model"
	tb "gopkg.in/tucnak/telebot.v2"
)

// CreateBookConversation conversation for creaing books
type CreateBookConversation struct {
	SequenceHandler

	client *apiclient.BurstClient

	book    model.Book
	preview *tb.Photo
	file    *tb.Document
}

// NewCreateBookConversation helper function
func NewCreateBookConversation(client *apiclient.BurstClient) *CreateBookConversation {
	conversation := &CreateBookConversation{}

	sequence := ConversationSequence{
		&ConversationPart{
			Text: "Let's create a new type book, enter book name:",
			Want: Text,
			Set:  conversation.handleBookName,
		},
		&ConversationPart{
			Text: "Please enter name of type:",
			Want: Text,
			Set:  conversation.handleBookType,
		},
		&ConversationPart{
			Text: "Please enter decsription of the book:",
			Want: Text,
			Set:  conversation.handleBookDescription,
		},
		&ConversationPart{
			Text: "Please enter review of the book:",
			Want: Text,
			Set:  conversation.handleBookReview,
		},
		&ConversationPart{
			Text: "Please enter rating of the book (0-5)",
			Want: Text,
			Set:  conversation.handleBookRating,
		},
		&ConversationPart{
			Text: "Please send preview file:",
			Want: Text,
			Set:  conversation.handlePreview,
		},
		&ConversationPart{
			Text:      "Please send PDF file:",
			ReplyText: "The book has been created",
			Want:      Text,
			Set:       conversation.handleFile,
		},
	}

	conversation.setSequence(sequence)

	return conversation
}

func (c *CreateBookConversation) findTypeByName(name string) model.Type {
	types, _ := c.client.GetAllTypes()

	for _, current := range types {
		if current.Name == name {
			return *current
		}
	}

	return model.Type{}
}

func (c *CreateBookConversation) handleBookName(object interface{}, value interface{}) bool {

	book, typeState := object.(*model.Book)
	text, textState := value.(string)

	if (typeState == false) || (textState == false) {
		return false
	}
	book.Name = text

	return true
}

func (c *CreateBookConversation) handleBookType(object interface{}, value interface{}) bool {

	book, typeState := object.(*model.Book)
	text, textState := value.(string)

	if (typeState == false) || (textState == false) {
		return false
	}
	bookType := c.findTypeByName(text)
	book.Type = bookType.ID

	return true
}

func (c *CreateBookConversation) handleBookDescription(object interface{}, value interface{}) bool {

	// This is looks ugly, I will refactor it later
	book, typeState := object.(*model.Book)
	text, textState := value.(string)

	if (typeState == false) || (textState == false) {
		return false
	}
	book.Description = text

	return true
}

func (c *CreateBookConversation) handleBookReview(object interface{}, value interface{}) bool {

	// This is looks ugly, I will refactor it later
	book, typeState := object.(*model.Book)
	text, textState := value.(string)

	if (typeState == false) || (textState == false) {
		return false
	}
	book.Review = text

	return true
}

func (c *CreateBookConversation) handleBookRating(object interface{}, value interface{}) bool {

	// This is looks ugly, I will refactor it later
	book, typeState := object.(*model.Book)
	text, textState := value.(string)
	rating, _ := strconv.Atoi(text)
	book.Rating = rating

	if (typeState == false) || (textState == false) {
		return false
	}
	return true
}

func (c *CreateBookConversation) handlePreview(object interface{}, value interface{}) bool {

	photo := value.(*tb.Photo)
	c.preview = photo

	return true
}

func (c *CreateBookConversation) handleFile(object interface{}, value interface{}) bool {
	document := value.(*tb.Document)
	c.file = document

	return true
}

func (c *CreateBookConversation) SetText(text string) *Reply {
	current := c.CurrentPart()

	if current.Want == Text {
		current.Set(&c.book, text)
		c.Next()

		//isEnd := c.IsEnd()

		//if isEnd {
		//	c.client.CreateType(&c._type)
		//}

		return &Reply{
			IsEnd: c.IsEnd(),
			Text:  current.ReplyText,
		}
	}

	return &Reply{}
}

// SetDocument unused in this conversation
func (c *CreateBookConversation) SetDocument(text *tb.Document) *Reply {

	return &Reply{}
}

// SetPhoto unused in this conversation
func (c *CreateBookConversation) SetPhoto(photo *tb.Photo) *Reply {

	return &Reply{}
}

func (c *CreateBookConversation) uploadBookData() {

}
