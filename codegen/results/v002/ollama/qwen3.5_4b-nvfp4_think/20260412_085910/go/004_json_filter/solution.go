package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name  string            `json:"name"`
	Age   int               `json:"age"`
	Active bool             `json:"active"`
	Score float64           `json:"score"`
}

func main() {
	// Open and read the input file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode JSON data into slice of records
	var records []Record
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter records where active is true and age >= 30
	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort by name ascending
	for i := 0; i < len(filtered)-1; i++ {
		for j := i + 1; j < len(filtered); j++ {
			if filtered[i].Name > filtered[j].Name {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}

	// Encode and write output to stdout
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}