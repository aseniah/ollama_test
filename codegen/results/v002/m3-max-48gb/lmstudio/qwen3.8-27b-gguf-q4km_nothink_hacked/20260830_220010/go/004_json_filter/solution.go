package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter: active is true and age >= 30
	var filtered []Person
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	// Sort by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal to JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}