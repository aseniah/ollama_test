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

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header row
	if len(rows) < 2 {
		fmt.Println("[]")
		return
	}

	header := rows[0]
	dataRows := rows[1:]

	// Find column indices
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[col] = i
	}

	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var persons []Person

	for _, row := range dataRows {
		name := strings.TrimSpace(row[colIndex["Name"]])
		parts := strings.Fields(name)
		var firstName, lastName string
		if len(parts) >= 2 {
			firstName = parts[0]
			lastName = strings.Join(parts[1:], " ")
		} else {
			firstName = parts[0]
			lastName = ""
		}

		birthdayStr := strings.TrimSpace(row[colIndex["Birthday"]])
		birthday, _ := time.Parse("1/2/2006", birthdayStr)
		birthdayFormatted := birthday.Format("2006-01-02")

		// Calculate age as of July 1, 2025
		age := referenceDate.Year() - birthday.Year()
		if referenceDate.Month() < birthday.Month() || (referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
			age--
		}

		var relatives []Relative

		father := strings.TrimSpace(row[colIndex["Father"]])
		if father != "null" && father != "" {
			fParts := strings.Fields(father)
			if len(fParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fParts[0],
					LastName:     strings.Join(fParts[1:], " "),
					Relationship: "Father",
				})
			}
		}

		mother := strings.TrimSpace(row[colIndex["Mother"]])
		if mother != "null" && mother != "" {
			mParts := strings.Fields(mother)
			if len(mParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    mParts[0],
					LastName:     strings.Join(mParts[1:], " "),
					Relationship: "Mother",
				})
			}
		}

		brother := strings.TrimSpace(row[colIndex["Brother"]])
		if brother != "null" && brother != "" {
			bParts := strings.Fields(brother)
			if len(bParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    bParts[0],
					LastName:     strings.Join(bParts[1:], " "),
					Relationship: "Brother",
				})
			}
		}

		sister := strings.TrimSpace(row[colIndex["Sister"]])
		if sister != "null" && sister != "" {
			sParts := strings.Fields(sister)
			if len(sParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sParts[0],
					LastName:     strings.Join(sParts[1:], " "),
					Relationship: "Sister",
				})
			}
		}

		if relatives == nil {
			relatives = []Relative{}
		}

		persons = append(persons, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}