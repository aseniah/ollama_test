package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Relative struct {
	FirstName  string
	LastName   string
	Relationship string
}

type Person struct {
	FirstName  string
	LastName   string
	Birthday   string
	Age        int
	Relatives  []Relative
}

func main() {
	// Read input CSV
	data, err := os.ReadFile("input/input.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	// Skip header line
	var people []Person
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			continue
		}

		// Parse Birthday
		birthdayStr := parts[1]
		date, err := time.Parse("02/01/2006", birthdayStr)
		if err != nil {
			continue
		}

		// Calculate age as of July 1, 2025
		referDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		age := int(referDate.Sub(date).Hours() / 24 / 3600 / 24)

		// Adjust age if birthday hasn't occurred yet in the year
		if referDate.Year() > date.Year()+1 {
			// Already handled correctly by simple subtraction for this case
			// But need to check if birthday is in the past within the same year
			if referDate.Before(date) {
				age--
			}
		}

		// Parse relatives
		var relatives []Relative
		for j := 4; j < len(parts); j++ {
			relStr := strings.TrimSpace(parts[j])
			if relStr == "null" {
				continue
			}
			
			relParts := strings.Split(relStr, ",")
			firstName := strings.TrimSpace(relParts[0])
			lastName := strings.TrimSpace(relParts[1])
			relationship := strings.TrimSpace(relParts[2])
			
			relatives = append(relatives, Relative{
				FirstName:  firstName,
				LastName:   lastName,
				Relationship: relationship,
			})
		}

		people = append(people, Person{
			FirstName:  strings.TrimSpace(parts[0]),
			LastName:   strings.TrimSpace(parts[2]),
			Birthday:   date.Format("2006-01-02"),
			Age:        age,
			Relatives:  relatives,
		})
	}

	// Output JSON
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}