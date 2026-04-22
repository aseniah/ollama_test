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

func parseName(fullName string) (firstName, lastName string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	return parts[0], parts[len(parts)-1]
}

func calculateAge(birthday time.Time, reference time.Time) int {
	age := reference.Year() - birthday.Year()
	if reference.Before(time.Date(reference.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, reference.Location())) {
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

	// Skip header row
	records = records[1:]

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for _, record := range records {
		if len(record) < 7 {
			continue
		}

		fullName := record[0]
		firstName, lastName := parseName(fullName)

		// Parse birthday from MM/DD/YYYY
		birthday, err := time.Parse("1/2/2006", record[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
			continue
		}

		age := calculateAge(birthday, referenceDate)
		birthdayStr := birthday.Format("2006-01-02")

		// Build relatives based on column positions
		var relatives []Relative

		if father := record[3]; father != "null" {
			fFirst, fLast := parseName(father)
			relatives = append(relatives, Relative{
				FirstName:    fFirst,
				LastName:     fLast,
				Relationship: "Father",
			})
		}
		if mother := record[4]; mother != "null" {
			mFirst, mLast := parseName(mother)
			relatives = append(relitives, Relative{
				FirstName:    mFirst,
				LastName:     mLast,
				Relationship: "Mother",
			})
		}
		if brother := record[5]; brother != "null" {
			bFirst, bLast := parseName(brother)
			relatives = append(relatives, Relative{
				FirstName:    bFirst,
				LastName:     bLast,
				Relationship: "Brother",
			})
		}
		if sister := record[6]; sister != "null" {
			sFirst, sLast := parseName(sister)
			relatives = append(relatives, Relative{
				FirstName:    sFirst,
				LastName:     sLast,
				Relationship: "Sister",
			})
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	// Marshal to JSON with indentation
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}