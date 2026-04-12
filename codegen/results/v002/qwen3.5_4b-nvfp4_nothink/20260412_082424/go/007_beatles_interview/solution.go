package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Helper function to parse a "Month/Day/Year" string into time.Time
func parseDate(d string) (time.Time, error) {
	if d == "" || d == "null" {
		return time.Time{}, nil // Return zero time for null values to avoid parsing errors
	}
	parts := strings.Split(d, "/")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", d)
	}

	month, _ := strconv.Atoi(parts[0])
	day, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

func main() {
	// Read input file
	inputFile, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	var csvData []map[string]string
	scanner := bufio.NewReader(inputFile)
	// Parse CSV rows (handles commas in data if necessary, though names don't seem to have commas here based on context)
	csvReader := csv.NewReader(scanner)
	
	// Skip header row
	_, _ = csvReader.Read()

	var people []Person
	
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
			os.Exit(1)
		}

		var person Person
		
		// Name is the first column (Name), but we need to split it if needed? 
		// Looking at input: "John Winston Lennon" -> firstName: "John", lastName: "Lennon"
		// The CSV columns are: Name, Birthday, Died, Father, Mother, Brother, Sister
		// Format looks like FirstLast or LastFirst. Based on Beatles, it's usually FirstName LastName.
		// However, the input data format is weird for George: "George Harrison".
		// Let's look at the mapping logic derived from expected output vs input:
		
		// Input Name | Birth      | Father            | Mother             | Brother          | Sister
		// John         10/9/1940   Alfred Lennon     Julia Stanley       null              null
		// James        6/18/1942    Jim McCartney      Mary McCartney       Mike McGear       null
		// Ringo        7/7/1940     Richard Starkey    Elsie Gleave         null          Marie Maguire
		// George       2/25/1943    Harold Harrison    Louise French      Peter Harrison     Louise Harrison

		// Logic for First/Last Name extraction:
		// Most names are FirstName LastName (e.g., "John Lennon").
		 // Some have middle names in the CSV but the output has 2 parts.
		 // We need to strip the Middle name if present?
		 // John Winston Lennon -> John, Lennon (Middle is Winston)
		 // James Paul McCartney -> James, McCartney (Middle is Paul)
		 // Ringo Starr -> Ringo, Starr (No middle in output, but input says "Ringo Starr" - wait, input says "Ringo Starr" at start of row?)
		 // Row: Name,Birthday... | Ringo,7/7/1940... 
		 // Input text shows: "Ringo Starr,7/7/1940..." -> First: Ringo, Last: Starr.
		
		// Let's assume the first word is FirstName and the last word is LastName for now?
		// Or check length > 2?
		// John Winston Lennon (3 words) -> "John", "Lennon"
		// James Paul McCartney (3 words) -> "James", "McCartney"
		
		fullName := strings.TrimSpace(record[0])
		parts := strings.Fields(fullName)
		
		if len(parts) == 1 {
			// Case: "George Harrison" -> First: George, Last: Harrison
			person.FirstName = parts[0]
			person.LastName = parts[len(parts)-1]
		} else if len(parts) == 2 {
			// Case: "Ringo Starr" -> First: Ringo, Last: Starr
			person.FirstName = parts[0]
			person.LastName = parts[1]
		} else {
			// Case: "John Winston Lennon" -> Take first and last
			person.FirstName = parts[0]
			person.LastName = parts[len(parts)-1]
		}

		// Parse Birthday
		birthDateStr := strings.TrimSpace(record[1])
		parsedBirth, _ := parseDate(birthDateStr)
		
		calcAge(person, parsedBirth) // We use the function to calculate age and set Age field
        
		// Construct Relatives slice
		for _, col := range []string{"Father", "Mother", "Brother", "Sister"} {
			val := strings.TrimSpace(record[2+3+index]) 
			// Wait, indices: Name(0), Birthday(1), Died(2), Father(3), Mother(4), Brother(5), Sister(6)
			colIndex := 3 + (col == "Father" ? 0 : 4) // No, that's wrong.
			// Let's map explicitly:
			// Father is index 3
			// Mother is index 4
			// Brother is index 5
			// Sister is index 6
			
			// Re-indexing properly based on slice length
			cols := []string{"Father", "Mother", "Brother", "Sister"}
			for i, colName := range cols {
				colIndex := 2 + i // Birthday(1) + Died(2) = 3. Father is index 3. Mother is index 4.
				
				// Wait: Name(0), Birth(1), Died(2). Father is 3. Correct.
				val := strings.TrimSpace(record[colIndex])
				
				if val == "" || val == "null" {
					continue
				}

				// Parse Full Name of Relative (Format: First Last)
				relNameParts := strings.Fields(val)
				// The expected output has only 2 parts (First, Last).
				// Input might have more if they had middle names? 
				// "Alfred Lennon" -> Alfred, Lennon.
				// "Mike McGear" -> Mike, McGear.
				
				relFirstName := relNameParts[0]
				// If there are 3 parts (e.g. John Winston Lennon), do we keep middle? 
				// Looking at input data, it seems like the names provided for relatives don't have middle names in the CSV.
				// But wait, "John Winston Lennon" has a middle name. The input CSV just lists "Alfred Lennon".
				// So the CSV relative field is First Last.
				
				if len(relNameParts) == 2 {
					person.Relatives = append(person.Relatives, Relative{
						FirstName: relNameParts[0],
						LastName:  relNameParts[1],
						Relationship: colName,
					})
				} else if len(relNameParts) == 3 {
					// Just in case, though unlikely based on typical data shown.
					// If middle name exists, do we use it? Expected format doesn't show a specific Middle field.
					// We will take First and Last as is.
					person.Relatives = append(person.Relatives, Relative{
						FirstName: relNameParts[0],
						LastName:  relNameParts[len(relNameParts)-1], // Take last word if > 2? 
						// Actually, let's assume the input relative names are "First Last".
					})
				}
			}
		}
		people = append(people, person)
	}

	var output []PersonJSON
	for _, p := range people {
		output = append(output, PersonJSON{
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Birthday:  fmt.Sprintf("%d-%02d-%02d", p.Birth.Year, p.Birth.Month, p.Birth.Day),
			Age:       p.Age,
			Relatives: p.Relatives,
		})
	}

	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}

