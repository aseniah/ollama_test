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

func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		return parts[0], ""
	}
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
}

func parseDate(s string) time.Time {
	return time.Parse("1/2/2006", s)
}

func ageAt(dateOfBirth time.Time, referenceDate time.Time) int {
	age := referenceDate.Year() - dateOfBirth.Year()
	// Check if birthday hasn't occurred yet in the reference year
	birthdayThisYear := time.Date(referenceDate.Year(), dateOfBirth.Month(), dateOfBirth.Day(), 0, 0, 0, 0, time.UTC)
	if referenceDate.Before(birthdayThisYear) {
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
	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	people := make([]Person, 0)

	for _, row := range records[1:] {
		if len(row) < 8 {
			continue
		}

		// Name
		fullName := strings.TrimSpace(row[0])
		firstName, lastName := splitName(fullName)

		// Birthday
		birthdayStr := strings.TrimSpace(row[1])
		birthDate := parseDate(birthdayStr)
		birthdayFormatted := fmt.Sprintf("%04d-%02d-%02d", birthDate.Year(), birthDate.Month(), birthDate.Day())

		// Age calculation
		var age int
		diedStr := strings.TrimSpace(row[2])
		if diedStr != "" && diedStr != "null" {
			deathDate := parseDate(diedStr)
			age = ageAt(birthDate, deathDate)
		} else {
			age = ageAt(birthDate, referenceDate)
		}

		// Relatives
		relatives := make([]Relative, 0)

		addRelative := func(value string, relationship string) {
			value = strings.TrimSpace(value)
			if value == "" || value == "null" {
				return
			}
			relFirstName, relLastName := splitName(value)
			relatives = append(relatives, Relative{
				FirstName:    relFirstName,
				LastName:     relLastName,
				Relationship: relationship,
			})
		}

		addRelative(row[3], "Father")
		addRelative(row[4], "Mother")
		addRelative(row[5], "Brother")
		addRelative(row[6], "Sister")

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	_ = strconv.Itoa(0) // suppress unused import if needed

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}