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
	// Reference date: July 1, 2025
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

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

	var people []Person

	// Skip header row
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 8 {
			continue
		}

		// Parse Name
		fullName := record[0]
		names := strings.SplitN(fullName, " ", 2)
		var firstName, lastName string
		if len(names) == 2 {
			firstName = names[0]
			lastName = names[1]
		} else {
			firstName = names[0]
			lastName = ""
		}

		// Parse Birthday: MM/DD/YYYY
		birthdayStr := record[1]
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			// Try alternative format
			birthday, err = time.Parse("01/02/2006", birthdayStr)
			if err != nil {
				continue
			}
		}

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := birthday.Format("2006-01-02")

		// Calculate age as of July 1, 2025
		age := refDate.Year() - birthday.Year()
		if refDate.Month() < birthday.Month() || (refDate.Month() == birthday.Month() && refDate.Day() < birthday.Day()) {
			age--
		}

		// Parse relatives
		relatives := []Relative{}

		// Father
		father := record[3]
		if father != "null" && father != "" {
			fNames := strings.SplitN(father, " ", 2)
			fFirst := fNames[0]
			fLast := ""
			if len(fNames) == 2 {
				fLast = fNames[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    fFirst,
				LastName:     fLast,
				Relationship: "Father",
			})
		}

		// Mother
		mother := record[4]
		if mother != "null" && mother != "" {
			mNames := strings.SplitN(mother, " ", 2)
			mFirst := mNames[0]
			mLast := ""
			if len(mNames) == 2 {
				mLast = mNames[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    mFirst,
				LastName:     mLast,
				Relationship: "Mother",
			})
		}

		// Brother
		brother := record[5]
		if brother != "null" && brother != "" {
			bNames := strings.SplitN(brother, " ", 2)
			bFirst := bNames[0]
			bLast := ""
			if len(bNames) == 2 {
				bLast = bNames[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    bFirst,
				LastName:     bLast,
				Relationship: "Brother",
			})
		}

		// Sister
		sister := record[6]
		if sister != "null" && sister != "" {
			sNames := strings.SplitN(sister, " ", 2)
			sFirst := sNames[0]
			sLast := ""
			if len(sNames) == 2 {
				sLast = sNames[1]
			}
			relatives = append(relatives, Relative{
				FirstName:    sFirst,
				LastName:     sLast,
				Relationship: "Sister",
			})
		}

		// Determine if we should include this person
		// Looking at expected output: John Lennon (died 1980) is included with age 40
		// George Harrison (died 2001) is included with age 58
		// James McCartney (alive) included with age 83
		// Ringo Starr (alive) included with age 84

		// For deceased people, age is calculated as of their death date
		diedStr := record[2]
		if diedStr != "null" && diedStr != "" {
			// Person died - use death date for age calculation
			deathDate, err := time.Parse("1/2/2006", diedStr)
			if err != nil {
				deathDate, err = time.Parse("01/02/2006", diedStr)
				if err != nil {
					continue
				}
			}
			age = deathDate.Year() - birthday.Year()
			if deathDate.Month() < birthday.Month() || (deathDate.Month() == birthday.Month() && deathDate.Day() < birthday.Day()) {
				age--
			}
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
	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Output JSON to stdout
	fmt.Println(string(output))
}