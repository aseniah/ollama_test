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
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Print("[]")
		return
	}

	header := records[0]
	var colName, colBirthday, colDied, colFather, colMother, colBrother, colSister int
	for i, h := range header {
		switch strings.TrimSpace(h) {
		case "Name":
			colName = i
		case "Birthday":
			colBirthday = i
		case "Died":
			colDied = i
		case "Father":
			colFather = i
		case "Mother":
			colMother = i
		case "Brother":
			colBrother = i
		case "Sister":
			colSister = i
		}
	}

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	var people []Person

	for i := 1; i < len(records); i++ {
		record := records[i]

		// Get name parts
		fullName := strings.TrimSpace(record[colName])
		nameParts := strings.Fields(fullName)
		var firstName, lastName string
		if len(nameParts) >= 2 {
			firstName = nameParts[0]
			lastName = nameParts[len(nameParts)-1]
		} else if len(nameParts) == 1 {
			firstName = nameParts[0]
		}

		// Parse birthday
		birthdayStr := strings.TrimSpace(record[colBirthday])
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			birthday, err = time.Parse("2006-01-02", birthdayStr)
			if err != nil {
				continue
			}
		}

		// Calculate age as of July 1, 2025
		age := referenceDate.Year() - birthday.Year()
		if referenceDate.Month() < birthday.Month() ||
			(referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
			age--
		}

		// Build relatives list
		var relatives []Relative

		fatherName := strings.TrimSpace(record[colFather])
		if fatherName != "" && fatherName != "null" {
			fatherParts := strings.Fields(fatherName)
			fName := fatherParts[0]
			fLName := fatherParts[len(fatherParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     fLName,
				Relationship: "Father",
			})
		}

		motherName := strings.TrimSpace(record[colMother])
		if motherName != "" && motherName != "null" {
			motherParts := strings.Fields(motherName)
			mName := motherParts[0]
			mLName := motherParts[len(motherParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    mName,
				LastName:     mLName,
				Relationship: "Mother",
			})
		}

		brotherName := strings.TrimSpace(record[colBrother])
		if brotherName != "" && brotherName != "null" {
			brotherParts := strings.Fields(brotherName)
			bName := brotherParts[0]
			bLName := brotherParts[len(brotherParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    bName,
				LastName:     bLName,
				Relationship: "Brother",
			})
		}

		sisterName := strings.TrimSpace(record[colSister])
		if sisterName != "" && sisterName != "null" {
			sisterParts := strings.Fields(sisterName)
			sName := sisterParts[0]
			sLName := sisterParts[len(sisterParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    sName,
				LastName:     sLName,
				Relationship: "Sister",
			})
		}

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := birthday.Format("2006-01-02")

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}

		if relatives == nil {
			relatives = []Relative{}
			person.Relatives = relatives
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}