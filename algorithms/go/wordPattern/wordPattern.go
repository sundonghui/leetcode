package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	word2Parttern := make(map[string]byte)
	pattern2Word := make(map[byte]string)
	for i, word := range words {
		ch := pattern[i]
		_, wordExist := word2Parttern[word]
		_, patternExist := pattern2Word[ch]
		if wordExist && word2Parttern[word] != ch ||
			patternExist && pattern2Word[ch] != word {
			return false
		}
		word2Parttern[word] = ch
		pattern2Word[ch] = word
	}
	return true
}
