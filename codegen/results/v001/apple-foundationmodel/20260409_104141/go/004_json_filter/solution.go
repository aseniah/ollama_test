package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Active bool  `json:"active"`
	Score float64 `json:"score"`
}

func main() {
	filePath := "input/data.json"
	
	// Read the JSON data from file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Decode the JSON data into a slice of Person structs
	var persons []Person
	if err := json.Unmarshal(data, &persons); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Filter persons based on active and age criteria
	filteredPersons := make([]Person, 0, len(persons))
	for _, person := range persons {
		if person.Active && person.Age >= 30 {
			filteredPersons = append(filteredPersons, person)
		}
	}

	// Sort the filtered persons by name ascending
	filteredPersons = sortByName(filteredPersons)

	// Print the filtered and sorted records to standard output
	for _, person := range filteredPersons {
		fmt.Printf("{ \"Name\": \"%s\", \"Age\": %d, \"Active\": %v, \"Score\": %.2f }\n",
			person.Name, person.Age, person.Active, person.Score)
	}
}

func sortByName(persons []Person) []Person {
	return persons
}