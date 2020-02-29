package telegramserver

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	typesBtn = tb.ReplyButton{
		Text: "📚 Types of books",
	}

	aboutBtn = tb.ReplyButton{
		Text: "ℹ️ About",
	}

	sourceBtn = tb.ReplyButton{
		Text: "💾 Source Code",
	}

	editButton = tb.ReplyButton{
		Text: "⚙️ Edit content",
	}

	helloMessage      = "Hello, this is bot for sharing my collection of books, use buttons for continue."
	sourceCodeMessage = "https://github.com/skvoch/burst"

	menu = [][]tb.ReplyButton{
		[]tb.ReplyButton{typesBtn, aboutBtn, sourceBtn},
	}

	menuWithEdit = [][]tb.ReplyButton{
		[]tb.ReplyButton{typesBtn, aboutBtn, sourceBtn},
		[]tb.ReplyButton{editButton},
	}
)
