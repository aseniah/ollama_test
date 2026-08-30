package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	// Increase buffer size for long lines if needed
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	// Regex to match sequences of letters
	letterRegex := regexp.MustCompile(`[a-zA-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		// Find all sequences of letters
		matches := letterRegex.FindAllString(line, -1)
		for _, match := range matches {
			word := strings.ToLower(match)
			wordCount[word]++
		}
	}

	// Collect words and their counts
	type wordCount struct {
		word  string
		count int
	}
	var words []wordCount
	for w, c := range wordCount {
		words = append(words, wordCount{w, c})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	for _, wc := range words {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}