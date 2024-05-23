package countWord

import (
	"fmt"
	"strings"
)

func CountWord(words string) string {

	wordMap := map[string]int{}
	var keys []string
	for _, word := range strings.Split(words, "") {
		if word != " " {
			w := strings.ToLower(word)
			if val, isKey := wordMap[w]; isKey {
				wordMap[w] = val + 1
			} else {
				wordMap[w] = 1
				keys = append(keys, w)
			}
		}
	}
	retunedStrings := `"`
	count := 0
	for _, key := range keys {
		count++
		if count == len(keys) {
			retunedStrings += fmt.Sprintf("%s=%d", key, wordMap[key])
		} else {
			retunedStrings += fmt.Sprintf("%s=%d, ", key, wordMap[key])
		}
	}
	return retunedStrings + `"`
}
