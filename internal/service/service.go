package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Converter(text string) string {
	checker := func(r rune) bool {
		return unicode.IsLetter(r)
	}

	if strings.ContainsFunc(text, checker) {
		return morse.ToMorse(text)
	}

	return morse.ToText(text)
}