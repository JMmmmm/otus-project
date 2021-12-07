package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(incomingText string) []string {
	var topWords []string
	if incomingText == "" {
		return topWords
	}

	wordAmountsMap := make(map[string]int)

	incomingText = strings.ReplaceAll(incomingText, "\n\t", " ")
	incomingText = strings.ReplaceAll(incomingText, "\t", "")
	incomingText = strings.ReplaceAll(incomingText, ",", "")
	incomingText = strings.ReplaceAll(incomingText, ".", "")
	incomingSlice := strings.Split(incomingText, " ")
	for _, incomingWord := range incomingSlice {
		if incomingWord == "-" || incomingWord == "" {
			continue
		}
		cleanedIncomingWord := strings.ToLower(incomingWord)
		amount := wordAmountsMap[cleanedIncomingWord]
		wordAmountsMap[cleanedIncomingWord] = amount + 1
	}

	type wordStruct struct {
		wordName string
		amount   int
	}
	sortedWordsStructs := make([]wordStruct, len(wordAmountsMap))
	i := 0
	for wordName, amount := range wordAmountsMap {
		sortedWordsStructs[i] = wordStruct{wordName, amount}
		i++
	}
	sort.Slice(sortedWordsStructs, func(i, j int) bool {
		if sortedWordsStructs[i].amount != sortedWordsStructs[j].amount {
			return sortedWordsStructs[i].amount > sortedWordsStructs[j].amount
		}

		return sortedWordsStructs[i].wordName < sortedWordsStructs[j].wordName
	})

	for i := 0; i < 10; i++ {
		topWords = append(topWords, sortedWordsStructs[i].wordName)
	}

	return topWords
}
