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

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], parts[0]
	}
	return parts[0], parts[len(parts)-1]
}

func parseBirthday(birthdayStr string) string {
	// Input format: M/D/YYYY (e.g., "10/9/1940")
	// Output format: YYYY-MM-DD
	t, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		t, _ = time.Parse("1/2/2006", birthdayStr)
	}
	return t.Format("2006-01-02")
}

func calculateAge(birthdayStr string, referenceDate time.Time) int {
	bday, _ := time.Parse("2006-01-02", parseBirthday(birthdayStr))
	age := referenceDate.Year() - bday.Year()
	
	// Adjust if birthday hasn't occurred yet in the reference year
	bdayInRefYear := time.Date(referenceDate.Year(), bday.Month(), bday.Day(), 0, 0, 0, 0, referenceDate.Location())
	if bdayInRefYear.After(referenceDate) {
		age--
	}
	return age
}

func main() {
	referenceDate, _ := time.Parse("2006-01-02", "2025-07-01")
	
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

	var people []Person

	for _, record := range records[1:] {
		// Parse name
		firstName, lastName := splitName(record[0])
		birthday := parseBirthday(record[1])
		age := calculateAge(record[1], referenceDate)

		// Parse relatives
		var relatives []Relative

		// Father (column 3)
		if record[3] != "" && record[3] != "null" {
			fName, lName := splitName(record[3])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Father"})
		}

		// Mother (column 4)
		if record[4] != "" && record[4] != "null" {
			fName, lName := splitName(record[4])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Mother"})
		}

		// Brother (column 5)
		if record[5] != "" && record[5] != "null" {
			fName, lName := splitName(record[5])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Brother"})
		}

		// Sister (column 6)
		if record[6] != "" && record[6] != "null" {
			fName, lName := splitName(record[6])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Sister"})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday,
			Age:       age,
			Relatives: relatives,
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(people)
}