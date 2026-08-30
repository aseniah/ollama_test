package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName  string   `json:"FirstName"`
	LastName   string   `json:"LastName"`
	Birthday   string   `json:"Birthday"`
	Age        int      `json:"Age"`
	Relatives  []Relativ `json:"Relatives"`
}

type Relativ struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	// Read input CSV
	data, err := os.ReadFile("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Parse CSV manually (simple approach)
	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid input file\n")
		os.Exit(1)
	}

	var people []Person
	// Skip header
	for i, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 6 {
			continue
		}

		name := fields[0]
		birthdayRaw := fields[1]
		birthday, err := parseDate(birthdayRaw)
		if err != nil {
			continue
		}

		age, err := calculateAge(birthday, time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))
		if err != nil {
			continue
		}

		relatives := []Relativ{}

		// Check Father
		fatherField := fields[3]
		if fatherField != "null" && fatherField != "" {
			relatives = append(relatives, Relativ{
				FirstName:  fatherField,
				LastName:   "Father'sLastName", // We don't have surname for relative
				Relationship: "Father"
			})
		}

		// Check Mother
		motherField := fields[4]
		if motherField != "null" && motherField != "" {
			relatives = append(relatives, Relativ{
				FirstName:  motherField,
				LastName:   "Mother'sLastName",
				Relationship: "Mother"
			})
		}

		// Check Brother
		brotherField := fields[5]
		if brotherField != "null" && brotherField != "" {
			relatives = append(relatives, Relativ{
				FirstName:  brotherField,
				LastName:   "Brother'sLastName",
				Relationship: "Brother"
			})
		}

		// Check Sister
		sisterField := fields[6]
		if sisterField != "null" && sisterField != "" {
			relatives = append(relatives, Relativ{
				FirstName:  sisterField,
				LastName:   "Sister'sLastName",
				Relationship: "Sister"
			})
		}

		// For actual surname, we'd need to parse the "Name" field better
		// But looking at the input, the name field is "FirstName LastName" format
		// The surname is after the first name.
		// Example: "John Winston Lennon" -> First: John, Last: Lennon
		
		// Actually, let's re-parse the name field to get surname
		nameParts := strings.Split(name, " ")
		if len(nameParts) >= 3 {
			person := Person{
				FirstName: nameParts[0],
				LastName:  nameParts[len(nameParts)-1],
				Birthday:  birthday.Format("2006-01-02"),
				Age:       age,
			}
			
			// Sort relatives to match expected order: Father, Mother, Brother, Sister
			sort.Slice(person.Relatives, func(i, j int) bool {
				relMap := map[string]int{
					"Father":    1,
					"Mother":    2,
					"Brother":   3,
					"Sister":    4,
				}
				return relMap[person.Relatives[i].Relationship] < relMap[person.Relatives[j].Relationship]
			})

			people = append(people, person)
		}
	}

	// Output JSON
	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func parseDate(dateStr string) (time.Time, error) {
	// Format: 10/9/1940
	return time.Parse("01/02/06", dateStr)
}

func calculateAge(birthday time.Time, today time.Time) (int, error) {
	years := today.Year() - birthday.Year()
	
	// If the birthday hasn't occurred this year yet, subtract 1
	if birthday.Month() > today.Month() || 
	   (birthday.Month() == today.Month() && birthday.Day() > today.Day()) {
		years--
	}
	
	return years, nil
}