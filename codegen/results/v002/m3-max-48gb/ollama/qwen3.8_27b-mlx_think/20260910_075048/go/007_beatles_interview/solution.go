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

func splitName(name string) (string, string) {
	parts := strings.Fields(name)
	first := parts[0]
	last := parts[len(parts)-1]
	return first, last
}

func calcAge(birth time.Time, ref time.Time) int {
	age := ref.Year() - birth.Year()
	// Check if birthday hasn't occurred yet in the reference year
	if ref.Month() < birth.Month() || (ref.Month() == birth.Month() && ref.Day() < birth.Day()) {
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

	var people []Person

	for i, rec := range records {
		if i == 0 {
			// Skip header
			continue
		}

		// Parse name
		firstName, lastName := splitName(rec[0])

		// Parse birthday
		birthDate, err := time.Parse("1/2/2006", rec[1])
		if err != nil {
			birthDate, err = time.Parse("1/2/2006", rec[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
				os.Exit(1)
			}
		}
		birthdayStr := birthDate.Format("2006-01-02")

		// Calculate age
		refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		if rec[2] != "null" {
			// Person is dead, calculate age at death
			deathDate, err := time.Parse("1/2/2006", rec[2])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing death date: %v\n", err)
				os.Exit(1)
			}
			refDate = deathDate
		}
		age := calcAge(birthDate, refDate)

		// Build relatives
		var relatives []Relative

		// Father
		if rec[3] != "null" {
			fName, lName := splitName(rec[3])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Father"})
		}
		// Mother
		if rec[4] != "null" {
			fName, lName := splitName(rec[4])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Mother"})
		}
		// Brother
		if rec[5] != "null" {
			fName, lName := splitName(rec[5])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Brother"})
		}
		// Sister
		if rec[6] != "null" {
			fName, lName := splitName(rec[6])
			relatives = append(relatives, Relative{FirstName: fName, LastName: lName, Relationship: "Sister"})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(people, "", "   ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}