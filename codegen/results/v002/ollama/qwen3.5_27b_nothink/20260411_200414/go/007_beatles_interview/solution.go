package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// Read the input file
	data, err := ioutil.ReadFile("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header
	if len(records) < 2 {
		os.Exit(1)
	}

	var people []Person
	calcDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for _, row := range records[1:] {
		if len(row) < 7 {
			continue
		}

		// Parse Name
		nameParts := strings.Split(strings.TrimSpace(row[0]), " ")
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			lastName = nameParts[len(nameParts)-1]
		}

		// Parse Birthday
		birthTime, err := time.Parse("1/2/2006", row[1])
		if err != nil {
			continue // Skip invalid rows
		}
		birthdayISO := birthTime.Format("2006-01-02")

		// Calculate Age
		age := calcDate.Year() - birthTime.Year()
		if calcDate.YearDay() < birthTime.YearDay() {
			age--
		}

		var relatives []Relative

		// Helper to parse relative name and add to list
		addRelative := func(name, relType string) {
			if name == "" || name == "null" {
				return
			}
			// Determine if input is "First Last" or "First Middle Last"
			parts := strings.Split(name, " ")
			var relFirstName, relLastName string
			if len(parts) >= 1 {
				relFirstName = parts[0]
			}
			if len(parts) >= 2 {
				relLastName = parts[len(parts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    relFirstName,
				LastName:     relLastName,
				Relationship: relType,
			})
		}

		// Columns: Name, Birthday, Died, Father, Mother, Brother, Sister
		// Indices: 0      , 1        , 2    , 3      , 4      , 5       , 6
		addRelative(row[3], "Father")
		addRelative(row[4], "Mother")
		addRelative(row[5], "Brother")
		addRelative(row[6], "Sister")

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayISO,
			Age:       age,
			Relatives: relatives,
		}
		people = append(people, person)
	}

	// Output JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(people); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}