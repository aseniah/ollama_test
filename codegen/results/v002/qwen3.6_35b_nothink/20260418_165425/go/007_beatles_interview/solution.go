package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Birthday  string `json:"Birthday"`
	Age       int    `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

type Relative struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	// Read the CSV file
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

	// Reference date for age calculation
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var results []Person

	// Skip header
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 6 {
			continue
		}

		nameParts := strings.Split(row[0], " ")
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthdayStr := row[1]
		diedStr := row[2]
		fatherStr := row[3]
		motherStr := row[4]
		brotherStr := row[5]
		sisterStr := row[6]

		// Parse birthday
		bday, err := parseDate(birthdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday for %s: %v\n", row[0], err)
			continue
		}

		// Calculate age as of July 1, 2025
		age := calculateAge(bday, refDate)

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := bday.Format("2006-01-02")

		p := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: []Relative{},
		}

		// Add father if not null
		if fatherStr != "" && fatherStr != "null" {
			nameParts := strings.Split(fatherStr, " ")
			if len(nameParts) >= 2 {
				p.Relatives = append(p.Relatives, Relative{
					FirstName:  nameParts[0],
					LastName:   nameParts[len(nameParts)-1],
					Relationship: "Father",
				})
			}
		}

		// Add mother if not null
		if motherStr != "" && motherStr != "null" {
			nameParts := strings.Split(motherStr, " ")
			if len(nameParts) >= 2 {
				p.Relatives = append(p.Relatives, Relative{
					FirstName:  nameParts[0],
					LastName:   nameParts[len(nameParts)-1],
					Relationship: "Mother",
				})
			}
		}

		// Add brother if not null
		if brotherStr != "" && brotherStr != "null" {
			nameParts := strings.Split(brotherStr, " ")
			if len(nameParts) >= 2 {
				p.Relatives = append(p.Relatives, Relative{
					FirstName:  nameParts[0],
					LastName:   nameParts[len(nameParts)-1],
					Relationship: "Brother",
				})
			}
		}

		// Add sister if not null
		if sisterStr != "" && sisterStr != "null" {
			nameParts := strings.Split(sisterStr, " ")
			if len(nameParts) >= 2 {
				p.Relatives = append(p.Relatives, Relative{
					FirstName:  nameParts[0],
					LastName:   nameParts[len(nameParts)-1],
					Relationship: "Sister",
				})
			}
		}

		results = append(results, p)
	}

	// Marshal to JSON
	output, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, string(output))
}

func parseDate(dateStr string) (time.Time, error) {
	// Format: M/D/YYYY
	t, err := time.Parse("1/2/2006", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func calculateAge(birthday time.Time, refDate time.Time) int {
	age := refDate.Year() - birthday.Year()
	// Check if birthday has not yet occurred in the reference year
	if refDate.Month() < birthday.Month() || (refDate.Month() == birthday.Month() && refDate.Day() < birthday.Day()) {
		age--
	}
	return age
}