package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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

func parseDate(mdy string) (month, day, year int) {
	parts := strings.Split(mdy, "/")
	fmt.Sscanf(parts[0], "%d", &month)
	fmt.Sscanf(parts[1], "%d", &day)
	fmt.Sscanf(parts[2], "%d", &year)
	return
}

func formatISO(month, day, year int) string {
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

func ageAsOf(bm, bd, by, rm, rd, ry int) int {
	age := ry - by
	// If the birthday hasn't occurred yet in the reference year, subtract 1
	if rm < bm || (rm == bm && rd < bd) {
		age--
	}
	return age
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

		// Parse name: First [Middle...] Last
		nameParts := strings.Fields(row[0])
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday
		bm, bd, by := parseDate(row[1])
		birthdayISO := formatISO(bm, bd, by)

		// Calculate age
		var age int
		diedStr := row[2]
		if strings.TrimSpace(diedStr) == "null" || strings.TrimSpace(diedStr) == "" {
			// Alive: age as of July 1, 2025
			age = ageAsOf(bm, bd, by, refMonth, refDay, refYear)
		} else {
			// Deceased: age at death
			dm, dd, dy := parseDate(diedStr)
			age = ageAsOf(bm, bd, by, dm, dd, dy)
		}

		// Parse relatives
		relatives := make([]Relative, 0)

		relFields := []string{row[3], row[4], row[5], row[6]}
		relNames := []string{"Father", "Mother", "Brother", "Sister"}

		for j, rf := range relFields {
			rf = strings.TrimSpace(rf)
			if rf == "null" || rf == "" {
				continue
			}
			rParts := strings.Fields(rf)
			rFirst := rParts[0]
			rLast := rParts[len(rParts)-1]
			relatives = append(relatives, Relative{
				FirstName:    rFirst,
				LastName:     rLast,
				Relationship: relNames[j],
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayISO,
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