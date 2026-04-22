package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
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

func calculateAge(birthdayStr string, referenceDate time.Time) int {
	// Parse birthday in format MM/DD/YYYY
	parts := strings.Split(birthdayStr, "/")
	if len(parts) != 3 {
		return 0
	}
	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0
	}

	birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// Calculate age as of referenceDate
	age := referenceDate.Year() - birthday.Year()
	// Check if birthday has occurred yet in the reference year
	refDateThisYear := time.Date(referenceDate.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)
	if referenceDate.Before(refDateThisYear) {
		age--
	}

	return age
}

func main() {
	referenceDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	// Skip header
	if len(records) < 1 {
		fmt.Println("[]")
		return
	}

	var people []Person

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		// Parse Name: "First Middle Last" -> First and Last
		nameParts := strings.Fields(record[0])
		firstName := ""
		lastName := ""
		if len(nameParts) == 0 {
			continue
		}
		firstName = nameParts[0]
		if len(nameParts) == 1 {
			lastName = ""
		} else {
			// Last name is the last part
			lastName = nameParts[len(nameParts)-1]
		}

		birthdayStr := record[1]
		fatherName := record[3]
		motherName := record[4]
		brotherName := record[5]
		sisterName := record[6]

		// Calculate age
		age := calculateAge(birthdayStr, referenceDate)

		// Format birthday as YYYY-MM-DD
		birthdayParts := strings.Split(birthdayStr, "/")
		if len(birthdayParts) == 3 {
			month := birthdayParts[0]
			day := birthdayParts[1]
			year := birthdayParts[2]
			// Pad month and day to 2 digits
			if len(month) == 1 {
				month = "0" + month
			}
			if len(day) == 1 {
				day = "0" + day
			}
			birthdayStr = year + "-" + month + "-" + day
		}

		var relatives []Relative

		if fatherName != "" && fatherName != "null" {
			fNameParts := strings.Fields(fatherName)
			fFirstName := fNameParts[0]
			fLastName := ""
			if len(fNameParts) > 1 {
				fLastName = fNameParts[len(fNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    fFirstName,
				LastName:     fLastName,
				Relationship: "Father",
			})
		}

		if motherName != "" && motherName != "null" {
			mNameParts := strings.Fields(motherName)
			mFirstName := mNameParts[0]
			mLastName := ""
			if len(mNameParts) > 1 {
				mLastName = mNameParts[len(mNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    mFirstName,
				LastName:     mLastName,
				Relationship: "Mother",
			})
		}

		if brotherName != "" && brotherName != "null" {
			bNameParts := strings.Fields(brotherName)
			bFirstName := bNameParts[0]
			bLastName := ""
			if len(bNameParts) > 1 {
				bLastName = bNameParts[len(bNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    bFirstName,
				LastName:     bLastName,
				Relationship: "Brother",
			})
		}

		if sisterName != "" && sisterName != "null" {
			sNameParts := strings.Fields(sisterName)
			sFirstName := sNameParts[0]
			sLastName := ""
			if len(sNameParts) > 1 {
				sLastName = sNameParts[len(sNameParts)-1]
			}
			relatives = append(relatives, Relative{
				FirstName:    sFirstName,
				LastName:     sLastName,
				Relationship: "Sister",
			})
		}

		// Ensure relatives is an empty array, not null
		if relatives == nil {
			relatives = []Relative{}
		}

		person := Person{
			FirstName:   firstName,
			LastName:    lastName,
			Birthday:    birthdayStr,
			Age:         age,
			Relatives:   relatives,
		}

		people = append(people, person)
	}

	// Output as JSON array
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "   ")
	encoder.Encode(people)
}