package telegramserver

func renderRatingToEmoji(rating int) string {
	value := rating

	if rating > 5 {
		value = 5
	}

	if rating < 0 {
		value = 0
	}

	emojiString := ""

	for i := 0; i < value; i++ {
		emojiString += "🌕"
	}

	for i := value; i < 5; i++ {
		emojiString += "🌑"
	}

	return emojiString
}
