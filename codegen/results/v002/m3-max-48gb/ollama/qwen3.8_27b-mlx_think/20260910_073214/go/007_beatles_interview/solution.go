package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName string     `json:"FirstName"`
	LastName  string     `json:"LastName"`
	Birthday  string     `json:"Birthday"`
	Age       int        `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		return parts[0], ""
	}
	first := parts[0]
	last := parts[len(parts)-1]
	return first, last
}

func calcAge(birthDate time.Time, refDate time.Time) int {
	age := refDate.Year() - birthDate.Year()
	// Check if birthday hasn't occurred yet in the reference year
	birthdayThisYear := time.Date(refDate.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, time.UTC)
	if refDate.Before(birthdayThisYear) {
		age--
	}
	return age
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header
	var persons []Person
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		// Parse Name
		fullName := row[0]
		firstName, lastName := splitName(fullName)

		// Parse Birthday (M/D/Y)
		birthDate, _ := time.Parse("1/2/2006", row[1])
		birthdayStr := fmt.Sprintf("%04d-%02d-%02d", birthDate.Year(), birthDate.Month(), birthDate.Day())

		// Determine age
		var age int
		diedStr := row[2]
		if strings.EqualFold(diedStr, "null") || diedStr == "" {
			// Alive: calculate as of July 1, 2025
			age = calcAge(birthDate, referenceDate)
		} else {
			// Dead: calculate at date of death
			diedDate, _ := time.Parse("1/2/2006", diedStr)
			age = calcAge(birthDate, diedDate)
		}

		// Parse relatives
		relatives := make([]Relative, 0)
		// Father
		if !strings.EqualFold(row[3], "null") && row[3] != "" {
			fn, ln := splitName(row[3])
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Father"})
		}
		// Mother
		if !strings.EqualFold(row[4], "null") && row[4] != "" {
			fn, ln := splitName(row[4])
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Mother"})
		}
		// Brother
		if !strings.EqualFold(row[5], "null") && row[5] != "" {
			fn, ln := splitName(row[5])
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Brother"})
		}
		// Sister
		if !strings.EqualFold(row[6], "null") && row[6] != "" {
			fn, ln := splitName(row[6])
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Sister"})
		}

		persons = append(persons, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(persons, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}