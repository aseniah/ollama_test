package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsLetter(r)
		})

		for _, word := range words {
			word = strings.ToLower(word)
			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	type wordFreq struct {
		word  string
		count int
	}

	var sortedWords []wordFreq
	for word, count := range wordCount {
		sortedWords = append(sortedWords, wordFreq{word, count})
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].count == sortedWords[j].count {
			return sortedWords[i].word < sortedWords[j].word
		}
		return sortedWords[i].count > sortedWords[j].count
	})

	for _, wf := range sortedWords {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}