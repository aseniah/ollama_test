package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Define a structure to represent a person
type Person struct {
	Name string
	Age  int
	Active bool
	Score float64
}

func main() {
	// Read data from JSON file
	data, err := ioutil.ReadFile("input/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON data into an array of Person structs
	var persons []Person
	if err := json.Unmarshal(data, &persons); err != nil {
		log.Fatal(err)
	}

	// Filter and sort the persons based on criteria
	filteredAndSorted := make([]Person, 0, len(persons))
	for _, person := range persons {
		if person.Active && person.Age >= 30 {
			filteredAndSorted = append(filteredAndSorted, person)
		}
	}

	// Sort the filtered array by name
	filteredAndSorted = sortByName(filteredAndSorted)

	// Output the results
	for _, person := range filteredAndSorted {
		fmt.Printf("{\"name\": \"%s\", \"age\": %d, \"active\": %t, \"score\": %.1f}\n",
			person.Name, person.Age, person.Active, person.Score)
	}
}

// Helper function to sort by name
func sortByName(persons []Person) []Person {
	return sort.Slice(persons, func(i, j int) bool {
		return strings.Compare(persons[i].Name, persons[j].Name) < 0
	})
}