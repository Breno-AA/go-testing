package words

import "strings"

func Count(s, w string) int {
	words := strings.Split(s, " ")
	if w != "" {
		return filter(words, w)
	}

	var total int
	for _, word := range words {
		if word == "" {
			continue
		}
		total++
	}

	return total
}

func filter(words []string, w string) (total int) {
	for _, word := range words {
		if word == w {
			total++
		}
	}

	return
}
