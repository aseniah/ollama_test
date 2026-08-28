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
	FirstName   string     `json:"FirstName"`
	LastName    string     `json:"LastName"`
	Birthday    string     `json:"Birthday"`
	Age         int        `json:"Age"`
	Relatives   []Relative `json:"Relatives"`
}

func main() {
	// Read CSV file
	data, err := os.ReadFile("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	// Parse CSV
	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CSV: %v\n", err)
		return
	}

	// Reference date for age calculation
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	// Process records (skip header)
	var people []Person
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 7 {
			continue
		}

		// Parse name
		name := row[0]
		nameParts := strings.Fields(name)
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday (M/D/YYYY format)
		birthdayStr := row[1]
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
			continue
		}

		// Calculate age as of July 1, 2025
		age := refDate.Year() - birthday.Year()
		if birthday.Month() > refDate.Month() || (birthday.Month() == refDate.Month() && birthday.Day() > refDate.Day()) {
			age--
		}

		// Create relatives
		var relatives []Relative

		// Father
		if row[3] != "null" && row[3] != "" {
			fatherParts := strings.Fields(row[3])
			relatives = append(relatives, Relative{
				FirstName:    fatherParts[0],
				LastName:     fatherParts[len(fatherParts)-1],
				Relationship: "Father",
			})
		}

		// Mother
		if row[4] != "null" && row[4] != "" {
			motherParts := strings.Fields(row[4])
			relatives = append(relatives, Relative{
				FirstName:    motherParts[0],
				LastName:     motherParts[len(motherParts)-1],
				Relationship: "Mother",
			})
		}

		// Brother
		if row[5] != "null" && row[5] != "" {
			brotherParts := strings.Fields(row[5])
			relatives = append(relatives, Relative{
				FirstName:    brotherParts[0],
				LastName:     brotherParts[len(brotherParts)-1],
				Relationship: "Brother",
			})
		}

		// Sister
		if row[6] != "null" && row[6] != "" {
			sisterParts := strings.Fields(row[6])
			relatives = append(relatives, Relative{
				FirstName:    sisterParts[0],
				LastName:     sisterParts[len(sisterParts)-1],
				Relationship: "Sister",
			})
		}

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := birthday.Format("2006-01-02")

		// Create person
		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	// Output JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}