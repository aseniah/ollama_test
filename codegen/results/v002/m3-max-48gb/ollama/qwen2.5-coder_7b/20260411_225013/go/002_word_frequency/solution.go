package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		for _, word := range words(line) {
			wordCounts[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort by count descending, then alphabetically ascending for ties
	type WordCount struct {
		word  string
		count int
	}
	var wordCountList []WordCount
	for word, count := range wordCounts {
		wordCountList = append(wordCountList, WordCount{word, count})
	}

	sort.SliceStable(wordCountList, func(i, j int) bool {
		if wordCountList[i].count == wordCountList[j].count {
			return wordCountList[i].word < wordCountList[j].word
		}
		return wordCountList[i].count > wordCountList[j].count
	})

	// Output the results
	for _, wc := range wordCountList {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}

func words(s string) []string {
	var words []string
	word := ""
	inWord := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			word += strings.ToLower(string(r))
			inWord = true
		} else if inWord {
			words = append(words, word)
			word = ""
			inWord = false
		}
	}

	if inWord {
		words = append(words, word)
	}

	return words
}