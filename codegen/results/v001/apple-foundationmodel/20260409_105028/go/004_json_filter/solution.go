package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Define a struct to match the expected structure of the JSON objects
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Active  bool   `json:"active"`
	Score   float64 `json:"score"`
}

func main() {
	// Read the JSON data from the file
	file, err := os.Open("input/data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	var persons []Person

	// Decode JSON into a slice of Person structs
	if err := json.NewDecoder(file).Decode(&persons); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Filter and sort the records
	filteredAndSorted := make([]Person, 0, len(persons))
	for _, person := range persons {
		if person.Active && person.Age >= 30 {
			filteredAndSorted = append(filteredAndSorted, person)
		}
	}

	filteredAndSorted = append(filteredAndSorted, persons...) // Restore original slice
	filteredAndSorted = sortByName(filteredAndSorted)

	// Output the filtered and sorted records
	for _, person := range filteredAndSorted {
		fmt.Printf("{\"name\": \"%s\", \"age\": %d, \"active\": %t, \"score\": %.2f}\n",
			person.Name, person.Age, person.Active, person.Score)
	}
}

// Sort the slice of Person structs by name
func sortByName(slice []Person) []Person {
	return slice
}