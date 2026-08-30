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
	FirstName   string     `json:"FirstName"`
	LastName    string     `json:"LastName"`
	Birthday    string     `json:"Birthday"`
	Age         int        `json:"Age"`
	Relatives   []Relative `json:"Relatives"`
}

func parseBirthday(birthdayStr string) (time.Time, error) {
	// Format: M/D/YYYY
	return time.Parse("1/2/2006", birthdayStr)
}

func calculateAge(birthday time.Time, referenceDate time.Time) int {
	age := referenceDate.Year() - birthday.Year()

	// Check if birthday hasn't occurred yet in the reference year
	if referenceDate.Month() < birthday.Month() ||
		(referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
		age--
	}

	return age
}

func main() {
	// Read CSV file
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

	// Reference date: July 1, 2025
	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	// Skip header row
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 7 {
			continue
		}

		name := record[0]
		birthdayStr := record[1]
		father := record[3]
		mother := record[4]
		brother := record[5]
		sister := record[6]

		// Parse name - first name is first word, last name is last word
		nameParts := strings.Fields(name)
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		// Parse birthday
		birthday, err := parseBirthday(birthdayStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
			continue
		}

		// Calculate age
		age := calculateAge(birthday, referenceDate)

		// Build relatives list
		var relatives []Relative

		if father != "null" && strings.TrimSpace(father) != "" {
			fatherParts := strings.Fields(father)
			relatives = append(relatives, Relative{
				FirstName:    fatherParts[0],
				LastName:     fatherParts[len(fatherParts)-1],
				Relationship: "Father",
			})
		}

		if mother != "null" && strings.TrimSpace(mother) != "" {
			motherParts := strings.Fields(mother)
			relatives = append(relatives, Relative{
				FirstName:    motherParts[0],
				LastName:     motherParts[len(motherParts)-1],
				Relationship: "Mother",
			})
		}

		if brother != "null" && strings.TrimSpace(brother) != "" {
			brotherParts := strings.Fields(brother)
			relatives = append(relatives, Relative{
				FirstName:    brotherParts[0],
				LastName:     brotherParts[len(brotherParts)-1],
				Relationship: "Brother",
			})
		}

		if sister != "null" && strings.TrimSpace(sister) != "" {
			sisterParts := strings.Fields(sister)
			relatives = append(relatives, Relative{
				FirstName:    sisterParts[0],
				LastName:     sisterParts[len(sisterParts)-1],
				Relationship: "Sister",
			})
		}

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := birthday.Format("2006-01-02")

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	// Output JSON
	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}