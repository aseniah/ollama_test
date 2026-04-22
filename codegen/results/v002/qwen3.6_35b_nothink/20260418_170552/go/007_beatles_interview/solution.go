package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func calculateAge(birthdayStr string, refDate time.Time) int {
	// Parse birthday in format MM/DD/YYYY
	t, err := time.Parse("1/2/2006", birthdayStr)
	if err != nil {
		return 0
	}

	age := refDate.Year() - t.Year()
	// Check if birthday has occurred this year
	if refDate.YearDay() < t.YearDay() {
		age--
	}
	return age
}

func main() {
	// Read CSV file
	data, err := ioutil.ReadFile("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
		return
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CSV: %v\n", err)
		return
	}

	// Skip header
	records = records[1:]

	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for _, record := range records {
		if len(record) < 7 {
			continue
		}

		nameParts := strings.Fields(record[0])
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]
		birthday := record[1]
		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		// Calculate age
		age := calculateAge(birthday, refDate)

		// Format birthday as YYYY-MM-DD
		t, err := time.Parse("1/2/2006", birthday)
		birthdayFormatted := ""
		if err == nil {
			birthdayFormatted = t.Format("2006-01-02")
		}

		var relatives []Relative

		if father != "null" && father != "" {
			fatherParts := strings.Fields(father)
			if len(fatherParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    fatherParts[0],
					LastName:     fatherParts[len(fatherParts)-1],
					Relationship: "Father",
				})
			}
		}

		if mother != "null" && mother != "" {
			motherParts := strings.Fields(mother)
			if len(motherParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    motherParts[0],
					LastName:     motherParts[len(motherParts)-1],
					Relationship: "Mother",
				})
			}
		}

		if brother != "null" && brother != "" {
			brotherParts := strings.Fields(brother)
			if len(brotherParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    brotherParts[0],
					LastName:     brotherParts[len(brotherParts)-1],
					Relationship: "Brother",
				})
			}
		}

		if sister != "null" && sister != "" {
			sisterParts := strings.Fields(sister)
			if len(sisterParts) >= 2 {
				relatives = append(relatives, Relative{
					FirstName:    sisterParts[0],
					LastName:     sisterParts[len(sisterParts)-1],
					Relationship: "Sister",
				})
			}
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}