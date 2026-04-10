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
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := extractWords(line)
		for _, word := range words {
			if word != "" {
				wordCounts[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	type wordCount struct {
		word  string
		count int
	}

	var counts []wordCount
	for word, count := range wordCounts {
		counts = append(counts, wordCount{word, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count != counts[j].count {
			return counts[i].count > counts[j].count
		}
		return counts[i].word < counts[j].word
	})

	for _, wc := range counts {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}

func extractWords(line string) []string {
	var words []string
	line = strings.ToLower(line)
	currentWord := strings.Builder{}

	for _, r := range line {
		if unicode.IsLetter(r) {
			currentWord.WriteRune(r)
		} else {
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		}
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}