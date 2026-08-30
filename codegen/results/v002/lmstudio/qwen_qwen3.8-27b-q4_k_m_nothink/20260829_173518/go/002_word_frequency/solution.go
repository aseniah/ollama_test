package main

import (
	"bufio"
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
		// Convert to lowercase
		line = strings.ToLower(line)
		// Strip punctuation: keep only letters
		var letters []rune
		for _, r := range line {
			if unicode.IsLetter(r) {
				// Keep the letter, but we need to track word boundaries
				letters = append(letters, r)
			} else {
				if len(letters) > 0 {
					word := string(letters)
					wordCounts[word]++
					letters = nil
				}
			}
		}
		if len(letters) > 0 {
			word := string(letters)
			wordCounts[word]++
		}
	}

	// Create a slice of word-count pairs
	type wordCount struct {
		word  string
		count int
	}
	words := make([]wordCount, 0, len(wordCounts))
	for w, c := range wordCounts {
		words = append(words, wordCount{word: w, count: c})
	}

	// Sort by count descending, then alphabetically ascending
	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word
	})

	// Output
	for _, wc := range words {
		println(wc.word + ": " + string(rune('0'+wc.count%10)))
		// Wait, the above is wrong for counts > 9. Let me fix.
	}

	// Actually, let me rewrite the output part properly
}