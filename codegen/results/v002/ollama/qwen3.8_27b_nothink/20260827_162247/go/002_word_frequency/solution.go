package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert to lowercase
		line = strings.ToLower(line)
		// Remove punctuation: keep only letters (a-z) and spaces
		var cleaned []string
		var currentWord []rune
		for _, ch := range line {
			if ch >= 'a' && ch <= 'z' {
				currentWord = append(currentWord, ch)
			} else {
				if len(currentWord) > 0 {
					cleaned = append(cleaned, string(currentWord))
					currentWord = nil
				}
			}
		}
		if len(currentWord) > 0 {
			cleaned = append(cleaned, string(currentWord))
		}
		for _, word := range cleaned {
			if word != "" {
				wordCount[word]++
			}
		}
	}

	type wordFreq struct {
		word  string
		count int
	}

	var words []wordFreq
	for w, c := range wordCount {
		words = append(words, wordFreq{word: w, count: c})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	for _, wf := range words {
		fmt.Printf("%s: %d\n", wf.word, wf.count)
	}
}