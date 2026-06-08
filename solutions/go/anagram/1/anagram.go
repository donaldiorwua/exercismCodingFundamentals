package anagram

import (
    "slices"
    "strings"
)

func Detect(subject string, candidates []string) []string {
	result := []string{}
	subjec := strings.ToLower(subject)
	sub := []rune(subjec)
	slices.Sort(sub)
	subnorm := string(sub)

	for _, words := range candidates {
		word := strings.ToLower(words)
		char := []rune(word)
		slices.Sort(char)
		charnorm := string(char)

		if word == subjec {
			continue
		}
		if charnorm == subnorm {
			result = append(result, words)
		}
	}
	return result
}