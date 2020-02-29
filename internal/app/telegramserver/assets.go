package telegramserver

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	typesBtn = tb.InlineButton{
		Unique: "type_btn",
		Text:   "📚 Types of books",
	}

	aboutBtn = tb.InlineButton{
		Unique: "about_btn",
		Text:   "ℹ️ About",
	}

	sourceBtn = tb.InlineButton{
		Unique: "source_btn",
		Text:   "💾 Source Code",
	}

	helloMessage = "Hello, this is bot for sharing my collection of books, use this buttons for continue."

	menu = [][]tb.InlineButton{
		[]tb.InlineButton{typesBtn, aboutBtn, sourceBtn},
		// ...
	}
)
