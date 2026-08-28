package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Person represents the structure for the output JSON
type Person struct {
	FirstName  string    `json:"FirstName"`
	LastName   string    `json:"LastName"`
	Birthday   string    `json:"Birthday"`
	Age        int       `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

// Relative represents the structure for the relative entries
type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read header
	header, err := reader.Read()
	if err != nil {
		return
	}

	// Map header names to column indices
	colMap := make(map[string]int)
	for i, name := range header {
		colMap[name] = i
	}

	// Reference date for age calculation: July 1, 2025
	refDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// Parse Name into First and Last
		fullName := record[colMap["Name"]]
		nameParts := strings.Split(fullName, " ")
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			// The last element of the split name is treated as Last Name
			// In the context of the CSV, the "LastName" in JSON is the last part of the Full Name
			lastName = nameParts[len(nameParts)-1]
		}

		// Parse Birthday (Format: M/D/YYYY)
		birthdayStr := record[colMap["Birthday"]]
		birthday, err := time.Parse("1/2/2006", birthdayStr)
		if err != nil {
			// Fallback for single digit months/days if needed
			birthday, err = time.Parse("01/02/2006", birthdayStr)
			if err != nil {
				continue
			}
		}

		// Calculate Age
		age := refDate.Year() - birthday.Year()
		if refDate.YearDay() < birthday.YearDay() {
			age--
		}

		person := Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthday.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Parse Relatives
		// Columns: Father, Mother, Brother, Sister
		relTypes := []struct {
			col  string
			rel  string
		}{
			{"Father", "Father"},
			{"Mother", "Mother"},
			{"Brother", "Brother"},
			{"Sister", "Sister"},
		}

		for _, rt := range relTypes {
			val := record[colMap[rt.col]]
			if val != "null" && val != "" {
				parts := strings.Split(val, " ")
				rFirst := parts[0]
				rLast := ""
				if len(parts) > 1 {
					rLast = parts[len(parts)-1]
				}
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    rFirst,
					LastName:     rLast,
					Relationship: rt.rel,
				})
			}
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(output))
}