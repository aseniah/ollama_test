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
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Birthday  string   `json:"Birthday"`
	Age       int      `json:"Age"`
	Relatives []Relative `json:"Relatives"`
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read CSV: %v\n", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Fprintln(os.Stdout, "[]")
		return
	}

	var people []Person
	calcDate := time.Date(2025, time.July, 1, 0, 0, 0, 0, time.UTC)

	for _, record := range records[1:] {
		if len(record) < 6 {
			continue
		}

		nameStr := record[0]
		birthdayStrRaw := record[1]
		brotherName := strings.TrimSpace(record[5])
		sisterName := strings.TrimSpace(record[6])

		first, last := splitName(nameStr)
		
		// Parse Birthday (MM/DD/YYYY)
		var bday time.Time
		if err := time.NewCSVDate("01/02/2006").Scan(birthdayStrRaw); err != nil {
			// Fallback or error handling if parsing fails, but based on input it should be MM/DD/YYYY
			// Let's try parsing manually to be safe with standard library quirks
			var month, day, year int
			fmt.Sscanf(birthdayStrRaw, "%d/%d/%d", &month, &day, &year)
			bday = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		} else {
			bday = time.Time{} // Placeholder if Scan worked, but time.NewCSVDate is for formatting usually. 
			// Let's stick to explicit parsing for robustness given the input format "MM/DD/YYYY"
			var m, d, y int
			fmt.Sscanf(birthdayStrRaw, "%d/%d/%d", &m, &d, &y)
			bday = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		}

		birthdayISO := bday.Format("2006-01-02")

		// Calculate Age
		age := calcDate.Year() - bday.Year()
		// Check if birthday has passed in the current reference year (2025)
		if calcDate.Day() < bday.Day() || (calcDate.Day() == bday.Day() && calcDate.Month() < bday.Month()) {
			age--
		}

		var relatives []Relative
		
		// Add Father
		fatherStr := strings.TrimSpace(record[3])
		if fatherStr != "" && fatherStr != "null" {
			fF, fL := splitName(fatherStr)
			relatives = append(relatives, Relative{fF, fL, "Father"})
		}

		// Add Mother
		motherStr := strings.TrimSpace(record[4])
		if motherStr != "" && motherStr != "null" {
			mF, mL := splitName(motherStr)
			relatives = append(relatives, Relative{mF, mL, "Mother"})
		}

		// Add Brother
		if brotherName != "" && brotherName != "null" {
			brF, brL := splitName(brotherName)
			relatives = append(relatives, Relative{brF, brL, "Brother"})
		}

		// Add Sister
		sisterStr := strings.TrimSpace(record[5]) // Wait, index logic check: Name(0), Bday(1), Died(2), Father(3), Mother(4), Brother(5), Sister(6)
		// Correction: The input CSV header says: Name,Birthday,Died,Father,Mother,Brother,Sister
		// Indices: 0, 1, 2, 3, 4, 5, 6. 
		// My previous loop used record[5] for brother and [6] for sister? 
		// Let's re-verify the row structure based on header provided in prompt.
		// Row: John Winston Lennon,10/9/1940,12/8/1980,Alfred Lennon,Julia Stanley,null,null
		// 0: Name, 1: Bday, 2: Died, 3: Father, 4: Mother, 5: Brother, 6: Sister.
		
		// Let's re-extract brother/sister with correct indices.
		brotherStr := strings.TrimSpace(record[5])
		sisterStr = strings.TrimSpace(record[6]) // This is inside the loop now

		if sisterStr != "" && sisterStr != "null" {
			sF, sL := splitName(sisterStr)
			relatives = append(relatives, Relative{sF, sL, "Sister"})
		} else if brotherName != "" && brotherName != "null" {
             // The logic above for brother was slightly duplicated in my thought process. 
             // Let's ensure we only add non-null/non-empty ones.
        }

		// Clean up the relatives slice to avoid duplicates from my logic above if I re-added them incorrectly.
		// Re-doing the relative construction cleanly:
		relatives = []Relative{}
		
		fatherStr = strings.TrimSpace(record[3])
		if fatherStr != "" && fatherStr != "null" {
			ff, fl := splitName(fatherStr)
			relatives = append(relatives, Relative{ff, fl, "Father"})
		}

		motherStr = strings.TrimSpace(record[4])
		if motherStr != "" && motherStr != "null" {
			mf, ml := splitName(motherStr)
			relatives = append(relatives, Relative{mf, ml, "Mother"})
		}

		brotherStr = strings.TrimSpace(record[5])
		if brotherStr != "" && brotherStr != "null" {
			bf, bl := splitName(brotherStr)
			relatives = append(relatives, Relative{bf, bl, "Brother"})
		}

		sisterStr = strings.TrimSpace(record[6])
		if sisterStr != "" && sisterStr != "null" {
			sf, sl := splitName(sisterStr)
			relatives = append(relatives, Relative{sf, sl, "Sister"})
		}

		person := Person{
			FirstName: first,
			LastName:  last,
			Birthday:  birthdayISO,
			Age:       age,
			Relatives: relatives,
		}

		people = append(people, person)
	}

	output, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func splitName(name string) (string, string) {
	name = strings.TrimSpace(name)
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	first := parts[0]
	last := ""
	if len(parts) > 1 {
		last = strings.Join(parts[1:], " ")
	}
	return first, last
}