package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		for _, r := range line {
			if unicode.IsLetter(r) {
				var currentWord string
				for i := 0; i < len(line); i++ {
					if unicode.IsLetter(rune(line[i])) {
						currentWord += string(line[i])
					} else {
						break
					}
				}
			}
		}
	}

	reader = bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		inWord := false
		lineLen := len(line)

		for i := 0; i < lineLen; i++ {
			r := rune(line[i])
			if unicode.IsLetter(r) {
				currentWord += string(r)
				inWord = true
			} else if inWord {
				currentWord = strings.ToLower(currentWord)
				wordCounts[currentWord]++
				inWord = false
			}
		}
	}

	wordCountSlice := []string{}
	for k := range wordCounts {
		wordCountSlice = append(wordCountSlice, k)
	}

	sort.Slice(wordCountSlice, func(i, j int) bool {
		countI := wordCounts[wordCountSlice[i]]
		countJ := wordCounts[wordCountSlice[j]]
		if countI != countJ {
			return countI > countJ
		}
		return wordCountSlice[i] < wordCountSlice[j]
	})

	for _, word := range wordCountSlice {
		fmt.Printf("%s: %d\n", word, wordCounts[word])
	}
}