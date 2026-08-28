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

func calculateAge(birth, reference time.Time) int {
	age := reference.Year() - birth.Year()
	if reference.Month() < birth.Month() || (reference.Month() == birth.Month() && reference.Day() < birth.Day()) {
		age--
	}
	return age
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

	result := make([]Person, 0)

	for _, record := range records[1:] {
		// Parse name: first and last name
		nameParts := strings.Fields(record[0])
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday
		birthDate, _ := time.Parse("1/2/2006", record[1])
		birthdayStr := fmt.Sprintf("%04d-%02d-%02d", birthDate.Year(), int(birthDate.Month()), birthDate.Day())

		// Parse death date (optional)
		var deathDate time.Time
		hasDeath := record[2] != "null" && record[2] != ""
		if hasDeath {
			deathDate, _ = time.Parse("1/2/2006", record[2])
		}

		// Calculate age as of July 1, 2025 (or at death if deceased)
		var age int
		if hasDeath {
			age = calculateAge(birthDate, deathDate)
		} else {
			refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
			age = calculateAge(birthDate, refDate)
		}

		// Parse relatives
		relatives := make([]Relative, 0)
		relFields := []struct {
			index int
			rel   string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, rf := range relFields {
			val := record[rf.index]
			if val != "null" && val != "" {
				relParts := strings.Fields(val)
				relatives = append(relatives, Relative{
					FirstName:    relParts[0],
					LastName:     relParts[len(relParts)-1],
					Relationship: rf.rel,
				})
			}
		}

		result = append(result, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}