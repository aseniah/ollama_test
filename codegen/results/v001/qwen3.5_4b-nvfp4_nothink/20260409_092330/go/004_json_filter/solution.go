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
	// Open and read the input file
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Parse JSON into a slice of Person structs
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Filter records: active is true and age >= 30
	var filtered []Person
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	// Sort by name ascending (case-insensitive is often safer, but strict byte-order is standard unless specified otherwise)
	// Using strings.Compare for byte-order sort. If case-insensitivity was needed, would use strings.CcompareFold.
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Create the output JSON
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output to stdout
	fmt.Print(string(output))
}