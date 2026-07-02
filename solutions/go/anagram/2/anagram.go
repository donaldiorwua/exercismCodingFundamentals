package anagram

import (
    "slices"
    "strings"
)
func normalize(s string) string {
    s = strings.ToLower(s)
    r := []rune(s)
    slices.Sort(r)

    return string(r)
}
func Detect(subject string, candidates []string) []string {
	result := []string{}
	subjectlow := strings.ToLower(subject)
	subjectnorm := normalize(subjectlow)

	for _, candidate := range candidates {
		if !strings.EqualFold(subject, candidate) && normalize(candidate) == subjectnorm {
            result = append(result, candidate)
        }
	}
	return result
}