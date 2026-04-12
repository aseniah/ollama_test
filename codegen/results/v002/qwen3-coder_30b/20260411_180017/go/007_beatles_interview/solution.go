package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName  string
	LastName   string
	Birthday   string
	Age        int
	Relatives  []Relative
}

type Relative struct {
	FirstName    string
	LastName     string
	Relationship string
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var people []Person

	// Skip header row
	for i, record := range records {
		if i == 0 {
			continue
		}

		// Parse name
		nameParts := strings.Split(record[0], " ")
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday
		birthday := record[1]
		birthday = convertDateFormat(birthday)

		// Calculate age as of July 1, 2025
		age := calculateAge(birthday, "2025-07-01")

		// Parse relatives
		var relatives []Relative
		if record[3] != "null" {
			relatives = append(relatives, Relative{
				FirstName:    strings.Split(record[3], " ")[0],
				LastName:     strings.Split(record[3], " ")[1],
				Relationship: "Father",
			})
		}
		if record[4] != "null" {
			relatives = append(relatives, Relative{
				FirstName:    strings.Split(record[4], " ")[0],
				LastName:     strings.Split(record[4], " ")[1],
				Relationship: "Mother",
			})
		}
		if record[5] != "null" {
			relatives = append(relatives, Relative{
				FirstName:    strings.Split(record[5], " ")[0],
				LastName:     strings.Split(record[5], " ")[1],
				Relationship: "Brother",
			})
		}
		if record[6] != "null" {
			relatives = append(relatives, Relative{
				FirstName:    strings.Split(record[6], " ")[0],
				LastName:     strings.Split(record[6], " ")[1],
				Relationship: "Sister",
			})
		}

		people = append(people, Person{
			FirstName:  firstName,
			LastName:   lastName,
			Birthday:   birthday,
			Age:        age,
			Relatives:  relatives,
		})
	}

	// Print JSON array
	fmt.Printf("[")
	for i, person := range people {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("\n  {\n    \"FirstName\": \"%s\",\n    \"LastName\": \"%s\",\n    \"Birthday\": \"%s\",\n    \"Age\": %d,\n    \"Relatives\": [\n", person.FirstName, person.LastName, person.Birthday, person.Age)
		for j, relative := range person.Relatives {
			if j > 0 {
				fmt.Printf(",\n")
			}
			fmt.Printf("      {\n        \"FirstName\": \"%s\",\n        \"LastName\": \"%s\",\n        \"Relationship\": \"%s\"\n      }", relative.FirstName, relative.LastName, relative.Relationship)
		}
		fmt.Printf("\n    ]\n  }")
	}
	fmt.Printf("\n]")
}

func convertDateFormat(date string) string {
	parts := strings.Split(date, "/")
	month := parts[0]
	day := parts[1]
	year := parts[2]

	// Ensure two-digit month and day
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}

	return fmt.Sprintf("%s-%s-%s", year, month, day)
}

func calculateAge(birthday, referenceDate string) int {
	birth, _ := time.Parse("2006-01-02", birthday)
	ref, _ := time.Parse("2006-01-02", referenceDate)

	age := ref.Year() - birth.Year()
	if ref.Month() < birth.Month() || (ref.Month() == birth.Month() && ref.Day() < birth.Day()) {
		age--
	}

	return age
}