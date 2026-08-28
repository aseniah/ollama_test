package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName  string   `json:"FirstName"`
	LastName   string   `json:"LastName"`
	Birthday   string   `json:"Birthday"`
	Age        int      `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

type Relative struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Calculate reference date: July 1, 2025
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			os.Exit(1)
		}

		// Skip header row
		if len(record) < 5 {
			continue
		}

		// Parse name
		name := record[0]
		parts := strings.Fields(name)
		if len(parts) == 1 {
			// Handle cases like "John Winston Lennon" - need first name and last name
			firstName := parts[0]
			lastName := strings.Join(parts[1:], " ")
		} else {
			firstName := parts[0]
			lastName := strings.Join(parts[1:], " ")
		}

		// Parse birthday
		birthdayStr := record[1]
		parts := strings.Split(birthdayStr, "/")
		if len(parts) == 3 {
			birthDate, err := time.Parse("02/01/06", birthdayStr)
			if err != nil {
				fmt.Println("Error parsing birthday:", err)
				os.Exit(1)
			}
			birthdayStr = birthDate.Format("2006-01-02")
		}

		// Calculate age as of July 1, 2025
		age := calculateAge(referenceDate, birthDate)

		// Parse relatives
		var relatives []Relative
		if record[3] != "null" {
			relatives = append(relatives, Relative{
				FirstName:  record[3],
				LastName:   record[4],
				Relationship: "Father",
			})
		}
		if record[5] != "null" {
			relatives = append(relatives, Relative{
				FirstName:  record[5],
				LastName:   record[6],
				Relationship: "Mother",
			})
		}
		if record[7] != "null" {
			relatives = append(relatives, Relative{
				FirstName:  record[7],
				LastName:   record[8],
				Relationship: "Brother",
			})
		}
		if record[9] != "null" {
			relatives = append(relatives, Relative{
				FirstName:  record[9],
				LastName:   record[10],
				Relationship: "Sister",
			})
		}

		person := Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   birthdayStr,
			Age:        age,
			Relatives:  relatives,
		}
		people = append(people, person)
	}

	// Output JSON
	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}

// calculateAge calculates the age of a person as of the reference date
func calculateAge(refDate, birthDate time.Time) int {
	diff := refDate.Sub(birthDate)
	years := int(diff / time.Hour) / time.Hour / time.Second / (time.Hour * time.Hour)
	age := refDate.Year() - birthDate.Year()

	// Adjust for month/day difference
	if refDate.Month() > birthDate.Month() ||
		(refDate.Month() == birthDate.Month() && refDate.Day() < birthDate.Day()) {
		age--
	}

	return age
}

// Helper import
import "io"