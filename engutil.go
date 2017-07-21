package engutil

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Number the noun s n times
func Numbered(s string, n int) string {
	if n <= 0 {
		return "An impossible amount of " + Pluralise(s)
	} else if n == 1 {
		return ASlashAn(s)
	} else {
		return strconv.Itoa(n) + " " + Pluralise(s)
	}
}

// Put an "A" or "An" before the noun s
func ASlashAn(s string) string {
	if len(s) == 0 {
		return s
	}
	first, _ := utf8.DecodeRuneInString(s)
	if Vowel(first) {
		return "An " + s
	} else {
		return "A " + s
	}
}

// Is the given rune a vowel
func Vowel(ru rune) bool {
	switch unicode.ToLower(ru) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

// Pluralise the noun s
func Pluralise(s string) string {
	matchedCy, _ := regexp.MatchString("[^aeiou]y", s)
	if matchedCy {
		return s[:len(s)-1] + "ies"
	} else if strings.HasSuffix(s, "s") {
		return s + "es"
	}
	return s + "s"
}
