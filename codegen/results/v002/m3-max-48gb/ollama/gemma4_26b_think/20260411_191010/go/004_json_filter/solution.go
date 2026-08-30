package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Person represents the structure of the objects in the JSON array.
type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	// Define the path to the input file.
	inputPath := "input/data.json"

	// 1. Read the file content.
	data, err := os.ReadFile(inputPath)
	if err != nil {
		// If the file cannot be read, we exit silently to adhere to 
		// the requirement "Do not output anything else."
		return
	}

	// 2. Parse the JSON array into a slice of Person structs.
	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		return
	}

	// 3. Filter the records.
	// Criteria: active is true AND age is 30 or older.
	var filtered []Person
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	// 4. Sort the filtered records by name in ascending order.
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// 5. Marshal the filtered slice back into a JSON array.
	// We use MarshalIndent to produce a readable output, 
	// but standard Marshal would also be valid.
	output, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		return
	}

	// 6. Output the resulting JSON array to stdout.
	fmt.Println(string(output))
}