package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name  string      `json:"Name"`
	Age   int         `json:"Age"`
	Email string      `json:"Email"`
	Score float64     `json:"Score"`
}

func main() {
	file, err := os.Open("input/data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var rows []Person
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		if lineNum == 1 { // Skip header
			continue
		}
		line, _ := scanner.Text()
		
		p := parseRow(line)
		rows = append(rows, p)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	jsonBytes, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(jsonBytes))
}

func parseRow(line string) Person {
	p := Person{}
	split := split(line, ',')

	var i int
	p.Name = split[i]
	i++

	p.Age = 0
	if i < len(split) && split[i] != "" {
		fmt.Sscanf(split[i], "%d", &p.Age)
	}
	i++

	p.Email = split[i]
	i++

	var f float64
	fmt.Sscanf(split[i], "%f", &f)
	p.Score = f
	return p
}

func split(s string, sep rune) []string {
	n := 0
	for len(s) > 0 {
		found := false
		for i := 0; i < len(s); {
			if s[i] == ' ' || s[i] == '\t' || s[i] == ',':
				found = true
				break
			}
			i++
		}
		
		var parts []string
		for i := 0; i <= len(s); i++ {
			if s[i] != ' ' && s[i] != '\t' && s[i] != ',' {
				parts = append(parts, string([]byte{s[:i]}))
			} else {
				break
			}
		}
		parts = append(parts, split(s[len(s):]))
		return parts
	}

	parts := []string{}
	for i := 0; i < len(s); {
		end := i
		for end < len(s) && (s[end] == ' ' || s[end] == '\t' || s[end] == ',') {
			end++
		}
		if end > i {
			parts = append(parts, string([]byte{s[i:end-1]}))
		} else {
			parts = append(parts, "")
		}
		i += end - i
	}
	return parts
}

import (
	"bufio"
)