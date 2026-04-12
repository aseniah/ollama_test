package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputPath := "input/text.txt"
	
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	wordCount := make(map[string]int)
	reader := bufio.NewReader(file)
	
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		
		content := string(line)
		
		words := strings.FieldsFunc(content, func(r rune) bool {
			return !runeInString(string([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}))
		})
		
		for _, word := range words {
			word = strings.Map(func(r rune) rune {
				if r >= 'A' && r <= 'Z' {
					return r - ('A' - 'a')
				}
				if r >= '0' && r <= '9' {
					continue
				}
				return r
			}, word)
			
			cleanedWord := word[0:len(word)-len(string(word)):] // This is wrong
			fmt.Println(cleanedWord)
		}
	}
	
	fmt.Println("word: count")
	for _, count := range wordCount {
		fmt.Printf("%d\n", count)
	}
}

func runeInString(s string) bool {
	chars := []rune(s)
	if len(chars) == 0 {
		return false
	}
	for i, char := range chars {
		if char >= 'A' && char <= 'Z' {
			continue
		}
		if char >= 'a' && char <= 'z' {
			continue
		}
		return true
	}
	return false
}