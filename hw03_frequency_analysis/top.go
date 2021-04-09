package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const WordCount = 10

var (
	textSplitterRegExp = regexp.MustCompile(`[\s\n]+`)
	workFinderRegExp   = regexp.MustCompile(`[\p{L}\w-]+`)
)

type WordCounter struct {
	Word  string
	Count int
}

func NewWorkCount(word string) *WordCounter {
	return &WordCounter{Word: word, Count: 1}
}

func Top10(text string) []string {
	wordMap := make(map[string]*WordCounter)
	words := textSplitterRegExp.Split(text, -1)

	for _, word := range words {
		matchWord := workFinderRegExp.FindString(strings.ToLower(word))
		if matchWord != "" && matchWord != "-" {
			if _, ok := wordMap[matchWord]; !ok {
				wordMap[matchWord] = NewWorkCount(matchWord)
			} else {
				wordMap[matchWord].Count++
			}
		}
	}

	listOfWordCounters := make([]WordCounter, 0, len(wordMap))
	for _, str := range wordMap {
		listOfWordCounters = append(listOfWordCounters, *str)
	}

	sort.SliceStable(listOfWordCounters, func(i, j int) bool {
		prev, next := listOfWordCounters[i], listOfWordCounters[j]
		if prev.Count == next.Count {
			return prev.Word < next.Word
		}

		return next.Count < prev.Count
	})

	answerSlice := make([]string, 0, WordCount)
	for i, val := range listOfWordCounters {
		if i > WordCount-1 {
			break
		}
		answerSlice = append(answerSlice, val.Word)
	}

	return answerSlice
}
