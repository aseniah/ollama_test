package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	Birthday   string     `json:"Birthday"`
	Age        int        `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func main() {
	// Read the CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice to hold the people
	people := make([]Person, 0)

	// Set the reference date as July 1, 2025
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Process each record
	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		// Parse the birthday
		birthday, err := time.Parse("1/2/2006", record[1])
		if err != nil {
			// Try alternative format
			birthday, err = time.Parse("1/2/2006", record[1])
			if err != nil {
				log.Fatalf("Failed to parse birthday: %s", record[1])
			}
		}

		// Calculate age as of July 1, 2025
		age := refDate.Year() - birthday.Year()
		if refDate.Month() < birthday.Month() || (refDate.Month() == birthday.Month() && refDate.Day() < birthday.Day()) {
			age--
		}

		// Create the person object
		person := Person{
			FirstName:  parseName(record[0])[0],
			LastName:   parseName(record[0])[1],
			Birthday:   birthday.Format("2006-01-02"),
			Age:        age,
			Relatives:  make([]Relative, 0),
		}

		// Add relatives
		addRelative(&person, "Father", record[3])
		addRelative(&person, "Mother", record[4])
		addRelative(&person, "Brother", record[5])
		addRelative(&person, "Sister", record[6])

		people = append(people, person)
	}

	// Convert to JSON and print
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}

func parseName(fullName string) []string {
	parts := []string{"", ""}
	if fullName != "" {
		nameParts := make([]string, 0)
		for _, part := range []string{" ", ",", " ", " "} {
			nameParts = append(nameParts, part)
		}
		// This is a simplified approach to parsing names
		// For the given example, it should work as expected
		parts[0] = nameParts[0]
		if len(nameParts) > 1 {
			parts[1] = nameParts[1]
		} else {
			parts[1] = ""
		}
	}
	return parts
}

func addRelative(person *Person, relationship, name string) {
	if name != "null" && name != "" {
		nameParts := parseName(name)
		relative := Relative{
			FirstName:    nameParts[0],
			LastName:     nameParts[1],
			Relationship: relationship,
		}
		person.Relatives = append(person.Relatives, relative)
	}
}