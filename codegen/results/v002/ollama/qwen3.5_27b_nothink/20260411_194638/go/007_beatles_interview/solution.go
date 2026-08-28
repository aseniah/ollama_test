package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
		os.Exit(1)
	}

	// Skip header
	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	var people []Person
	calcDate, _ := time.Parse("2006-01-02", "2025-07-01")

	for _, record := range records[1:] {
		if len(record) < 7 {
			continue
		}

		// Parse Name
		nameParts := strings.SplitN(record[0], " ", 2)
		firstName := nameParts[0]
		lastName := ""
		if len(nameParts) > 1 {
			lastName = nameParts[1]
		}

		// Parse Birthday
		var bday time.Time
		// Try common formats: M/D/YYYY or MM/DD/YYYY
		for _, layout := range []string{"1/2/2006", "1/2/06", "1/2"} {
			if b, err := time.Parse(layout, record[1]); err == nil {
				bday = b
				break
			}
		}
		// Fallback: handle strictly MM/DD/YYYY if standard parsing failed but input looks valid
		if bday.IsZero() {
			// Manual parse for safety with single digit months/days
			parts := strings.Split(record[1], "/")
			if len(parts) == 3 {
				m, _ := strconv.Atoi(parts[0])
				d, _ := strconv.Atoi(parts[1])
				y, _ := strconv.Atoi(parts[2])
				bday = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
			}
		}

		bdayStr := bday.Format("2006-01-02")

		// Calculate Age as of July 1, 2025
		age := calcDate.Year() - bday.Year()
		// Check if birthday hasn't occurred yet in 2025
		bdayInCalcYear := time.Date(calcDate.Year(), bday.Month(), bday.Day(), 0, 0, 0, 0, time.UTC)
		if calcDate.Before(bdayInCalcYear) {
			age--
		}

		// Parse Relatives
		var relatives []Relative
		relData := []struct {
			val string
			rel string
		}{
			{record[3], "Father"},
			{record[4], "Mother"},
			{record[5], "Brother"},
			{record[6], "Sister"},
		}

		for _, rd := range relData {
			val := strings.TrimSpace(rd.val)
			if val == "null" || val == "" {
				continue
			}
			// Parse Relative Name (Assume First Last format)
			relNameParts := strings.SplitN(val, " ", 2)
			rFirstName := relNameParts[0]
			rLastName := ""
			if len(relNameParts) > 1 {
				rLastName = relNameParts[1]
			}

			relatives = append(relatives, Relative{
				FirstName:    rFirstName,
				LastName:     rLastName,
				Relationship: rd.rel,
			})
		}

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  bdayStr,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}