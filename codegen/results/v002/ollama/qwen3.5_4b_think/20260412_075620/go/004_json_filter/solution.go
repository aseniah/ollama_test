package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	// Read the input file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Parse JSON data
	var data []Person
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		os.Exit(1)
	}

	// Filter records where active is true and age is 30 or older
	var filtered []Person
	for _, p := range data {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	// Sort the filtered records by name in ascending order
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output to stdout
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}