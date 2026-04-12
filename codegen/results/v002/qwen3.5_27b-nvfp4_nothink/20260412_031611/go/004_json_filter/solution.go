package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Record struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Active bool    `json:"active"`
	Score  float64 `json:"score"`
}

func main() {
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var records []Record
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&records); err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding JSON:", err)
		os.Exit(1)
	}

	filtered := make([]Record, 0)
	for _, r := range records {
		if r.Active && r.Age >= 30 {
			filtered = append(filtered, r)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	output, err := json.Marshal(filtered)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}