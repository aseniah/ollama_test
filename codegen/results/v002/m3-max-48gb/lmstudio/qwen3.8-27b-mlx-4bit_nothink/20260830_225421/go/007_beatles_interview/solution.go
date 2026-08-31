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
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Println("[]")
		return
	}

	// Parse header to find column indices
	header := records[0]
	colIdx := make(map[string]int)
	for i, h := range header {
		colIdx[strings.TrimSpace(h)] = i
	}

	getCol := func(row []string, name string) string {
		if idx, ok := colIdx[name]; ok && idx < len(row) {
			return strings.TrimSpace(row[idx])
		}
		return ""
	}

	// Helper to parse date from M/D/YYYY or MM/DD/YYYY
	parseDate := func(dateStr string) (int, int, int, bool) {
		if dateStr == "" || strings.ToLower(dateStr) == "null" {
			return 0, 0, 0, false
		}
		parts := strings.Split(dateStr, "/")
		if len(parts) != 3 {
			return 0, 0, 0, false
		}
		m, err1 := strconv.Atoi(parts[0])
		d, err2 := strconv.Atoi(parts[1])
		y, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return 0, 0, 0, false
		}
		return m, d, y, true
	}

	// Calculate age as of a given date (month, day, year)
	calcAge := func(birthM, birthD, birthY, refM, refD, refY int) int {
		age := refY - birthY
		// If birthday hasn't occurred yet this year, subtract 1
		if refM < birthM || (refM == birthM && refD < birthD) {
			age--
		}
		return age
	}

	// Format date as YYYY-MM-DD
	formatDate := func(m, d, y int) string {
		return fmt.Sprintf("%04d-%02d-%02d", y, m, d)
	}

	people := make([]Person, 0)

	for _, row := range records[1:] {
		if len(row) == 0 || strings.TrimSpace(row[0]) == "" {
			continue
		}

		name := strings.TrimSpace(getCol(row, "Name"))
		nameParts := strings.Fields(name)
		if len(nameParts) < 2 {
			continue
		}
		firstName := nameParts[0]
		lastName := nameParts[len(nameParts)-1]

		birthdayStr := getCol(row, "Birthday")
		diedStr := getCol(row, "Died")

		bM, bD, bY, ok := parseDate(birthdayStr)
		if !ok {
			continue
		}

		// Determine reference date
		var refM, refD, refY int
		if strings.ToLower(diedStr) != "null" && diedStr != "" {
			// Person died - use death date
			dM, dD, dY, dOk := parseDate(diedStr)
			if dOk {
				refM, refD, refY = dM, dD, dY
			} else {
				refM, refD, refY = 7, 1, 2025
			}
		} else {
			refM, refD, refY = 7, 1, 2025
		}

		age := calcAge(bM, bD, bY, refM, refD, refY)

		relatives := make([]Relative, 0)

		addRelative := func(nameStr, relationship string) {
			if strings.ToLower(nameStr) == "null" || nameStr == "" {
				return
			}
			parts := strings.Fields(nameStr)
			if len(parts) < 2 {
				return
			}
			rFirst := parts[0]
			rLast := parts[len(parts)-1]
			relatives = append(relatives, Relative{
				FirstName:    rFirst,
				LastName:     rLast,
				Relationship: relationship,
			})
		}

		addRelative(getCol(row, "Father"), "Father")
		addRelative(getCol(row, "Mother"), "Mother")
		addRelative(getCol(row, "Brother"), "Brother")
		addRelative(getCol(row, "Sister"), "Sister")

		people = append(people, Person{
			FirstName: firstName,
			LastName:  lastName,
			Birthday:  formatDate(bM, bD, bY),
			Age:       age,
			Relatives: relatives,
		})
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}