package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	
	// Regex to match only characters (letters)
	reg := regexp.MustCompile(`[^a-z]+`)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.Fields(line)
		for _, w := range words {
			// Strip punctuation: keep only letters
			cleaned := reg.ReplaceAllString(w, " ")
			// Split again in case punctuation was inside a word (e.g., "word,word")
			parts := strings.Fields(cleaned)
			for _, p := range parts {
				counts[p]++
			}
		}
	}

	var result []wordCount
	for k, v := range counts {
		result = append(result, wordCount{k, v})
	}

	// Sort by count descending, then by word ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].count != result[j].count {
			return result[i].count > result[j].count
		}
		return result[i].word < result[j].word
	})

	for _, wc := range result {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}