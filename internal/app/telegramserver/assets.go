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

	helloMessage      = "Hello, this is bot for sharing my collection of books, use this buttons for continue."
	sourceCodeMessage = "If you want to modify, or use this bot for your books collection, welcome to GitHub \nhttps://github.com/skvoch/burst"

	menu = [][]tb.ReplyButton{
		[]tb.ReplyButton{typesBtn, aboutBtn, sourceBtn},
	}
)
