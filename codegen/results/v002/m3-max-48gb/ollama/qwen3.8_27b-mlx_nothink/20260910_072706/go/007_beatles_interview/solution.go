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
	var people []Person

	// Reference date: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for i, record := range records {
		if i == 0 {
			continue // skip header
		}

		// Parse name
		nameParts := strings.Fields(record[0])
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday
		birthdayStr := strings.TrimSpace(record[1])
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday %q: %v\n", birthdayStr, err)
			os.Exit(1)
		}

		// Calculate age as of July 1, 2025
		age := refDate.Year() - birthday.Year()
		// If birthday hasn't occurred yet in 2025, subtract 1
		birthdayThisYear := time.Date(refDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
		if refDate.Before(birthdayThisYear) {
			age--
		}

		// Format birthday
		birthdayFormatted := birthday.Format("2006-01-02")

		// Build relatives
		var relatives []Relative

		// Father
		fatherStr := strings.TrimSpace(record[3])
		if fatherStr != "" && !strings.EqualFold(fatherStr, "null") {
			fatherParts := strings.Fields(fatherStr)
			if len(fatherParts) >= 2 {
				fatherFirst := fatherParts[0]
				fatherLast := fatherParts[len(fatherParts)-1]
				relatives = append(relatives, Relative{
					FirstName:    fatherFirst,
					LastName:     fatherLast,
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
		motherStr := strings.TrimSpace(record[4])
		if motherStr != "" && !strings.EqualFold(motherStr, "null") {
			motherParts := strings.Fields(motherStr)
			if len(motherParts) >= 2 {
				motherFirst := motherParts[0]
				motherLast := motherParts[len(motherParts)-1]
				relatives = append(relatives, Relative{
					FirstName:    motherFirst,
					LastName:     motherLast,
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
		brotherStr := strings.TrimSpace(record[5])
		if brotherStr != "" && !strings.EqualFold(brotherStr, "null") {
			brotherParts := strings.Fields(brotherStr)
			if len(brotherParts) >= 2 {
				brotherFirst := brotherParts[0]
				brotherLast := brotherParts[len(brotherParts)-1]
				relatives = append(relatives, Relative{
					FirstName:    brotherFirst,
					LastName:     brotherLast,
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
		sisterStr := strings.TrimSpace(record[6])
		if sisterStr != "" && !strings.EqualFold(sisterStr, "null") {
			sisterParts := strings.Fields(sisterStr)
			if len(sisterParts) >= 2 {
				sisterFirst := sisterParts[0]
				sisterLast := sisterParts[len(sisterParts)-1]
				relatives = append(relatives, Relative{
					FirstName:    sisterFirst,
					LastName:     sisterLast,
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

		// Use first name from the expected format (not full name from CSV)
		// Looking at expected output: John, James, Ringo, George
		// These are the first names from the name column
		// But James Paul McCartney -> James, John Winston Lennon -> John
		// Ringo Starr -> Ringo (single first name), George Harrison -> George
		// So just take the first word as first name and last word as last name

		// Initialize relatives slice even if empty to avoid null
		if relatives == nil {
			relatives = []Relative{}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	// Marshal to JSON
	jsonOutput, err := json.MarshalIndent(people, "", "   ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}