func parseDate(d string) (time.Time, error) {
	if d == "" || d == "null" {
		return time.Time{}, nil
	}
	parts := strings.Split(d, "/")
	month, _ := strconv.Atoi(parts[0])
	day, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

type Person struct {
	FirstName string
	LastName   string
	Birth      time.Time
	Age        int
	Relatives  []Relative
}

type Relative struct {
	FirstName  string
	LastName   string
	Relationship string
}

// Calculation logic for age as of July 1, 2025.
func calcAge(person *Person, birthTime time.Time) {
	refDate := time.Date(2025, time.Month(7), 1, 0, 0, 0, 0, time.Local)

	// Calculate full age: Years difference
	years := refDate.Year() - birthTime.Year()

	// Check if birthday has passed this year (July 1)
	birthMonth, _ := birthTime.MarshalBinary() // Wrong way to get month
	birthdayParts := strings.Split(fmt.Sprintf("%d/%d", birthTime.Month(), birthTime.Day()), "/") 
	// Simpler: string slice
	mArr := strings.Split(birthDateStr := fmt.Sprintf("%d-%02d-%02d", 1, 7, 2025), "-") // Wait, let's just check date directly
	
	// Compare time.Time
	if refDate.Before(birthTime) {
		years--
	} else if refDate.Year() == birthTime.Year() && 
	       (refDate.Month() < birthTime.Month() || (refDate.Month() == birthTime.Month() && refDate.Day() < birthTime.Day())) {
		// Birthday passed, so we are younger than 'year - 1'
		years-- // Actually, the logic is: if birthday passed, age = year_diff. If not passed, age = year_diff - 1.
		// If refDate is Jul 1 and Birth is Jan 5 (Year Y). RefDate > Birth? No, RefDate is same year but later month.
		// Wait, let's simplify.
	}

	if !refDate.After(birthTime) { // Person died before or on ref date (if Died != null)
		// If the person has already had their birthday in the target year, age = current_year - birth_year
		// But if the reference date is Jul 1, and Birth is Aug 5. They haven't turned that age yet.
		
		if refDate.Year() == birthTime.Year() && (refDate.Month() < birthTime.Month() || (refDate.Month() == birthTime.Month() && refDate.Day() < birthTime.Day())) {
			person.Age = years - 1
		} else {
			person.Age = years
		}
		return
	}

	// Logic: 
	// If reference date is BEFORE birthday: Age = YearDiff - 1
	// Else (after or same): Age = YearDiff

	if refDate.Before(birthTime) {
		person.Age = years - 1
	} else if refDate.After(birthTime) {
		person.Age = years
	} else {
		// Same day? 
		// If July 1 is same as Birthday: Age is YearDiff.
		// Example: Born July 1. Age on July 1 is YearDiff (just turned).
		person.Age = years
	}
}

type PersonJSON struct {
	FirstName string      `json:"FirstName"`
	LastName   string      `json:"LastName"`
	Birthday   string      `json:"Birthday"`
	Age        int         `json:"Age"`
	Relatives  []RelRelative `json:"Relatives"`
}

type RelRelative struct {
	FirstName  string      `json:"FirstName"`
	LastName   string      `json:"LastName"`
	Relationship string     `json:"Relationship"`
}

func main() {
	inputFile, err := os.Open("input/input.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	var csvData []map[string]string
	scanner := bufio.NewReader(inputFile)
	csvReader := csv.NewReader(scanner)
	
	_, _ = csvReader.Read() // Skip header
	
	var people []Person
	birthDateStr := "" // Global scope not supported, make local inside loop or pass it.

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading CSV:", err)
			os.Exit(1)
		}

		var person Person
		
		fullName := strings.TrimSpace(record[0])
		parts := strings.Fields(fullName)
		
		if len(parts) == 1 {
			person.FirstName = parts[0]
			person.LastName = parts[0]
		} else if len(parts) == 2 {
			person.FirstName = parts[0]
			person.LastName = parts[1]
		} else {
			person.FirstName = parts[0]
			person.LastName = parts[len(parts)-1]
		}

		birthStr := strings.TrimSpace(record[1])
		parsedBirth, _ := parseDate(birthStr)
		
		calcAge(person, parsedBirth, birthStr) // Need to pass birth string for date comparison inside calc if time parsing is tricky
		
		// Relative Mapping Indices:
		// 0: Name, 1: Birth, 2: Died
		// 3: Father, 4: Mother, 5: Brother, 6: Sister
		cols := []string{"Father", "Mother", "Brother", "Sister"}
		
		for _, colName := range cols {
			valIdx := 3 + (len([]string{colName}) - 1) // Not quite right mapping logic again.
			// Fixed indices: Father=3, Mother=4, Brother=5, Sister=6
			if colName == "Father" { idx := 3 }
			if colName == "Mother" { idx := 4 }
			if colName == "Brother" { idx := 5 }
			if colName == "Sister" { idx := 6 }
			
			val := strings.TrimSpace(record[idx])
			if val == "" || val == "null" {
				continue
			}
			
			relParts := strings.Fields(val)
			if len(relParts) >= 2 {
				person.Relatives = append(person.Relatives, RelRelative{
					FirstName:  relParts[0],
					LastName:   relParts[len(relParts)-1], // Take last if > 2? Or just 1st if always 2.
					Relationship: colName,
				})
			}
		}
		
		// Wait, the calculation function needs the birth date to work correctly.
		calcAgeAndSet(person, parsedBirth)
		
		people = append(people, person)
	}

	var output []PersonJSON
	for _, p := range people {
		output = append(output, PersonJSON{
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Birthday:  fmt.Sprintf("%d-%02d-%02d", p.Birth.Year, p.Birth.Month, p.Birth.Day),
			Age:       p.Age,
			Relatives: p.Relatives,
		})
	}

	jsonOutput, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(jsonOutput))
}

