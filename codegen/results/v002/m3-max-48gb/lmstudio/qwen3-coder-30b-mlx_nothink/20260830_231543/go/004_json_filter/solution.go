package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	// Read the input file
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse JSON array
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Filter records where active is true and age is 30 or older
	var filtered []Person
	for _, person := range people {
		if person.Active && person.Age >= 30 {
			filtered = append(filtered, person)
		}
	}

	// Sort by name ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Name < filtered[j].Name
	})

	// Output the result as JSON
	output, err := json.Marshal(filtered)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Println(string(output))
}