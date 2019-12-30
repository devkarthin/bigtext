package main

import (
	"errors"
	"strings"
)

// check if emoji exists in map "emoji"
func getEmoji(emojiName string) (string, error) {
	if value, ok := emojis[strings.ToLower(emojiName)]; ok {
		return value, nil
	}
	return emojiName, errors.New("Couldn't find emoji")
}

var emojis = map[string]string{
	"a": "🇦",
	"b": "🇧",
	"c": "🇨",
	"d": "🇩",
	"e": "🇪",
	"f": "🇫",
	"g": "🇬",
	"h": "🇭",
	"i": "🇮",
	"j": "🇯",
	"k": "🇰",
	"l": "🇱",
	"m": "🇲",
	"n": "🇳",
	"o": "🇴",
	"p": "🇵",
	"q": "🇶",
	"r": "🇷",
	"s": "🇸",
	"t": "🇹",
	"u": "🇺",
	"v": "🇻",
	"w": "🇼",
	"x": "🇽",
	"y": "🇾",
	"z": "🇿",
	"0": "0️⃣",
	"1": "1️⃣",
	"2": "2️⃣",
	"3": "3️⃣",
	"4": "4️⃣",
	"5": "5️⃣",
	"6": "6️⃣",
	"7": "7️⃣",
	"8": "8️⃣",
	"9": "9️⃣",
	"!": "❗️",
	"?": "❓",
	"#": "#️⃣"}
