package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Read the input file
	inputFile := "input/data.json"
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter and sort the records
	var filteredRecords []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filteredRecords = append(filteredRecords, r)
		}
	}

	sort.Slice(filteredRecords, func(i, j int) bool {
		return filteredRecords[i].Name < filteredRecords[j].Name
	})

	// Output as JSON array to stdout
	output, err := json.Marshal(filteredRecords)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling output: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}