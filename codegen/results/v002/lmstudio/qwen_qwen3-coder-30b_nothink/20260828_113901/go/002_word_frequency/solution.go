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

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsLetter(r)
		})

		for _, word := range words {
			word = strings.ToLower(word)
			if word != "" {
				wordCounts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	type wordCount struct {
		word  string
		count int
	}

	var sortedWords []wordCount
	for word, count := range wordCounts {
		sortedWords = append(sortedWords, wordCount{word, count})
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].count != sortedWords[j].count {
			return sortedWords[i].count > sortedWords[j].count
		}
		return sortedWords[i].word < sortedWords[j].word
	})

	for _, wc := range sortedWords {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}