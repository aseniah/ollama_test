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

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	var people []Person
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 2 {
			continue
		}

		name := strings.TrimSpace(record[0])
		birthdayStr := strings.TrimSpace(record[1])

		nameParts := strings.Fields(name)
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthDate, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			continue
		}

		age := refDate.Year() - birthDate.Year()
		if refDate.Month() < birthDate.Month() || (refDate.Month() == birthDate.Month() && refDate.Day() < birthDate.Day()) {
			age--
		}

		var relatives []Relative
		relCols := []string{"Father", "Mother", "Brother", "Sister"}
		for j, col := range relCols {
			if j+2 < len(record) {
				val := strings.TrimSpace(record[j+2])
				if val != "null" && val != "" {
					parts := strings.Fields(val)
					if len(parts) >= 2 {
						relatives = append(relatives, Relative{
							FirstName:    parts[0],
							LastName:     parts[len(parts)-1],
							Relationship: col,
						})
					}
				}
			}
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthDate.Format("2006-01-02"),
			Age:       age,
			Relatives: relatives,
		})
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}