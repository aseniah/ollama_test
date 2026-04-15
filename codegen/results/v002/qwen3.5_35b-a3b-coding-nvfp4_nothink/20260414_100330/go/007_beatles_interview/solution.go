package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Reference date for age calculation
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Open input file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	// Skip header
	header := records[0]
	dataRows := records[1:]

	var results []Person

	for _, row := range dataRows {
		if len(row) < 7 {
			continue
		}

		// Parse Name: "FirstName MiddleName LastName" -> extract FirstName and LastName
		// Assuming format: "FirstName MiddleName LastName" or "FirstName LastName"
		// We take the first word as FirstName and last word as LastName
		nameParts := strings.Fields(row[0])
		if len(nameParts) == 0 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse Birthday: "M/D/YYYY"
		bdayStr := row[1]
		bday, err := time.Parse("1/2/2006", bdayStr)
		if err != nil {
			// Try alternative format if needed, but problem uses M/D/YYYY
			continue
		}

		// Calculate Age
		age := refDate.Year() - bday.Year()
		if bday.After(time.Date(refDate.Year(), bday.Month(), bday.Day(), 0, 0, 0, 0, time.UTC)) {
			age--
		}

		var relatives []Relative

		// Father
		if row[3] != "null" && strings.TrimSpace(row[3]) != "" {
			fNameParts := strings.Fields(row[3])
			if len(fNameParts) > 0 {
				fRel := Relative{
					FirstName:    fNameParts[0],
					LastName:     fNameParts[len(fNameParts)-1],
					Relationship: "Father",
				}
				relatives = append(relatives, fRel)
			}
		}

		// Mother
		if row[4] != "null" && strings.TrimSpace(row[4]) != "" {
			mNameParts := strings.Fields(row[4])
			if len(mNameParts) > 0 {
				mRel := Relative{
					FirstName:    mNameParts[0],
					LastName:     mNameParts[len(mNameParts)-1],
					Relationship: "Mother",
				}
				relatives = append(relatives, mRel)
			}
		}

		// Brother
		if row[5] != "null" && strings.TrimSpace(row[5]) != "" {
			bNameParts := strings.Fields(row[5])
			if len(bNameParts) > 0 {
				bRel := Relative{
					FirstName:    bNameParts[0],
					LastName:     bNameParts[len(bNameParts)-1],
					Relationship: "Brother",
				}
				relatives = append(relatives, bRel)
			}
		}

		// Sister
		if row[6] != "null" && strings.TrimSpace(row[6]) != "" {
			sNameParts := strings.Fields(row[6])
			if len(sNameParts) > 0 {
				sRel := Relative{
					FirstName:    sNameParts[0],
					LastName:     sNameParts[len(sNameParts)-1],
					Relationship: "Sister",
				}
				relatives = append(relatives, sRel)
			}
		}

		person := Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   bday.Format("2006-01-02"),
			Age:        age,
			Relatives:  relatives,
		}
		results = append(results, person)
	}

	// Output JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}

type Person struct {
	FirstName  string    `json:"FirstName"`
	LastName   string    `json:"LastName"`
	Birthday   string    `json:"Birthday"`
	Age        int       `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}