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

	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		tokens := strings.Fields(line)

		for _, token := range tokens {
			cleaned := ""
			for _, ch := range token {
				if unicode.IsLetter(ch) {
					cleaned += string(ch)
				}
			}

			if cleaned != "" {
				wordCount[cleaned]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	type wordFreq struct {
		word  string
		count int
	}

	var results []wordFreq
	for w, c := range wordCount {
		results = append(results, wordFreq{w, c})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].count != results[j].count {
			return results[i].count > results[j].count
		}
		return results[i].word < results[j].word
	})

	for _, r := range results {
		fmt.Printf("%s: %d\n", r.word, r.count)
	}
}