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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}

	text := content.String()
	
	// Convert to lowercase
	text = strings.ToLower(text)
	
	// Remove punctuation - keep only letters and spaces
	re := regexp.MustCompile(`[^a-z\s]`)
	text = re.ReplaceAllString(text, " ")
	
	// Split into words
	words := strings.Fields(text)
	
	// Count frequencies
	freq := make(map[string]int)
	for _, w := range words {
		// Only keep words with at least one letter
		if w != "" {
			freq[w]++
		}
	}
	
	// Create slice of word-count pairs
	type wordCount struct {
		word  string
		count int
	}
	var pairs []wordCount
	for w, c := range freq {
		pairs = append(pairs, wordCount{word: w, count: c})
	}
	
	// Sort by count descending, then alphabetically ascending
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})
	
	// Output
	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}