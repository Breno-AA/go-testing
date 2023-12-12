package words

import "strings"

func Count(s, w string) int {
	arr := strings.Split(s, " ")

	words := make([]string, 0)

	for _, word := range arr {
		if word == "" {
			continue
		}
		words = append(words, word)
	}

	if w != "" {
		return filter(words, w)
	}

	return len(words)
}

func filter(words []string, w string) (total int) {
	for _, word := range words {
		if word == w {
			total++
		}
	}

	return
}
