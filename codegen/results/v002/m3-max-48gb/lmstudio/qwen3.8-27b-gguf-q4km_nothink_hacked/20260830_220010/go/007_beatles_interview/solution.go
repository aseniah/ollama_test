package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

	var people []Person

	for i := 1; i < len(records); i++ {
		row := records[i]
		// Parse name into first and last
		nameParts := strings.Fields(row[0])
		var firstName, lastName string
		if len(nameParts) >= 2 {
			firstName = nameParts[0]
			lastName = nameParts[len(nameParts)-1]
		} else if len(nameParts) == 1 {
			firstName = nameParts[0]
			lastName = ""
		}

		// Birthday is already in the right format
		birthday := row[1]

		// Calculate age as of July 1, 2025
		age := calculateAge(birthday)

		// Build relatives list
		var relatives []Relative
		// Father: index 3
		if row[3] != "null" && row[3] != "" {
			fName := strings.Fields(row[3])
			if len(fName) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fName[0],
					LastName:     fName[len(fName)-1],
					Relationship: "Father",
				})
			}
		}
		// Mother: index 4
		if row[4] != "null" && row[4] != "" {
			mName := strings.Fields(row[4])
			if len(mName) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    mName[0],
					LastName:     mName[len(mName)-1],
					Relationship: "Mother",
				})
			}
		}
		// Brother: index 5
		if row[5] != "null" && row[5] != "" {
			bName := strings.Fields(row[5])
			if len(bName) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    bName[0],
					LastName:     bName[len(bName)-1],
					Relationship: "Brother",
				})
			}
		}
		// Sister: index 6
		if row[6] != "null" && row[6] != "" {
			sName := strings.Fields(row[6])
			if len(sName) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sName[0],
					LastName:     sName[len(sName)-1],
					Relationship: "Sister",
				})
			}
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday,
			Age:       age,
			Relatives: relatives,
		}
		people = append(people, person)
	}

	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output to stdout
	os.Stdout.Write(output)
}

func calculateAge(birthday string) int {
	// Birthday format: M/D/YYYY
	parts := strings.Split(birthday, "/")
	if len(parts) != 3 {
		return 0
	}
	var month, day, year int
	fmt.Sscanf(parts[0], "%d", &month)
	fmt.Sscanf(parts[1], "%d", &day)
	fmt.Sscanf(parts[2], "%d", &year)

	// Reference date: July 1, 2025
	refYear := 2025
	refMonth := 7
	refDay := 1

	age := refYear - year
	// If birthday hasn't occurred yet this year, subtract 1
	if refMonth < month || (refMonth == month && refDay < day) {
		age--
	}
	if age < 0 {
		age = 0
	}
	return age
}