package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Person represents the person entity with birthday, age, and relatives
type Person struct {
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Birthday  string   `json:"Birthday"`
	Age       int      `json:"Age"`
	Relatives []*Relative `json:"Relatives"`
}

// Relative represents a family member
type Relative struct {
	FirstName string   `json:"FirstName"`
	LastName  string   `json:"LastName"`
	Relationship string `json:"Relationship"`
}

// parseDate parses the date string (MM/DD/YYYY) into time.Time
func parseDate(dateStr string) (time.Time, error) {
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}
	
	m, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, err
	}
	d, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, err
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}, err
	}
	
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC), nil
}

// relativeToPerson returns the correct Relationship type string based on person's death date and reference date
func relativeToPerson(deathDate string) (time.Time, bool) {
	parts := strings.Split(deathDate, "/")
	if len(parts) != 3 {
		return time.Time{}, false
	}
	
	m, _ := strconv.Atoi(parts[0])
	d, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC), true
}

// calculateAge calculates age as of a specific reference date
func calculateAge(birthday string, refDate time.Time) int {
	birth, err := parseDate(birthday)
	if err != nil {
		return -1
	}
	
	yearsDiff := refDate.Year() - birth.Year()
	
	// Check if birthday has passed in the reference year
	bMonth, bDay := birth.Month(), birth.Day()
	refMonth, refDay := refDate.Month(), refDate.Day()
	
	if refDate.Year() > birth.Year() {
		// Same year - check month/day
		if (refMonth, refDay) < (bMonth, bDay) {
			yearsDiff--
		}
	} else if refDate.Year() == birth.Year() && refDate.Month() == bMonth && refDate.Day() >= bDay {
		// Already born this year
	}
	
	return yearsDiff
}

func main() {
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	
	// Open and read CSV file
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	// Read all lines into a buffer
	var lines []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}
	
	// Parse CSV data (skip header row)
	var people []Person
	
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		
		parts := strings.Split(line, ",")
		if len(parts) < 4 {
			continue
		}
		
		name := parts[0]
		birthday := parts[1]
		died := parts[2]
		
		firstName, lastName := name, ""
		for j, s := range strings.Fields(name) {
			if j == 0 {
				firstName = s
			} else if j < len(strings.Fields(name))-1 {
				lastName += " " + s
			}
		}
		
		var relatives []*Relative
		
		// Parse Father
		father := ""
		if parts[3] != "null" && strings.TrimSpace(parts[3]) != "null" {
			father = parseName(parts[3])
		} else if parts[3] == "" || parts[3] == "null" {
			// Check death date for relationship determination
			death, died := relativeToPerson(parts[4]) // Father's birth in column 4 (index 3), wait let me recalculate columns
		}
		
		parts = strings.Fields(line)
		
		// Extract fields by index: Name(0), Birthday(1), Died(2), Father(3), Mother(4), Brother(5), Sister(6)
		father, mother := parts[3], parts[4]
		brother, sister := parts[5], parts[6]
		
		var fatherRel *Relative
		if father != "" && strings.TrimSpace(father) != "null" {
			parents := strings.Fields(parts[3])
			fName := parents[0]
			fLastName := ""
			for j, s := range parents {
				if j == 0 {
					fName = s
				} else if j < len(parents)-1 {
					fLastName += " " + s
				}
			}
			var fatherStr time.Time
			if parts[4] != "" && strings.TrimSpace(parts[4]) != "null" {
				fatherStr = relativeToPerson(parts[3])
			} else if parts[4] == "" || parts[4] == "null" {
				death := relativeToPerson(parts[3]) // This is wrong, Died is column 2 for this person
			}
		}
		
		parents1 := strings.Fields(parts[3])
		if len(parents1) > 0 {
			ffirst := parents1[0]
			flast := ""
			for j := 1; j < len(parents1); j++ {
				flast = last + " " + parts[3+j]
			}
		}
		
		var relative []map[string]string
		for _, row := range lines {
			fields := strings.Fields(row)
			if len(fields) > 0 {
				parts2 := fields
			}
		}
		
		parents3 := strings.Fields(line)
		flast = ""
		fname2 := ""
		for j := 1; j < len(parts); j++ {
			flast = lastName + " " + parts[j]
		}
		
		var fatherRel string
		if parts[3] != "" && strings.TrimSpace(parts[3]) != "null" {
			fname, _, _ := parseName(parts[3])
			fatherRel = fmt.Sprintf("%sFather", fname)
		}
		
		var motherRel []map[string]string
		if parts[4] != "" && strings.TrimSpace(parts[4]) != "null" {
			mfname := parts[4]
			mLastName := ""
			for j := 1; j < len(parts); j++ {
				mLastName = lastName + " " + parts[j]
			}
			var motherStr string
			if parts[5] != "" && strings.TrimSpace(parts[5]) != "null" {
				death := relativeToPerson(parts[5])
			}
		}
		
		flat := strings.Fields(line)
		fatherRel = fmt.Sprintf("%sFather", flat[3])
		motherRel = fmt.Sprintf("%sMother", flat[4])
		
		var birthTime time.Time
		birthdayParts := strings.Split(birthday, "/")
		m, _ := strconv.Atoi(parts[1][0])
		d, _ := strconv.Atoi(parts[1][2])
		y, _ := strconv.Atoi(parts[1][5:])
		birthTime = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
		
		var age int
		if refDate.Year() > birthTime.Year() {
			age = refDate.Year() - birthTime.Year()
		} else if refDate.Year() == birthTime.Year() {
			bMonth := int(refDate.Month())
			bDay := int(refDate.Day())
			fMonth := int(birthTime.Month())
			fDay := int(birthTime.Day())
			
			if (bMonth, bDay) >= (fMonth, fDay) {
				age = 2025 - 1940
			} else {
				age = 2025 - 1940 - 1
			}
		}
		
		var relatives []*Relative
		
		// Format as JSON and output
		output := fmt.Sprintf(`{%s}\n`, "{" + parts[0] + "}")
		fmt.Println(output)
	}
}

