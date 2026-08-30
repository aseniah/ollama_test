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
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Birthday  string   `json:"Birthday"`
	Age       int      `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func parseName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	first := parts[0]
	last := "unknown"
	if len(parts) > 1 {
		last = parts[len(parts)-1] // Assuming last name is the last word if multiple
	}
	return first, last
}

func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" || dateStr == "null" {
		return time.Time{}, fmt.Errorf("empty or null date")
	}

	// Try common formats: M/D/YYYY, MM/DD/YYYY, etc.
	formats := []string{
		"1/2/2006",
		"1/2/06",
		"1/02/2006",
		"1/02/06",
	}

	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t, nil
		}
	}

	// Try to manually parse if standard formats fail (though CSV usually consistent)
	// Split by '/'
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month: %v", err)
	}
	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %v", err)
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year: %v", err)
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func calculateAge(birthDate time.Time, referenceDate time.Time) int {
	years := referenceDate.Year() - birthDate.Year()

	// Check if birthday has occurred in the reference year yet
	birthdayInRefYear := birthDate.Month() == referenceDate.Month() && birthDate.Day() <= referenceDate.Day()

	if !birthdayInRefYear || (referenceDate.Month() < birthDate.Month()) {
		years--
	} else if referenceDate.Month() == birthDate.Month() && referenceDate.Day() < birthDate.Day() {
		years--
	}

	return years
}

func formatOutputDate(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format("2006-01-02")
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

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	// Skip header
	dataRecords := records[1:]
	var people []Person

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)

	for _, record := range dataRecords {
		if len(record) < 6 {
			continue // Skip malformed rows
		}

		fullName := strings.TrimSpace(record[0])
		bdayStr := strings.TrimSpace(record[1])
		diedStr := strings.TrimSpace(record[2])
		fatherStr := strings.TrimSpace(record[3])
		motherStr := strings.TrimSpace(record[4])
		brotherStr := strings.TrimSpace(record[5])
		var sisterStr string
		if len(record) > 6 {
			sisterStr = strings.TrimSpace(record[6])
		}

		first, last := parseName(fullName)

		bday, err := parseDate(bdayStr)
		if err != nil {
			continue // Skip if birthday cannot be parsed
		}

		// If the person is dead, we might need to cap age? 
		// The prompt says "Calculate ages as of July 1, 2025". 
		// Usually, for deceased people, age is calculated at death. 
		// HOWEVER, looking at the expected output:
		// John Lennon (died 1980) -> Age 40. (2025 - 1940 = 85). But output says 40.
		// George Harrison (died 2001) -> Age 58. (2025 - 1943 = 82). But output says 58.
		// Alive people:
		// McCartney (b 1942): 2025-1942 = 83. Output says 83.
		// Starr (b 1940): 2025-1940 = 85. Wait, July 1 2025 vs July 7 1940. 
		//   Born July 7, 2025 is before July 7? No. July 1 is before July 7.
		//   So 2025 - 1940 = 85. But he hasn't had birthday yet in 2025. So 84.
		//   Output says 84. Correct.
		//
		// Conclusion: 
		// If "Died" is not null, calculate age based on Death Date vs Birth Date?
		// OR calculate age at death date relative to today? No.
		// Let's re-verify Lennon: Born 10/9/1940. Died 12/8/1980.
		// Age at death: 1980 - 1940 = 40. (Oct is before Dec, so he had birthday). Correct.
		// Harrison: Born 2/25/1943. Died 11/29/2001.
		// Age at death: 2001 - 1943 = 58. (Feb is before Nov, so he had birthday). Correct.
		//
		// Rule: If dead, calculate age as of Death Date. If alive, calculate age as of July 1, 2025.

		var finalAge int
		if diedStr != "" && diedStr != "null" {
			diedDate, err := parseDate(diedStr)
			if err == nil {
				finalAge = calculateAge(bday, diedDate)
			} else {
				// If death date fails, fall back to reference date? 
				// But data seems clean. Let's assume valid if string exists.
				finalAge = calculateAge(bday, referenceDate)
			}
		} else {
			finalAge = calculateAge(bday, referenceDate)
		}

		var relatives []Relative

		processRel := func(name string, relType string) {
			if name != "" && name != "null" {
				fn, ln := parseName(name)
				relatives = append(relatives, Relative{
					FirstName:    fn,
					LastName:     ln,
					Relationship: relType,
				})
			}
		}

		processRel(fatherStr, "Father")
		processRel(motherStr, "Mother")
		processRel(brotherStr, "Brother")
		processRel(sisterStr, "Sister")

		person := Person{
			FirstName: first,
			LastName:  last,
			Birthday:  formatOutputDate(bday),
			Age:       finalAge,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}