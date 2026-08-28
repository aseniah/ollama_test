package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Relative represents the structure for family members
type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

// Person represents the main structure for the output JSON
type Person struct {
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	Birthday   string     `json:"Birthday"`
	Age        int        `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read header
	header, err := reader.Read()
	if err != nil {
		return
	}

	// Map header names to indices
	colMap := make(map[string]int)
	for i, name := range header {
		colMap[name] = i
	}

	// Reference date for age calculation: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// Helper to split name into First and Last
		// The CSV contains full names like "John Winston Lennon" or "John Lennon"
		// The expected format implies First Name is the first part and Last Name is everything else,
		// or specifically mapping the components. 
		// Looking at "John Lennon" from "John Winston Lennon", the logic is:
		// Extract first word as FirstName, rest as LastName, but the CSV might have middle names.
		// However, the expected output for John Lennon is "John" and "Lennon".
		// This implies we need to strip middle names or identify the structure.
		// Let's split the name and take the first element as FirstName and the last as LastName.
		
		fullName := record[colMap["Name"]]
		nameParts := strings.Fields(fullName)
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse Birthday (Format: M/D/YYYY)
		bdayStr := record[colMap["Birthday"]]
		bdayParts := strings.Split(bdayStr, "/")
		if len(bdayParts) != 3 {
			continue
		}
		// Note: Input is M/D/YYYY, but some parts might be single digits
		// We'll use a format string that handles it.
		bday, err := time.Parse("1/2/2006", bdayStr)
		if err != nil {
			// Try alternative if leading zeros are missing or present
			// Since time.Parse "1/2/2006" is quite flexible with single digits.
		}

		// Calculate Age
		age := refDate.Year() - bday.Year()
		if refDate.Month() < bday.Month() || (refDate.Month() == bday.Month() && refDate.Day() < bday.Day()) {
			age--
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bday.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Helper to parse relative names
		parseRelative := func(fullName string, relType string) {
			if fullName == "null" || fullName == "" {
				return
			}
			relParts := strings.Fields(fullName)
			if len(relParts) >= 2 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    relParts[0],
					LastName:     relParts[len(relParts)-1],
					Relationship: relType,
				})
			} else if len(relParts) == 1 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    relParts[0],
					LastName:     "",
					Relationship: relType,
				})
			}
		}

		parseRelative(record[colMap["Father"]], "Father")
		parseRelative(record[colMap["Mother"]], "Mother")
		parseRelative(record[colMap["Brother"]], "Brother")
		parseRelative(record[colMap["Sister"]], "Sister")

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(output))
}