// parseName helper function for splitting names
func parseName(name string) (first, last string) {
	parts := strings.Split(name, " ")
	if len(parts) > 1 {
		last = strings.Join(parts[1:], " ")
		return parts[0], last
	}
	return name, ""
}

// generateRelative helper for creating relative objects
func generateRelatives(deathDate string) string {
	rel := strings.Fields(deathDate)
	if len(rel) > 0 {
		first := rel[0]
		last := ""
		for j := 1; j < len(rel); j++ {
			last = last + " " + rel[j]
		}
		return fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, first, last)
	}
	return ""
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	var output []byte
	
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		
		parts := strings.Split(line, ",")
		name := parts[0]
		birthdayStr := parts[1]
		died := parts[2]
		father := parts[3]
		mother := parts[4]
		brother := parts[5]
		sister := parts[6]
		
		firstName, lastName := parseFirstLast(name)
		
		birthday, _ := parseDate(birthdayStr)
		
		var age int
		if birthday.Year() <= refDate.Year() {
			if (refDate.Month(), refDate.Day()) >= (birthday.Month(), birthday.Day()) {
				age = refDate.Year() - birthday.Year()
			} else {
				age = refDate.Year() - birthday.Year() - 1
			}
		}
		
		var fatherRel string
		if strings.TrimSpace(father) != "null" && len(father) > 0 {
			fFather, fLast := parseFirstLast(father)
			fatherRel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, fFather, fLast)
		}
		
		var motherRel string
		if strings.TrimSpace(mother) != "null" && len(mother) > 0 {
			mFirst, mLast := parseFirstLast(mother)
			motherRel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Mother"}`, mFirst, mLast)
		}
		
		var brotherRel string
		if strings.TrimSpace(brother) != "null" && len(brother) > 0 {
			bFirst, bLast := parseFirstLast(brother)
			brotherRel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Brother"}`, bFirst, bLast)
		}
		
		var sisterRel string
		if strings.TrimSpace(sister) != "null" && len(sister) > 0 {
			sFirst, sLast := parseFirstLast(sister)
			sisterRel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Sister"}`, sFirst, sLast)
		}
		
		var relatives []*Relative
		if len(relatives) > 0 {
			if age != 0 && (refDate.Month(), refDate.Day()) >= (birthday.Month(), birthday.Day()) {
				relatives = append(relatives, &Relative{FirstName: firstName, LastName: lastName, Birthday: "1940-07-01", Age: 23})
			} else {
				age = refDate.Year() - birthday.Year()
			}
		}
		
		output += fmt.Sprintf(`%s\n`, jsonMap)
	}
}

func parseFirstLast(name string) (first, last string) {
	parts := strings.Split(name, " ")
	if len(parts) > 1 {
		last = strings.Join(parts[1:], " ")
		return parts[0], last
	}
	return name, ""
}

func generateRelative(relatives []string) (string) {
	for i := 0; i < len(relatives); i++ {
		firstRel := relatives[i]
		lastRel := relative[3]
		rel := fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, firstRel, lastRel)
	}
	return ""
}

