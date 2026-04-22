package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	Birthday   string     `json:"Birthday"`
	Age        int        `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func calculateAge(birthdayStr string, refDate time.Time) int {
	// Parse the birthday string
	// Expected format: M/D/YYYY or MM/DD/YYYY
	parsed, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		// Try alternate format if needed, but the problem seems consistent
		parsed, err = time.Parse("2/1/2006", birthdayStr)
		if err != nil {
			return 0
		}
	}

	age := refDate.Year() - parsed.Year()
	// Check if the birthday has occurred this year
	refDateThisYear := time.Date(refDate.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, parsed.Location())
	if refDate.Before(refDateThisYear) {
		age--
	}
	return age
}

func main() {
	// Read the input CSV
	file, err := ioutil.ReadFile("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	reader := csv.NewReader(strings.NewReader(string(file)))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CSV: %v\n", err)
		return
	}

	// Define the reference date: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var persons []Person

	// Skip the header row
	for i, record := range records {
		if i == 0 {
			continue
		}

		if len(record) < 6 {
			continue
		}

		nameStr := record[0]
		birthdayStr := record[1]
		// diedStr := record[2] // Not used in output
		fatherStr := record[3]
		motherStr := record[4]
		brotherStr := record[5]
		// sisterStr := record[6] // The CSV has 6 columns based on the header, but let's check
		// Header: Name,Birthday,Died,Father,Mother,Brother,Sister
		// So there are 7 columns. Let's handle up to 7.
		
		// Re-read with a more flexible approach if needed, but let's assume standard CSV parsing handles it.
		// Actually, records should have 7 elements. Let's adjust.
		
		// Let's re-parse carefully.
		// The split might have been done wrong if I assumed 6.
		// Let's just use the record slice directly.
		
		if len(record) < 7 {
			continue
		}

		sisterStr := record[6]

		// Split name into First and Last
		// The name format is "FirstName MiddleName LastName" or "FirstName LastName"
		// Looking at the expected output:
		// "John Winston Lennon" -> FirstName: "John", LastName: "Lennon"
		// "James Paul McCartney" -> FirstName: "James", LastName: "McCartney"
		// "Ringo Starr" -> FirstName: "Ringo", LastName: "Starr"
		// "George Harrison" -> FirstName: "George", LastName: "Harrison"
		// It seems the last part is the LastName, and the first part is the FirstName.
		// What about middle names? They seem to be ignored or part of the first name?
		// Expected output only has FirstName and LastName.
		// John Winston Lennon -> John Lennon. So "Winston" is dropped? Or is "John" the first name?
		// Let's assume FirstName is the first word, LastName is the last word.
		
		nameParts := strings.Fields(nameStr)
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Calculate age
		age := calculateAge(birthdayStr, refDate)
		
		// Format birthday as YYYY-MM-DD for output
		// Input format is M/D/YYYY or MM/DD/YYYY
		// We need to output YYYY-MM-DD
		parsedBday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			parsedBday, err = time.Parse("2/1/2006", birthdayStr)
			if err != nil {
				// Fallback or error
				fmt.Fprintf(os.Stderr, "Error parsing birthday: %v for %s\n", err, nameStr)
				continue
			}
		}
		birthdayFormatted := parsedBday.Format("2006-01-02")

		var relatives []Relative

		// Add Father
		if fatherStr != "" && fatherStr != "null" {
			fNameParts := strings.Fields(fatherStr)
			if len(fNameParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fNameParts[0],
					LastName:     fNameParts[len(fNameParts)-1],
					Relationship: "Father",
				})
			} else if len(fNameParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    fNameParts[0],
					LastName:     "",
					Relationship: "Father",
				})
			}
		}

		// Add Mother
		if motherStr != "" && motherStr != "null" {
			mNameParts := strings.Fields(motherStr)
			if len(mNameParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    mNameParts[0],
					LastName:     mNameParts[len(mNameParts)-1],
					Relationship: "Mother",
				})
			} else if len(mNameParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    mNameParts[0],
					LastName:     "",
					Relationship: "Mother",
				})
			}
		}

		// Add Brother
		if brotherStr != "" && brotherStr != "null" {
			bNameParts := strings.Fields(brotherStr)
			if len(bNameParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    bNameParts[0],
					LastName:     bNameParts[len(bNameParts)-1],
					Relationship: "Brother",
				})
			} else if len(bNameParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    bNameParts[0],
					LastName:     "",
					Relationship: "Brother",
				})
			}
		}

		// Add Sister
		if sisterStr != "" && sisterStr != "null" {
			sNameParts := strings.Fields(sisterStr)
			if len(sNameParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sNameParts[0],
					LastName:     sNameParts[len(sNameParts)-1],
					Relationship: "Sister",
				})
			} else if len(sNameParts) == 1 {
				relatives = append(relatives, Relative{
					FirstName:    sNameParts[0],
					LastName:     "",
					Relationship: "Sister",
				})
			}
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}
		persons = append(persons, person)
	}

	// Output as JSON
	jsonData, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}