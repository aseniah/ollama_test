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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[^a-z]")

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.FieldsFunc(line, func(r rune) bool {
			return (r < 'a' || r > 'z') && r != '\''
		})

		for _, word := range words {
			cleaned := re.ReplaceAllString(word, "")
			if cleaned != "" {
				wordFreq[cleaned]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	type wordCount struct {
		word  string
		count int
	}

	var pairs []wordCount
	for word, count := range wordFreq {
		pairs = append(pairs, wordCount{word, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	for _, p := range pairs {
		fmt.Printf("%s: %d\n", p.word, p.count)
	}
}
