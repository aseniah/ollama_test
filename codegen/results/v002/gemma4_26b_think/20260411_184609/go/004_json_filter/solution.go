package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Person represents the structure of the objects in the JSON input.
type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the path to the input file.
	filePath := "input/data.json"

	// Read the content of the JSON file.
	data, err := os.ReadFile(filePath)
	if err != nil {
		// If there is an error reading the file, we exit silently 
		// to adhere to the "Do not output anything else" requirement.
		return
	}

	// Unmarshal the JSON array into a slice of Person structs.
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return
	}

	// Filter the slice based on the criteria:
	// active is true AND age is 30 or older.
	var filtered []Person
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	// Sort the filtered slice by the "name" field in ascending order.
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Marshal the filtered and sorted slice back into a JSON array.
	// Using MarshalIndent to provide a clean, readable output format.
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	// Output the resulting JSON array to stdout.
	fmt.Println(string(output))
}