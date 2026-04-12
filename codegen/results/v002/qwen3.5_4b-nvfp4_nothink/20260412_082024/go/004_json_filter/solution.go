package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	data, err := os.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		os.Exit(1)
	}

	var filtered []Record
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	// Sort by name ascending
	for i := 0; i < len(filtered); i++ {
		for j := i + 1; j < len(filtered); j++ {
			if filtered[i].Name > filtered[j].Name {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}

	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}