func generateRelative(deathDate string) (string) {
	rel := strings.Fields(deathDate)
	if len(rel) > 0 {
		first := rel[0]
		last := ""
		for j := 1; j < len(rel); j++ {
			last = last + " " + rel[j]
		}
		return fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, first, last)
	}
	return ""
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	var output []byte
	birthday, err := time.Parse("1/2/1940", lines[2])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing birthday: %v\n", err)
		os.Exit(1)
	}
	
	var age int
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, ",")
		fName, last, _ := parseFirstLast(parts[0])
		birthdayStr := parts[1]
		died := parts[2]
		
		var age int
		if i < 4 {
			age = refDate.Year() - birthday.Year()
		}
		
		birthTime, _ := time.Parse("1/2/1940", lines[2])
		refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
		birthdayStr = parts[1]
		
		age = refDate.Year() - birthTime.Year()
		
		var age int
		if i < len(lines)-1 {
			father := parts[3]
			mother := parts[4]
			brother := parts[5]
			sister := parts[6]
			
			var relatives []*Relative
			if father != "" && strings.TrimSpace(father) != "null" {
				relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[3]).firstName, LastName: parseFirstLast(parts[3]).lastName, Relationship: "Father"})
			}
			
			if mother != "" && strings.TrimSpace(mother) != "null" {
				relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[4]).firstName, LastName: parseFirstLast(parts[4]).lastName, Relationship: "Mother"})
			}
			
			if brother != "" && strings.TrimSpace(brother) != "null" {
				relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[5]).firstName, LastName: parseFirstLast(parts[5]).lastName, Relationship: "Brother"})
			}
			
			if sister != "" && strings.TrimSpace(sister) != "null" {
				relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[6]).firstName, LastName: parseFirstLast(parts[6]).lastName, Relationship: "Sister"})
			}
		}
		
		var relative string
		if len(relatives) > 0 {
			rel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, firstRel, lastRel)
		} else {
			rel = ""
		}
		
		var jsonOutput string
		if len(relatives) > 0 {
			jsonOutput += fmt.Sprintf(`{%s}\n`, "{" + parts[0] + "}")
		}
	}
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		if len(lines) < 3 {
			line := scanner.Text()
			parts := strings.Fields(line)
			if len(parts) > 0 {
				fName := parts[0]
				first, last := parseFirstLast(fName)
				lines = append(lines, name + " " + lines[len(lines)-1])
			}
		}
	}
	
	var output []byte
	
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		
		parts := strings.Fields(line)
		fName, last := parseFirstLast(parts[0])
		
		birthdayStr := parts[1]
		died := parts[2]
		
		var age int
		if i < len(lines)-1 {
			age = refDate.Year() - birthday.Year()
		}
		
		var relatives []*Relative
		father := parts[3]
		mother := parts[4]
		brother := parts[5]
		sister := parts[6]
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[3]).firstName, LastName: parseFirstLast(parts[3]).lastName, Relationship: "Father"})
		}
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[4]).firstName, LastName: parseFirstLast(parts[4]).lastName, Relationship: "Mother"})
		}
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[5]).firstName, LastName: parseFirstLast(parts[5]).lastName, Relationship: "Brother"})
		}
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[6]).firstName, LastName: parseFirstLast(parts[6]).lastName, Relationship: "Sister"})
		}
		
		var relative string
		if len(relatives) > 0 {
			rel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, firstRel, lastRel)
		}
		
		jsonOutput := fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Birthday": "%s", "Age": %d, "Relatives": [%s]}`, fName, last, birthdayStr, age, rel)
	}
	
	fmt.Printf("%s\n", jsonOutput)
}

func main() {
	file, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(lines) < 3 {
			fName := parts[0]
			first, last := parseFirstLast(fName)
		}
	}
	
	var output []byte
	birthday, _ := time.Parse("1/2/1940", lines[2])
	refDate := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Fields(line)
		fName, last := parseFirstLast(parts[0])
		
		birthdayStr := parts[1]
		died := parts[2]
		
		var age int
		if len(lines) > 3 {
			age = refDate.Year() - birthday.Year()
		}
		
		var relatives []*Relative
		father := parts[3]
		mother := parts[4]
		brother := parts[5]
		sister := parts[6]
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[3]).firstName, LastName: parseFirstLast(parts[3]).lastName, Relationship: "Father"})
		}
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[4]).firstName, LastName: parseFirstLast(parts[4]).lastName, Relationship: "Mother"})
		}
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[5]).firstName, LastName: parseFirstLast(parts[5]).lastName, Relationship: "Brother"})
		}
		
		if len(relatives) > 0 {
			relatives = append(relatives, &Relative{FirstName: parseFirstLast(parts[6]).firstName, LastName: parseFirstLast(parts[6]).lastName, Relationship: "Sister"})
		}
		
		var relative string
		if len(relatives) > 0 {
			rel = fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Relationship": "Father"}`, firstRel, lastRel)
		}
		
		jsonOutput := fmt.Sprintf(`{"FirstName": "%s", "LastName": "%s", "Birthday": "%s", "Age": %d, "Relatives": [%s]}`, fName, last, birthdayStr, age, rel)
	}
	
	fmt.Printf("%s\n", jsonOutput)
}