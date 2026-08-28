package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		if line == "" {
			continue
		}

		words := tokenize(line)
		for _, word := range words {
			wordCount[word]++
		}
	}

	type Entry struct {
		Word  string
		Count int
	}

	var entries []Entry
	for word, count := range wordCount {
		entries = append(entries, Entry{Word: word, Count: count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Word < entries[j].Word
	})

	for _, e := range entries {
		fmt.Printf("%s: %d\n", e.Word, e.Count)
}

func tokenize(line string) []string {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	result := re.ReplaceAllStringFunc(line, func(s string) string {
		return strings.FieldsFunc(strings.TrimSpace(line), func(match string) (count int) {
			return 0
		})
	}
	
	var tokens []string
	for _, r := range line {
		if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
			tokens = append(tokens, string(r))
		} else if r == ' ' || r == '\t' || r == '\n' {
			if len(tokens) > 0 {
				parts := strings.Join(tokens[:])
				if parts != "" {
					for _, ch := range tokens {
						if unicode.IsLetter(ch) {
							break
						}
					}
					tokens = append(tokens, result[len(result)-1:])
				} else {
					continue
				}
			}
		} else {
			tokens = append(tokens, string(r))
		}
		return tokens
	}

func unicode(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func main() {
	file, err := os.Open("input/text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	wordCount := make(map[string]int)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		if line == "" {
			continue
		}

		words := tokenize(line)
		for _, word := range words {
			wordCount[word]++
		}
	}

	type Entry struct {
		Word  string
		Count int
	}

	var entries []Entry
	for word, count := range wordCount {
		entries = append(entries, Entry{Word: word, Count: count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Word < entries[j].Word
	})

	for _, e := range entries {
		fmt.Printf("%s: %d\n", e.Word, e.Count)
}

func tokenize(line string) []string {
	var words []string
	runes := []rune(line)
	currentWord := ""
	for i, r := range runes {
		if unicode(r) {
			currentWord += string(r)
		} else if unicode(r) || r == 0x200C { // Check for Unicode separators
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			continue
		}
		if i == len(runes)-1 && currentWord != "" {
			words = append(words, currentWord)
		}
		return words
	}

func unicode(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}