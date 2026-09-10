package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(rows) == 0 {
		os.Exit(1)
	}

	// Parse header
	header := rows[0]
	// Expected columns: Name, Birthday, Died, Father, Mother, Brother, Sister
	// Map column names to indices
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[strings.TrimSpace(col)] = i
	}

	dataRows := rows[1:]

	// Sort data rows by birthday (earliest first)
	// First, parse birthdays for sorting
	type rowWithBirthday struct {
		row      []string
		birthDay time.Time
	}
	parsedRows := make([]rowWithBirthday, len(dataRows))
	for i, row := range dataRows {
		// Parse birthday in format M/D/YYYY
		birthdayStr := strings.TrimSpace(row[colIndex["Birthday"]])
		birthDay, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			// Try another format
			birthDay, err = time.Parse("01/02/2006", birthdayStr)
			if err != nil {
				// Try yet another
				birthDay, err = time.Parse("1/2/06", birthdayStr)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing birthday %s: %v\n", birthdayStr, err)
					// Skip this row or handle gracefully
					continue
				}
			}
		}
		parsedRows[i] = rowWithBirthday{row: row, birthDay: birthDay}
	}

	// Sort by birthday
	sort.Slice(parsedRows, func(i, j int) bool {
		return parsedRows[i].birthDay.Before(parsedRows[j].birthDay)
	})

	// Calculate age as of July 1, 2025
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for _, pr := range parsedRows {
		row := pr.row
		name := strings.TrimSpace(row[colIndex["Name"]])
		// Split name into first and last
		parts := strings.Fields(name)
		if len(parts) < 2 {
			continue
		}
		firstName := parts[0]
		lastName := parts[len(parts)-1]

		// Birthday in YYYY-MM-DD format
		birthdayStr := strings.TrimSpace(row[colIndex["Birthday"]])
		birthDay, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthDay, err = time.Parse("01/02/2006", birthdayStr)
			if err != nil {
				birthDay, err = time.Parse("1/2/06", birthdayStr)
				if err != nil {
					continue
				}
			}
		}
		birthdayISO := birthDay.Format("2006-01-02")

		// Calculate age as of July 1, 2025
		age := refDate.Year() - birthDay.Year()
		if refDate.Before(birthDay.AddDate(age, 0, 0)) ||
			refDate.Before(birthDay.AddDate(age-1, 0, 0)) {
			// Check if birthday hasn't occurred yet in the reference year
			birthdayThisYear := time.Date(refDate.Year(), birthDay.Month(), birthDay.Day(), 0, 0, 0, 0, time.UTC)
			if refDate.Before(birthdayThisYear) {
				age--
			}
		}

		// Build relatives list
		var relatives []Relative

		// Father
		fatherStr := strings.TrimSpace(row[colIndex["Father"]])
		if fatherStr != "" && fatherStr != "null" {
			fatherParts := strings.Fields(fatherStr)
			if len(fatherParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fatherParts[0],
					LastName:     fatherParts[len(fatherParts)-1],
					Relationship: "Father",
				})
			} else if len(fatherParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    fatherParts[0],
					LastName:     "",
					Relationship: "Father",
				})
			}
		}

		// Mother
		motherStr := strings.TrimSpace(row[colIndex["Mother"]])
		if motherStr != "" && motherStr != "null" {
			motherParts := strings.Fields(motherStr)
			if len(motherParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    motherParts[0],
					LastName:     motherParts[len(motherParts)-1],
					Relationship: "Mother",
				})
			} else if len(motherParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    motherParts[0],
					LastName:     "",
					Relationship: "Mother",
				})
			}
		}

		// Brother
		brotherStr := strings.TrimSpace(row[colIndex["Brother"]])
		if brotherStr != "" && brotherStr != "null" {
			brotherParts := strings.Fields(brotherStr)
			if len(brotherParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    brotherParts[0],
					LastName:     brotherParts[len(brotherParts)-1],
					Relationship: "Brother",
				})
			} else if len(brotherParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    brotherParts[0],
					LastName:     "",
					Relationship: "Brother",
				})
			}
		}

		// Sister
		sisterStr := strings.TrimSpace(row[colIndex["Sister"]])
		if sisterStr != "" && sisterStr != "null" {
			sisterParts := strings.Fields(sisterStr)
			if len(sisterParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sisterParts[0],
					LastName:     sisterParts[len(sisterParts)-1],
					Relationship: "Sister",
				})
			} else if len(sisterParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    sisterParts[0],
					LastName:     "",
					Relationship: "Sister",
				})
			}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayISO,
			Age:       age,
			Relatives: relatives,
		})
	}

	// Marshal to JSON
	jsonOutput, err := json.MarshalIndent(people, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}