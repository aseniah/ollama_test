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

func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		return parts[0], ""
	}
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
}

func calculateAge(birthday time.Time, refDate time.Time) int {
	age := refDate.Year() - birthday.Year()
	// If birthday hasn't occurred yet this year, subtract 1
	if refDate.Month() < birthday.Month() ||
		(refDate.Month() == birthday.Month() && refDate.Day() < birthday.Day()) {
		age--
	}
	return age
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
		fmt.Println("[]")
		return
	}

	header := records[0]
	data := records[1:]

	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var result []Person

	for _, record := range data {
		// Map header indices
		nameIdx := 0
		birthdayIdx := 1
		fatherIdx := 3
		motherIdx := 4
		brotherIdx := 5
		sisterIdx := 6

		// Parse name
		firstName, lastName := splitName(record[nameIdx])

		// Parse birthday
		birthday, err := time.Parse("1/2/2006", record[birthdayIdx])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
			continue
		}
		birthdayStr := birthday.Format("2006-01-02")

		// Calculate age
		age := calculateAge(birthday, refDate)

		// Parse relatives
		relatives := make([]Relative, 0)

		father := record[fatherIdx]
		if father != "null" && father != "" {
			fName, fLastName := splitName(father)
			relatives = append(relatives, Relative{
				FirstName:    fName,
				LastName:     fLastName,
				Relationship: "Father",
			})
		}

		mother := record[motherIdx]
		if mother != "null" && mother != "" {
			mName, mLastName := splitName(mother)
			relatives = append(relatives, Relative{
				FirstName:    mName,
				LastName:     mLastName,
				Relationship: "Mother",
			})
		}

		brother := record[brotherIdx]
		if brother != "null" && brother != "" {
			bName, bLastName := splitName(brother)
			relatives = append(relatives, Relative{
				FirstName:    bName,
				LastName:     bLastName,
				Relationship: "Brother",
			})
		}

		sister := record[sisterIdx]
		if sister != "null" && sister != "" {
			sName, sLastName := splitName(sister)
			relatives = append(relatives, Relative{
				FirstName:    sName,
				LastName:     sLastName,
				Relationship: "Sister",
			})
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		}
		result = append(result, person)
	}

	if result == nil {
		result = []Person{}
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}