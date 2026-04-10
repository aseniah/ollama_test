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

	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		lowerline := strings.ToLower(line)

		// Remove non-letter characters
		pattern := regexp.MustCompile("[^a-z]+")
		words := pattern.FindAllString(lowerline, -1)

		for _, w := range words {
			if w != "" {
				wordFreq[w]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var words []string
	for word, count := range wordFreq {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if wordFreq[words[i]] != wordFreq[words[j]] {
			return wordFreq[words[i]] > wordFreq[words[j]]
		}
		return words[i] < words[j]
	})

	for _, w := range words {
		fmt.Printf("%s: %d\n", w, wordFreq[w])
	}
}