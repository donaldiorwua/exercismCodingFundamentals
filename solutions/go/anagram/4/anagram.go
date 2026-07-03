package anagram

import (
    "slices"
    "strings"
)

func normalize(s string) string {
    r := []rune(strings.ToLower(s))
    slices.Sort(r)
    return string(r)
}

func Detect(subject string, candidates []string) []string {
	result := []string{}
	subjectnorm := normalize(subject)
	for _, candidate := range candidates {
		if !strings.EqualFold(subject, candidate) && normalize(candidate) == subjectnorm {
            result = append(result, candidate)
        }
	}
	return result
}
