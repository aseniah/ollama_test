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

type Relative struct {
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Relationship string `json:"Relationship"`
}

type Person struct {
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
	Birthday   string     `json:"Birthday"`
	Age        int        `json:"Age"`
	Relatives  []Relative `json:"Relatives"`
}

func parseName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	// The last part is the LastName, everything before is FirstName (including middle names)
	// However, the expected output shows "John" for "John Winston Lennon".
	// Looking at the expected output: 
	// "John Winston Lennon" -> FirstName: "John", LastName: "Lennon"
	// "James Paul McCartney" -> FirstName: "James", LastName: "McCartney"
	// "Ringo Starr" -> FirstName: "Ringo", LastName: "Starr"
	// "George Harrison" -> FirstName: "George", LastName: "Harrison"
	// It seems we take the first word as FirstName and the last word as LastName.
	return parts[0], parts[len(parts)-1]
}

func parseDate(dateStr string) (time.Time, error) {
	// Input format in CSV: 10/9/1940 (M/D/YYYY)
	return time.Parse("1/2/2006", dateStr)
}

func calculateAge(birthday time.Time, referenceDate time.Time) int {
	age := referenceDate.Year() - birthday.Year()
	if referenceDate.Month() < birthday.Month() || (referenceDate.Month() == birthday.Month() && referenceDate.Day() < birthday.Day()) {
		age--
	}
	return age
}

func parseRelative(name string, relationship string) *Relative {
	if name == "null" || name == "" {
		return nil
	}
	f, l := parseName(name)
	return &Relative{
		FirstName:    f,
		LastName:     l,
		Relationship: relationship,
	}
	// Wait, looking at the CSV vs JSON:
	// "Alfred Lennon" -> FirstName: "Alfred", LastName: "Lennon"
	// "Julia Stanley" -> FirstName: "Julia", LastName: "Stanley"
	// For relatives, the input is "FirstName LastName".
	// Let's check "Mike McGear" -> FirstName: "Mike", LastName: "McGear"
	// The logic is: split by space. First word is FirstName, rest is LastName.
	// Actually, most are just two words. Let's use a simpler split.
}

func splitRelativeName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[len(parts)-1]
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return
	}

	// Map header to indices
	cols := make(map[string]int)
	for i, h := range header {
		cols[h] = i
	}

	referenceDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	var results []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		// Parse Main Person
		fullName := record[cols["Name"]]
		bdayStr := record[cols["Birthday"]]
		bday, _ := time.Parse("1/2/2006", bdayStr)
		
		pFirstName, pLastName := splitRelativeName(fullName)
		// Special check: for the main person, the expected format keeps 
		// the first name as the first word and last name as the last word.
		// Looking at John Winston Lennon -> FirstName: John, LastName: Lennon.
		
		age := calculateAge(bday, referenceDate)

		person := Person{
			FirstName: pFirstName,
			LastName:  pLastName,
			Birthday:  bday.Format("2006-01-02"),
			Age:       age,
			Relatives: []Relative{},
		}

		// Parse Relatives
		relTypes := []string{"Father", "Mother", "Brother", "Sister"}
		for _, relType := range relTypes {
			relName := record[cols[relType]]
			if relName != "null" && relName != "" {
				rFirst, rLast := splitRelativeName(relName)
				person.Relatives = append(processRelative(relName, relType, person.Relatives))
			}
		}
		
		// Re-implementing relative logic to match exact expected output structure
		person.Relatives = []Relative{}
		relMappings := []struct {
			col  string
			rel  string
		}{
			{"Father", "Father"},
			{"Mother", "Mother"},
			{"Brother", "Brother"},
			{"Sister", "Sister"},
		}

		for _, mapping := range relMappings {
			val := record[cols[mapping.col]]
			if val != "null" && val != "" {
				parts := strings.Fields(val)
				f, l := "", ""
				if len(parts) > 0 {
					f = parts[0]
					if len(parts) > 1 {
						l = parts[len(parts)-1]
					}
				}
				person.Relatives = append(person.Relatives, Relative{
					FirstName:    f,
					LastName:     l,
					Relationship: mapping.rel,
				})
			}
		}

		results = append(results, person)
	}

	output, _ := json.MarshalIndent(results, "", "  ")
	fmt.Println(string(output))
}

// Helper to handle the relative logic properly within the loop
func processRelative(name string, rel string, current []Relative) []Relative {
	// This was a placeholder for the logic moved into main
	return current
}

// Redoing the logic to be clean in one block.