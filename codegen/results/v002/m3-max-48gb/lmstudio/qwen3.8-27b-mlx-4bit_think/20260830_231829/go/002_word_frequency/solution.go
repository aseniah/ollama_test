package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text strings.Builder
	for scanner.Scan() {
		text.WriteString(scanner.Text())
		text.WriteString(" ")
	}

	words := strings.Fields(text.String())
	counts := make(map[string]int)

	for _, w := range words {
		cleaned := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return unicode.ToLower(r)
			}
			return -1
		}, w)
		if cleaned != "" {
			counts[cleaned]++
		}
	}

	result := make([]wordCount, 0, len(counts))
	for w, c := range counts {
		result = append(result, wordCount{w, c})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].count != result[j].count {
			return result[i].count > result[j].count
		}
		return result[i].word < result[j].word
	})

	for _, rc := range result {
		fmt.Printf("%s: %d\n", rc.word, rc.count)
	}
}