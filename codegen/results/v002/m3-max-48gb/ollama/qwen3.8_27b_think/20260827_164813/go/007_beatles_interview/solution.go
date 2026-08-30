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

func parseName(fullName string) (string, string) {
	parts := strings.Split(strings.TrimSpace(fullName), " ")
	firstName := parts[0]
	lastName := parts[len(parts)-1]
	return firstName, lastName
}

func formatDate(dateStr string) string {
	t, _ := time.Parse("1/2/2006", dateStr)
	return t.Format("2006-01-02")
}

func calculateAge(birthStr string, diedStr string) int {
	birth, _ := time.Parse("1/2/2006", birthStr)

	var reference time.Time
	if diedStr != "null" && diedStr != "" {
		reference, _ = time.Parse("1/2/2006", diedStr)
	} else {
		reference = time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	}

	age := reference.Year() - birth.Year()
	if reference.YearDay() < birth.YearDay() {
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
	// Skip header row
	for _, record := range records[1:] {
		name := strings.TrimSpace(record[0])
		birthday := strings.TrimSpace(record[1])
		died := strings.TrimSpace(record[2])
		father := strings.TrimSpace(record[3])
		mother := strings.TrimSpace(record[4])
		brother := strings.TrimSpace(record[5])
		sister := strings.TrimSpace(record[6])

		firstName, lastName := parseName(name)

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  formatDate(birthday),
			Age:       calculateAge(birthday, died),
			Relatives: []Relative{},
		}

		if father != "null" && father != "" {
			fn, ln := parseName(father)
			person.Relatives = append(person.Relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Father"})
		}
		if mother != "null" && mother != "" {
			fn, ln := parseName(mother)
			person.Relatives = append(person.Relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Mother"})
		}
		if brother != "null" && brother != "" {
			fn, ln := parseName(brother)
			person.Relatives = append(person.Relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Brother"})
		}
		if sister != "null" && sister != "" {
			fn, ln := parseName(sister)
			person.Relatives = append(person.Relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Sister"})
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