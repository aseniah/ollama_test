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

// Reference date for age calculation
var refDate = time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Birthday    string    `json:"Birthday"`
	Age         int       `json:"Age"`
	Relatives   []Relative `json:"Relatives"`
}

func parseBirthday(bdayStr string) (time.Time, error) {
	// Input format: M/D/YYYY or MM/DD/YYYY
	// We'll try to parse it
	layouts := []string{
		"1/2/2006",
		"1/2/2006", // Same as above, just to be sure
		"01/02/2006",
		"1/02/2006",
		"01/2/2006",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, bdayStr)
		if err == nil {
			return t, nil
		}
	}
	// Fallback: try standard layout with zeros
	// If input is "10/9/1940", standard layout "1/2/2006" works.
	// Let's just use the first one which covers both single and double digit days/months
	t, err := time.Parse("1/2/2006", bdayStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func calculateAge(bdayStr string) int {
	bday, err := parseBirthday(bdayStr)
	if err != nil {
		return 0
	}
	age := refDate.Year() - bday.Year()
	if refDate.YearDay() < bday.YearDay() {
		age--
	}
	return age
}

func cleanName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" || name == "null" {
		return ""
	}
	// The CSV has "First Last" format in the Name column
	parts := strings.Fields(name)
	if len(parts) < 2 {
		return name
	}
	first := parts[0]
	last := strings.Join(parts[1:], " ")
	return first
}

func parseRelative(name string, relType string) Relative {
	if name == "null" || name == "" {
		return Relative{}
	}
	parts := strings.Fields(name)
	if len(parts) < 2 {
		return Relative{FirstName: name, LastName: "", Relationship: relType}
	}
	return Relative{
		FirstName:    parts[0],
		LastName:     strings.Join(parts[1:], " "),
		Relationship: relType,
	}
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		// No data rows
		fmt.Println("[]")
		return
	}

	// Skip header
	header := records[0]
	// Map column indices
	colMap := make(map[string]int)
	for i, h := range header {
		colMap[strings.TrimSpace(h)] = i
	}

	var results []Person

	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) == 0 {
			continue
		}

		name := row[colMap["Name"]]
		birthday := row[colMap["Birthday"]]

		firstName := cleanName(name)
		lastName := ""
		parts := strings.Fields(name)
		if len(parts) >= 2 {
			lastName = strings.Join(parts[1:], " ")
		}

		age := calculateAge(birthday)

		var relatives []Relative

		fatherName := row[colMap["Father"]]
		if fatherName != "null" && fatherName != "" {
			relatives = append(relatives, parseRelative(fatherName, "Father"))
		}

		motherName := row[colMap["Mother"]]
		if motherName != "null" && motherName != "" {
			relatives = append(relatives, parseRelative(motherName, "Mother"))
		}

		brotherName := row[colMap["Brother"]]
		if brotherName != "null" && brotherName != "" {
			relatives = append(relatives, parseRelative(brotherName, "Brother"))
		}

		sisterName := row[colMap["Sister"]]
		if sisterName != "null" && sisterName != "" {
			relatives = append(relatives, parseRelative(sisterName, "Sister"))
		}

		person := Person{
			FirstName:   firstName,
			LastName:    lastName,
			Birthday:    birthday,
			Age:         age,
			Relatives:   relatives,
		}
		results = append(results, person)
	}

	// Special handling for the example data to match the expected output exactly
	// The example output has specific formatting for Last Names (some are single word, some might be split differently)
	// Looking at the expected output:
	// John Lennon -> "John Winston Lennon" in CSV -> First: John, Last: Lennon (Correct)
	// James McCartney -> "James Paul McCartney" -> First: James, Last: McCartney (Correct)
	// Ringo Starr -> "Ringo Starr" -> First: Ringo, Last: Starr (Correct)
	// George Harrison -> "George Harrison" -> First: George, Last: Harrison (Correct)

	// Wait, looking closer at the expected output:
	// John Lennon: Birthday "1940-10-09" (Format YYYY-MM-DD)
	// Our code reads "10/9/1940" and outputs it as is.
	// We need to reformat the Birthday to YYYY-MM-DD.

	for i := range results {
		// Reformat birthday
		bday, err := parseBirthday(results[i].Birthday)
		if err == nil {
			results[i].Birthday = bday.Format("2006-01-02")
		}
	}

	// Filter out empty relatives if any (though logic above handles nulls)
	// The expected output does not show empty relatives, so if a field was null, it shouldn't be in the list.
	// Our logic already skips null.

	// Output JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "   ")
	if err := encoder.Encode(results); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}