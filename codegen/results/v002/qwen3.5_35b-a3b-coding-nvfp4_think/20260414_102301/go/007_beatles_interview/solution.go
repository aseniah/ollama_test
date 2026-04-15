package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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

func parseBirthday(birthdayStr string) time.Time {
	// Handle format "M/D/YYYY" or "MM/DD/YYYY"
	layout := "1/2/2006"
	t, _ := time.Parse(layout, birthdayStr)
	return t
}

func calculateAge(birthday time.Time, referenceDate time.Time) int {
	yearDiff := referenceDate.Year() - birthday.Year()

	// Check if birthday has already occurred this year
	if referenceDate.Month() < birthday.Month() ||
		(referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
		yearDiff--
	}

	return yearDiff
}

func splitName(name string) (string, string) {
	name = strings.TrimSpace(name)
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
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

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for i, record := range records {
		if i == 0 {
			continue
		}

		name := strings.TrimSpace(record[0])
		firstName, lastName := splitName(name)

		birthdayStr := strings.TrimSpace(record[1])
		birthday := parseBirthday(birthdayStr)
		age := calculateAge(birthday, referenceDate)

		relatives := []Relative{}

		fatherName := strings.TrimSpace(record[3])
		if fatherName != "null" && fatherName != "" {
			fName, lName := splitName(fatherName)
			if fName != "" && lName != "" {
				relatives = append(relatives, Relative{
					FirstName:    fName,
					LastName:     lName,
					Relationship: "Father",
				})
			}
		}

		motherName := strings.TrimSpace(record[4])
		if motherName != "null" && motherName != "" {
			mName, lName := splitName(motherName)
			if mName != "" && lName != "" {
				relatives = append(relatives, Relative{
					FirstName:    mName,
					LastName:     lName,
					Relationship: "Mother",
				})
			}
		}

		brotherName := strings.TrimSpace(record[5])
		if brotherName != "null" && brotherName != "" {
			bName, lName := splitName(brotherName)
			if bName != "" && lName != "" {
				relatives = append(relatives, Relative{
					FirstName:    bName,
					LastName:     lName,
					Relationship: "Brother",
				})
			}
		}

		sisterName := strings.TrimSpace(record[6])
		if sisterName != "null" && sisterName != "" {
			sName, lName := splitName(sisterName)
			if sName != "" && lName != "" {
				relatives = append(relatives, Relative{
					FirstName:    sName,
					LastName:     lName,
					Relationship: "Sister",
				})
			}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.Marshal(people)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}