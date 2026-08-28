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

func parseName(name string) (first, last string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	first = parts[0]
	if len(parts) > 1 {
		last = parts[len(parts)-1]
	} else {
		last = ""
	}
	return
}

func calculateAge(birthday string, asOf time.Time) int {
	birth, _ := time.Parse("1/2/2006", birthday)
	age := asOf.Year() - birth.Year()
	// If birthday hasn't occurred yet this year, subtract 1
	if asOf.YearDay() < birth.YearDay() {
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

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Parse header
	header := records[0]
	nameIdx := -1
	birthdayIdx := -1
	fatherIdx := -1
	motherIdx := -1
	brotherIdx := -1
	sisterIdx := -1

	for i, h := range header {
		switch strings.TrimSpace(strings.ToLower(h)) {
		case "name":
			nameIdx = i
		case "birthday":
			birthdayIdx = i
		case "father":
			fatherIdx = i
		case "mother":
			motherIdx = i
		case "brother":
			brotherIdx = i
		case "sister":
			sisterIdx = i
		}
	}

	asOf := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for _, record := range records[1:] {
		if len(record) <= nameIdx {
			continue
		}

		name := strings.TrimSpace(record[nameIdx])
		first, last := parseName(name)

		birthday := strings.TrimSpace(record[birthdayIdx])
		// Convert M/D/YYYY to YYYY-MM-DD
		var birthDate time.Time
		var birthStr string
		if b, err := time.Parse("1/2/2006", birthday); err == nil {
			birthDate = b
			birthStr = b.Format("2006-01-02")
		} else {
			birthDate = time.Now()
			birthStr = birthday
		}

		age := calculateAge(birthday, asOf)

		var relatives []Relative

		getRelative := func(idx int, relationship string) {
			if idx >= 0 && idx < len(record) {
				val := strings.TrimSpace(record[idx])
				if val != "" && strings.ToLower(val) != "null" {
					rFirst, rLast := parseName(val)
					relatives = append(relatives, Relative{
						FirstName:    rFirst,
						LastName:     rLast,
						Relationship: relationship,
					})
				}
			}
		}

		getRelative(fatherIdx, "Father")
		getRelative(motherIdx, "Mother")
		getRelative(brotherIdx, "Brother")
		getRelative(sisterIdx, "Sister")

		person := Person{
			FirstName: first,
			LastName:  last,
			Birthday:  birthStr,
			Age:       age,
			Relatives: relatives,
		}
		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}