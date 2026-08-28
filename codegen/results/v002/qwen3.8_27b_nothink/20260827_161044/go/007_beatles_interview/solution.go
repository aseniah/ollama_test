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

func calculateAge(birthday time.Time, asOf time.Time) int {
	age := asOf.Year() - birthday.Year()
	if asOf.YearDay() < birthday.YearDay() {
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

	// Skip header
	header := records[0]
	data := records[1:]

	_ = header

	asOf := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var result []Person

	for _, record := range data {
		if len(record) < 7 {
			continue
		}

		fullName := strings.TrimSpace(record[0])
		nameParts := strings.SplitN(fullName, " ", 2)
		var firstName, lastName string
		if len(nameParts) == 2 {
			firstName = nameParts[0]
			lastName = nameParts[1]
		} else {
			firstName = nameParts[0]
			lastName = ""
		}

		birthdayStr := strings.TrimSpace(record[1])
		birthdayLayout := "1/2/2006"
		birthday, err := time.Parse(birthdayLayout, birthdayStr)
		if err != nil {
			birthdayLayout = "1/2/06"
			birthday, err = time.Parse(birthdayLayout, birthdayStr)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
			continue
		}

		age := calculateAge(birthday, asOf)

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := fmt.Sprintf("%04d-%02d-%02d", birthday.Year(), birthday.Month(), birthday.Day())

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: []Relative{},
		}

		// Father
		fatherStr := strings.TrimSpace(record[3])
		if fatherStr != "null" && fatherStr != "" {
			fatherParts := strings.SplitN(fatherStr, " ", 2)
			if len(fatherParts) == 2 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    fatherParts[0],
					LastName:     fatherParts[1],
					Relationship: "Father",
				})
			} else if len(fatherParts) == 1 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    fatherParts[0],
					LastName:     "",
					Relationship: "Father",
				})
			}
		}

		// Mother
		motherStr := strings.TrimSpace(record[4])
		if motherStr != "null" && motherStr != "" {
			motherParts := strings.SplitN(motherStr, " ", 2)
			if len(motherParts) == 2 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    motherParts[0],
					LastName:     motherParts[1],
					Relationship: "Mother",
				})
			} else if len(motherParts) == 1 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    motherParts[0],
					LastName:     "",
					Relationship: "Mother",
				})
			}
		}

		// Brother
		brotherStr := strings.TrimSpace(record[5])
		if brotherStr != "null" && brotherStr != "" {
			brotherParts := strings.SplitN(brotherStr, " ", 2)
			if len(brotherParts) == 2 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    brotherParts[0],
					LastName:     brotherParts[1],
					Relationship: "Brother",
				})
			} else if len(brotherParts) == 1 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    brotherParts[0],
					LastName:     "",
					Relationship: "Brother",
				})
			}
		}

		// Sister
		sisterStr := strings.TrimSpace(record[6])
		if sisterStr != "null" && sisterStr != "" {
			sisterParts := strings.SplitN(sisterStr, " ", 2)
			if len(sisterParts) == 2 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    sisterParts[0],
					LastName:     sisterParts[1],
					Relationship: "Sister",
				})
			} else if len(sisterParts) == 1 {
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    sisterParts[0],
					LastName:     "",
					Relationship: "Sister",
				})
			}
		}

		result = append(result, person)
	}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}