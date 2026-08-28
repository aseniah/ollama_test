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

func parseName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	first := parts[0]
	last := "null"
	if len(parts) > 1 {
		last = strings.Join(parts[1:], " ")
	}
	return first, last
}

func parseNameToRelative(name string, relType string) (Relative, bool) {
	name = strings.TrimSpace(name)
	if name == "" || name == "null" {
		return Relative{}, false
	}
	first, last := parseName(name)
	return Relative{
		FirstName:    first,
		LastName:     last,
		Relationship: relType,
	}, true
}

func calculateAge(birthday string) int {
	layout := "1/2/2006" // Go's reference time format for M/D/YYYY
	birthDate, err := time.Parse(layout, birthday)
	if err != nil {
		// Fallback to YYYY-MM-DD if the above fails (unlikely here but safe)
		birthDate, _ = time.Parse("2006-01-02", birthday)
	}
	
	calcDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	age := calcDate.Year() - birthDate.Year()

	if calcDate.Before(birthDate.AddDate(age, 0, 0)) {
		age--
	}

	return age
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening input CSV:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	var people []Person

	// Skip header row
	for _, record := range records[1:] {
		if len(record) < 7 {
			continue
		}

		name := record[0]
		birthdayRaw := record[1]
		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		firstName, lastName := parseName(name)
		birthdayFormatted, err := time.Parse("1/2/2006", birthdayRaw).MarshalText()
		if err != nil {
			// Fallback in case of parsing error (should not happen with valid data)
			birthdayFormatted = []byte("\"" + birthdayRaw + "\"")
		}
		// Convert byte slice to string
		var bDayStr string
		if len(birthdayFormatted) > 0 {
			bDayStr = time.Date(0, time.Month(birthDate), birthDate.Day(), 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		}

		// Re-parse properly for formatting
		layout := "1/2/2006"
		birthTime, _ := time.Parse(layout, birthdayRaw)
		bDayStr = birthTime.Format("2006-01-02")

		age := calculateAge(birthdayRaw)

		var relatives []Relative

		if rel, ok := parseNameToRelative(father, "Father"); ok {
			relatives = append(relatives, rel)
		}
		if rel, ok := parseNameToRelative(mother, "Mother"); ok {
			relatives = append(relatives, rel)
		}
		if rel, ok := parseNameToRelative(brother, "Brother"); ok {
			relatives = append(relatives, rel)
		}
		if rel, ok := parseNameToRelative(sister, "Sister"); ok {
			relatives = append(relatives, rel)
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bDayStr,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	jsonData, err := json.MarshalIndent(people, "", "   ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}