package nlp

import (
	"regexp"
	"strings"

	"github.com/dhiller/go-teach/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

// Tokenize returns a slice of the string with normalized words
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if len(token) > 0 {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
