package groupWord

import (
	"sort"
	"strings"
)

type Word struct {
	Key   string
	Value int
}

func GroupWord(inputWords string) string {
	
	words := []string{}
	words = append(words, strings.Split(inputWords, " ")...)

	wordsMap := map[string]Word{}
	var keys []string
	for _, word := range words {
		for _, splitedWord := range strings.Split(word, "") {
			if val, isKey := wordsMap[splitedWord]; isKey {
				wordsMap[splitedWord] = Word{splitedWord, val.Value + 1}
			} else {
				wordsMap[splitedWord] = Word{splitedWord, 1}
				keys = append(keys, splitedWord)
			}
		}
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if wordsMap[keys[i]].Value == wordsMap[keys[j]].Value {
			return wordsMap[keys[i]].Key < wordsMap[keys[j]].Key
		}
		return wordsMap[keys[i]].Value > wordsMap[keys[j]].Value
	})

	returnedStrings := `"`
	for _, key := range keys {
		returnedStrings += key
	}
	return returnedStrings + `"`
}
