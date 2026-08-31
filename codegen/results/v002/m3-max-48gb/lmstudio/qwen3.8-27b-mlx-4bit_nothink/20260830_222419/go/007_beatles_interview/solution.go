package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// First row is header
	if len(records) < 2 {
		// No data rows
		output, _ := json.Marshal([]Person{})
		fmt.Println(string(output))
		return
	}

	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for i, record := range records {
		if i == 0 {
			continue // skip header
		}

		if len(record) < 8 {
			continue
		}

		fullName := strings.TrimSpace(record[0])
		nameParts := strings.Fields(fullName)
		var firstName, lastName string
		if len(nameParts) >= 2 {
			firstName = nameParts[0]
			lastName = nameParts[len(nameParts)-1]
		} else if len(nameParts) == 1 {
			firstName = nameParts[0]
			lastName = ""
		}

		// Parse birthday
		birthdayStr := strings.TrimSpace(record[1])
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthday, err = time.Parse("01/02/2006", birthdayStr)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday for %s: %v\n", fullName, err)
			continue
		}

		// Calculate age as of July 1, 2025
		age := referenceDate.Year() - birthday.Year()
		// If birthday hasn't occurred yet this year, subtract 1
		referenceMonthDay := time.Date(referenceDate.Year(), referenceDate.Month(), referenceDate.Day(), 0, 0, 0, 0, time.UTC)
		birthdayThisYear := time.Date(referenceDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
		if referenceMonthDay.Before(birthdayThisYear) || referenceMonthDay.Equal(time.Date(referenceDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)) && false {
			// Check if we've passed the birthday
			if time.Date(referenceDate.Year(), referenceDate.Month(), referenceDate.Day(), 0, 0, 0, 0, time.UTC).Before(time.Date(referenceDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)) {
				age--
			}
		}
		// Simpler approach:
		age = referenceDate.Year() - birthday.Year()
		if referenceDate.Month() < birthday.Month() || (referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
			age--
		}

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := birthday.Format("2006-01-02")

		// Build relatives
		var relatives []Relative

		father := strings.TrimSpace(record[3])
		if father != "" && father != "null" {
			relatives = append(relatives, parseName(father, "Father"))
		}

		mother := strings.TrimSpace(record[4])
		if mother != "" && mother != "null" {
			relatives = append(relatives, parseName(mother, "Mother"))
		}

		brother := strings.TrimSpace(record[5])
		if brother != "" && brother != "null" {
			relatives = append(relatives, parseName(brother, "Brother"))
		}

		sister := strings.TrimSpace(record[6])
		if sister != "" && sister != "null" {
			relatives = append(relatives, parseName(sister, "Sister"))
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	if people == nil {
		people = []Person{}
	}

	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func parseName(fullName, relationship string) Relative {
	parts := strings.Fields(fullName)
	var firstName, lastName string
	if len(parts) >= 2 {
		firstName = parts[0]
		lastName = parts[len(parts)-1]
	} else if len(parts) == 1 {
		firstName = parts[0]
		lastName = ""
	} else {
		firstName = ""
		lastName = ""
	}

	return Relative{
		FirstName:    firstName,
		LastName:     lastName,
		Relationship: relationship,
	}
}

// Helper to avoid unused import
var _ = strconv.Itoa
var _ = fmt.Sprint