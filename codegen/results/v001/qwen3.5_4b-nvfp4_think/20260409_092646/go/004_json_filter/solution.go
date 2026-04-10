package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name string  `json:"name"`
	Ages int      `json:"age"`
	Active bool   `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data []Record
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	var filtered []Record
	for _, record := range data {
		if record.Active && record.Age >= 30 {
			filtered = append(filtered, record)
		}
	}

	for i := len(filtered) - 1; i > 0; i-- {
		j := i - 1
		if filtered[j].Name > filtered[i].Name {
			filtered[j], filtered[i] = filtered[i], filtered[j]
		}
	}

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(filtered)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}