func calcAgeAndSet(person *Person, birthTime time.Time) {
	refDate := time.Date(2025, time.Month(7), 1, 0, 0, 0, 0, time.Local)
	
	// Simple integer math logic for age calculation on a specific date.
	// Age = YearDiff - 1 if birthday hasn't passed yet (or is after ref day).
	// Age = YearDiff if birthday has passed or is today.
	
	if !refDate.After(birthTime) { // Passed birthdate or died before
		// If the person is still alive? The code doesn't check death date because we only have birth date for age calc.
		// But we need to know if they are dead? 
		// The expected output shows Age based purely on calculation. 
		// Let's assume: Age = YearRef - YearBorn - 1 if (Month < RefMonth OR (Month == RefMonth AND Day > RefDay)) else YearDiff.
		
		if refDate.Before(birthTime) || (refDate.Year() < birthTime.Year()) { // Impossible to be older than born
			person.Age = 0 // Should not happen
		} else if (refDate.Month() < birthTime.Month()) || (refDate.Month() == birthTime.Month() && refDate.Day() < birthTime.Day()) {
			// Birthday not reached yet
			person.Age = refDate.Year() - birthTime.Year() - 1
		} else {
			// Birthday passed or is today
			person.Age = refDate.Year() - birthTime.Year()
		}
	} else { // Already born past birthday
		person.Age = refDate.Year() - birthTime.Year()
	}
}