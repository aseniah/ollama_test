package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	f, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(rows) < 2 {
		fmt.Println("[]")
		return
	}

	header := rows[0]
	var nameIdx, birthdayIdx, diedIdx, fatherIdx, motherIdx, brotherIdx, sisterIdx int
	for i, col := range header {
		switch col {
		case "Name":
			nameIdx = i
		case "Birthday":
			birthdayIdx = i
		case "Died":
			diedIdx = i
		case "Father":
			fatherIdx = i
		case "Mother":
			motherIdx = i
		case "Brother":
			brotherIdx = i
		case "Sister":
			sisterIdx = i
		}
	}

	// Reference date: July 1, 2025
	refYear := 2025
	refMonth := 7
	refDay := 1

	var result []Person

	for _, row := range rows[1:] {
		if len(row) <= nameIdx {
			continue
		}

		fullName := strings.TrimSpace(row[nameIdx])
		nameParts := strings.Fields(fullName)
		var firstName, lastName string
		if len(nameParts) >= 2 {
			// Handle middle names: first part is first name, last part is last name
			firstName = nameParts[0]
			lastName = nameParts[len(nameParts)-1]
		} else if len(nameParts) == 1 {
			firstName = nameParts[0]
			lastName = ""
		}

		birthdayStr := strings.TrimSpace(row[birthdayIdx])
		birthdayDate := parseDate(birthdayStr)

		// Age as of July 1, 2025
		age := refYear - birthdayDate.Year
		// Check if birthday hasn't occurred yet this year
		if birthdayDate.Month > refMonth || (birthdayDate.Month == refMonth && birthdayDate.Day > refDay) {
			age--
		}

		// Format birthday as YYYY-MM-DD
		birthdayFormatted := fmt.Sprintf("%04d-%02d-%02d", birthdayDate.Year, birthdayDate.Month, birthdayDate.Day)

		var relatives []Relative

		fatherStr := getCell(row, fatherIdx)
		if fatherStr != "null" && fatherStr != "" {
			fn, ln := splitName(fatherStr)
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Father"})
		}

		motherStr := getCell(row, motherIdx)
		if motherStr != "null" && motherStr != "" {
			fn, ln := splitName(motherStr)
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Mother"})
		}

		brotherStr := getCell(row, brotherIdx)
		if brotherStr != "null" && brotherStr != "" {
			fn, ln := splitName(brotherStr)
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Brother"})
		}

		sisterStr := getCell(row, sisterIdx)
		if sisterStr != "null" && sisterStr != "" {
			fn, ln := splitName(sisterStr)
			relatives = append(relatives, Relative{FirstName: fn, LastName: ln, Relationship: "Sister"})
		}

		if relatives == nil {
			relatives = []Relative{}
		}

		result = append(result, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  birthdayFormatted,
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

type date struct {
	Year  int
	Month int
	Day   int
}

func parseDate(s string) date {
	s = strings.TrimSpace(s)
	// Expected format: M/D/YYYY or MM/DD/YYYY
	parts := strings.Split(s, "/")
	if len(parts) != 3 {
		return date{}
	}
	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])
	return date{Year: year, Month: month, Day: day}
}

func getCell(row []string, idx int) string {
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[idx])
}

func splitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) >= 2 {
		firstName := parts[0]
		lastName := parts[len(parts)-1]
		return firstName, lastName
	} else if len(parts) == 1 {
		return parts[0], ""
	}
	return "", ""
}