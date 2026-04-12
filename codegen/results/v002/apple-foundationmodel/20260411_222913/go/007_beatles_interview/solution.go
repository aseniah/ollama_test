package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Person represents a person with their details and calculated age.
type Person struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Birthday       string `json:"birthday"`
	Age            int    `json:"age"`
	Relatives      map[string]string `json:"relatives"`
}

// readCSV reads data from the CSV file and returns a slice of Person.
func readCSV(filePath string) ([]Person, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var persons []Person
	reader := csv.NewReader(strings.NewReader(string(data)))
	if err := reader.ReadAll(&persons); err != nil {
		return nil, fmt.Errorf("failed to parse CSV: %w", err)
	}

	return persons, nil
}

// calculateAge calculates the age of a person as of July 1, 2025.
func calculateAge(birthday string, currentDate time.Time) int {
	birth := time.Parse("02/01/1000", birthday)
	return currentDate.Year()-birth.Year()-((currentDate.Month()-birth.Month())+1)%12
}

func main() {
	inputFile := "input/input.csv"
	expectedFile := "input/expected_format.json"

	persons, err := readCSV(inputFile)
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}

	// Set a reference date to calculate ages against
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Calculate ages and update Person structs
	for i, person := range persons {
		age := calculateAge(person.Birthday, referenceDate)
		persons[i].Age = age
		// Update Relatives field to include age if possible
		for _, relative := range person.Relatives {
			if relative["Relationship"] == "Father" {
				relative["Age"] = fmt.Sprintf("%d", age)
			} else if relative["Relationship"] == "Mother" {
				relative["Age"] = fmt.Sprintf("%d", age)
			}
		}
	}

	// Output the JSON array
	var jsonArray []map[string]interface{}
	for _, person := range persons {
		jsonArray = append(jsonArray, map[string]interface{}{
			"FirstName":      person.FirstName,
			"LastName":       person.LastName,
			"Birthday":       person.Birthday,
			"Age":            person.Age,
			"Relatives":      person.Relatives,
		})
	}

	err = json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		log.Fatalf("failed to serialize JSON: %v", err)
	}

	fmt.Println(string(err))
}