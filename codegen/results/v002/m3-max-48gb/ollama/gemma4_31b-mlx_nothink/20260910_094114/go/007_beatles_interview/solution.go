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

func splitName(fullName string) (string, string) {
	fullName = strings.TrimSpace(fullName)
	if fullName == "null" || fullName == "" {
		return "", ""
	}
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	// Based on the expected format, "John Winston Lennon" becomes "John" and "Lennon"
	// It seems to take the first part as First Name and last part as Last Name
	return parts[0], parts[len(parts)-1]
}

func calculateAge(birthdayStr string, targetDate time.Time) int {
	// Input format is M/D/YYYY
	t, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		return 0
	}

	age := targetDate.Year() - t.Year()
	if targetDate.YearDay() < t.YearDay() {
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

	// Target date: July 1, 2025
	targetDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)
	var result []Person

	// Skip header
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 7 {
			continue
		}

		fullFirstName, fullLastName := splitName(row[0])
		
		p := Person{
			FirstName: fullFirstName,
			LastName:  fullLastName,
			Birthday:  "", // Will format to YYYY-MM-DD
			Age:       calculateAge(row[1], targetDate),
			Relatives: []Relative{},
		}

		// Format Birthday to YYYY-MM-DD
		t, err := time.Parse("1/2/2006", row[1])
		if err == nil {
			p.Birthday = t.Format("2006-01-02")
		}

		// Map Relatives
		relationalCols := []struct {
			index int
			label string
		}{
			{3, "Father"},
			{4, "Mother"},
			{5, "Brother"},
			{6, "Sister"},
		}

		for _, col := range relationalCols {
			val := row[col.index]
			if val == "null" || val == "" {
				continue
			}
			fName, lName := splitName(val)
			p.Relatives = append(p.Relatives, Relative{
				FirstName:    fName,
				LastName:     lName,
				Relationship: col.label,
			})
		}

		result = append(result, p)
	}

	jsonOut, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOut))
}