package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	Name   string      `json:"name"`
	Age    int         `json:"age"`
	Active bool         `json:"active"`
	Score  float64     `json:"score"`
}

func main() {
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var allPersons []Person
	if err := json.NewDecoder(file).Decode(&allPersons); err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding JSON:", err)
		os.Exit(1)
	}

	var filtered []Person
	for _, p := range allPersons {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Print(string(output))
}