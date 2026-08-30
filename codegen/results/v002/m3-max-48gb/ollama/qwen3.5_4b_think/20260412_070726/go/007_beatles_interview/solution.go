package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Birthday   string `json:"Birthday"`
	Age        int    `json:"Age"`
	Relatives  []Relate `json:"Relatives"`
}

type Relate struct {
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type PersonData struct {
	Name      string
	Birthday  string
	Died      string
	Father    string
	Mother    string
	Brother   string
	Sister    string
}

func calculateAge(birthDate time.Time, refDate time.Time) int {
	// If birthday has not occurred in the reference year yet, age is one less
	age := refDate.Year() - birthDate.Year()
	if refDate.Month() < birthDate.Month() || 
		(refDate.Month() == birthDate.Month() && refDate.Day() < birthDate.Day()) {
		age--
	}
	return age
}

func main() {
	// Read input CSV
	csvContent, err := os.ReadFile("input/input.csv")
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Parse the CSV
	lines := strings.Split(string(csvContent), "\n")
	personData := make([]PersonData, 0)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i == 0 { // Skip header
			continue
		}
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			continue
		}

		// Parse fields
		name := fields[0]
		birthdayStr := fields[1]
		diedStr := fields[2]
		fatherStr := fields[3]
		motherStr := fields[4]
		brotherStr := fields[5]
		sisterStr := fields[6]

		// Split name into first and last name
		nameParts := strings.Fields(name)
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			lastName = strings.Join(nameParts[1:], " ")
		}

		// Parse birthday: MM/DD/YYYY
		parts := strings.Split(birthdayStr, "/")
		if len(parts) == 3 {
			birthMonth, _ := strconv.Atoi(parts[0])
			birthDay, _ := strconv.Atoi(parts[1])
			birthYear, _ := strconv.Atoi(parts[2])
			birthDate := time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)
		} else {
			birthDate = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
		}

		// Reference date: July 1, 2025
		refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

		personData = append(personData, PersonData{
			Name:      name,
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  strings.Join(parts, "-"),
			Died:      diedStr,
			Father:    fatherStr,
			Mother:    motherStr,
			Brother:   brotherStr,
			Sister:    sisterStr,
		})
	}

	// Calculate age for each person
	for i := range personData {
		pd := &personData[i]
		if pd.Birthday == "" || pd.Birthday == "null" {
			continue
		}

		// Parse the birthday string (YYYY-MM-DD)
		if strings.Contains(pd.Birthday, "-") {
			year, month, day, err := strconv.Atoi(pd.Birthday[0:4])
			dayInt := day
			yearInt := year
			monthInt := month

			// For age calculation, use the original date
			birthDate := time.Date(year, time.Month(month), dayInt, 0, 0, 0, 0, time.UTC)
			pd.Age = calculateAge(birthDate, refDate)
		}
	}

	// Build output JSON
	var out []Person
	for _, pd := range personData {
		if pd.Birthday == "" {
			continue
		}

		person := Person{
			FirstName: pd.FirstName,
			LastName:  pd.LastName,
			Birthday:  pd.Birthday,
			Age:       pd.Age,
		}

		// Process relatives
		rels := []Relate{}

		// Father
		if pd.Father != "null" {
			nameParts := strings.Fields(pd.Father)
			rel := Relate{
				FirstName:  nameParts[0],
				LastName:   nameParts[1],
				Relationship: "Father",
			}
			rels = append(rels, rel)
		}

		// Mother
		if pd.Mother != "null" {
			nameParts := strings.Fields(pd.Mother)
			rel := Relate{
				FirstName:  nameParts[0],
				LastName:   nameParts[1],
				Relationship: "Mother",
			}
			rels = append(rels, rel)
		}

		// Brother
		if pd.Brother != "null" {
			nameParts := strings.Fields(pd.Brother)
			rel := Relate{
				FirstName:  nameParts[0],
				LastName:   nameParts[1],
				Relationship: "Brother",
			}
			rels = append(rels, rel)
		}

		// Sister
		if pd.Sister != "null" {
			nameParts := strings.Fields(pd.Sister)
			rel := Relate{
				FirstName:  nameParts[0],
				LastName:   nameParts[1],
				Relationship: "Sister",
			}
			rels = append(rels, rel)
		}

		person.Relatives = rels
		out = append(out, person)
	}

	// Output JSON
	jsonData, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}