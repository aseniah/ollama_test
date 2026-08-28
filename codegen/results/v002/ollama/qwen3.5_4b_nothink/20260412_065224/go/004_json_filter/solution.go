package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Active bool  `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	// Read input data
	inputJSON, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		fmt.Fprintln(err.Error)
		return
	}

	var people []Person
	if err := json.Unmarshal(inputJSON, &people); err != nil {
		fmt.Fprintln(err.Error)
		return
	}

	// Filter people: active is true and age >= 30
	filtered := []Person{}
	for _, p := range people {
		if p.Active && p.Age >= 30 {
			filtered = append(filtered, p)
		}
	}

	// Sort by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output as JSON
	outputJSON, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Fprintln(err.Error)
		return
	}

	fmt.Println(string(outputJSON))
}