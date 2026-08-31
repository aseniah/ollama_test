package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseDate(s string) (month, day, year int) {
	parts := strings.Split(s, "/")
	month, _ = strconv.Atoi(parts[0])
	day, _ = strconv.Atoi(parts[1])
	year, _ = strconv.Atoi(parts[2])
	return
}

func formatDate(month, day, year int) string {
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

func calculateAge(birthMonth, birthDay, birthYear int, refMonth, refDay, refYear int) int {
	age := refYear - birthYear
	if refMonth < birthMonth || (refMonth == birthMonth && refDay < birthDay) {
		age--
	}
	return age
}

func splitName(fullName string) (first, last string) {
	parts := strings.Fields(fullName)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	if len(parts) >= 3 {
		first = parts[0]
		last = parts[len(parts)-1]
		return
	}
	return parts[0], ""
}

func main() {
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Reference date: July 1, 2025
	refMonth, refDay, refYear := 7, 1, 2025

	people := make([]Person, 0)

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}

		// Name
		fullName := strings.TrimSpace(row[0])
		firstName, lastName := splitName(fullName)

		// Birthday
		bm, bd, by := parseDate(strings.TrimSpace(row[1]))
		birthdayStr := formatDate(bm, bd, by)

		// Died
		died := strings.TrimSpace(row[2])

		// Age
		var age int
		if died != "" && died != "null" {
			dm, dd, dy := parseDate(died)
			age = calculateAge(bm, bd, by, dm, dd, dy)
		} else {
			age = calculateAge(bm, bd, by, refMonth, refDay, refYear)
		}

		// Relatives
		relatives := make([]Relative, 0)
		relFields := []string{
			"Father", "Mother", "Brother", "Sister",
		}
		relValues := []string{
			strings.TrimSpace(row[3]),
			strings.TrimSpace(row[4]),
			strings.TrimSpace(row[5]),
			strings.TrimSpace(row[6]),
		}

		for j, val := range relValues {
			if val == "" || val == "null" {
				continue
			}
			rFirst, rLast := splitName(val)
			relatives = append(relatives, Relative{
				FirstName:    rFirst,
				LastName:     rLast,
				Relationship: relFields[j],
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}