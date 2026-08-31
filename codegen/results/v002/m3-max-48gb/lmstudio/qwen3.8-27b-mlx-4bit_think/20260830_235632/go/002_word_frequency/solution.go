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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text strings.Builder
	for scanner.Scan() {
		if text.Len() > 0 {
			text.WriteString(" ")
		}
		text.WriteString(scanner.Text())
	}

	words := strings.Fields(text.String())

	freq := make(map[string]int)
	for _, w := range words {
		var cleaned strings.Builder
		for _, r := range w {
			if unicode.IsLetter(r) {
				cleaned.WriteRune(unicode.ToLower(r))
			}
		}
		cs := cleaned.String()
		if cs != "" {
			freq[cs]++
		}
	}

	type wc struct {
		word  string
		count int
	}
	var pairs []wc
	for w, c := range freq {
		pairs = append(pairs, wc{w, c})
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