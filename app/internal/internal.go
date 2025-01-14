package internal

import (
	"strings"
)

func Ascii(text, banner string) (asciiArt string, status int) {
	if UserInputChecker(text) != "" {
		return "", 400
	}

	bannerMap, status := MapBuilder(banner)
	if status != 200 {
		return "", status
	}
	asciiArt = BuildAsciiArt(strings.Split(text, "\r\n"), bannerMap)
	return asciiArt, 200
}
