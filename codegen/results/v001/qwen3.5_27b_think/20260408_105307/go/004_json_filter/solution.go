package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Entry struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the input file
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse JSON into a slice of Entries
	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter entries where active is true and age is >= 30
	var filtered []Entry
	for _, e := range entries {
		if e.Active && e.Age >= 30 {
			filtered = append(filtered, e)
		}
	}

	// Sort the filtered entries by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the result back to JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Write output to stdout
	if _, err := os.Stdout.Write(output); err != nil {
		fmt.Fprintf(os.Stderr, "error writing output: %v\n", err)
		os.Exit(1)
	}
}