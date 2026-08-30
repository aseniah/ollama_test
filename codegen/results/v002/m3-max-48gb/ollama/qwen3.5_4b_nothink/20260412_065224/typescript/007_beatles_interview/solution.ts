import * as fs from 'fs';
import * as path from 'path';

// Read input CSV
const csvContent = fs.readFileSync(new URL('input/input.csv', import.meta.url), 'utf-8');

// Split CSV content by lines
const lines = csvContent.trim().split('\n');

// Parse CSV data
// Handle potential quoted fields or multiple spaces (though the sample seems simple)
// We'll split by comma, but be careful with potential edge cases.
// The sample shows standard comma separation without quotes around fields, 
// except potentially if the format changes, but the provided sample is straightforward.

function parseLine(line: string) {
    // Split by comma
    const fields = line.split(',');
    
    // Extract first name and last name from the "Name" field
    // Format: "First Last"
    const parts = fields[0].trim().split(/\s+/);
    if (parts.length < 2) {
        throw new Error(`Invalid name format in line: ${line}`);
    }
    const firstName = parts[0];
    const lastName = parts[1];
    
    // Birthday is MM/DD/YYYY or DD/MM/YYYY?
    // Looking at the sample:
    // John Lennon: 10/9/1940 -> 1940-10-09 (Month/Day/Year)
    // James McCartney: 6/18/1942 -> 1942-06-18 (Month/Day/Year)
    // Ringo Starr: 7/7/1940 -> 1940-07-07 (Month/Day/Year)
    // George Harrison: 2/25/1943 -> 1943-02-25 (Month/Day/Year)
    
    const dateParts = fields[1].trim().split('/');
    if (dateParts.length !== 3) {
        throw new Error(`Invalid date format in line: ${line}`);
    }
    const month = parseInt(dateParts[0], 10);
    const day = parseInt(dateParts[1], 10);
    const year = parseInt(dateParts[2], 10);
    
    // Birthday as YYYY-MM-DD
    const birthday = `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
    
    // Calculate age as of July 1, 2025
    const referenceDate = new Date(2025, 6, 1); // July is index 6
    
    function calculateAge(birthdayString: string): number {
        const birthDate = new Date(birthdayString);
        let age = referenceDate.getFullYear() - birthDate.getFullYear();
        
        // Check if birthday has already occurred in the current year (July 1st)
        // Reference: July 1st, 2025
        const refMonth = referenceDate.getMonth(); // 6 (July)
        const refDay = referenceDate.getDate(); // 1
        
        const bMonth = birthDate.getMonth();
        const bDay = birthDate.getDate();
        
        // If birthday is before or on July 1st, age is correct
        // If birthday is after July 1st, subtract 1
        if (bMonth < refMonth || (bMonth === refMonth && bDay <= refDay)) {
            return age;
        } else {
            return age - 1;
        }
    }
    
    const age = calculateAge(birthday);
    
    // Build relative names array
    const relatives = [];
    
    // Father
    const father = fields[3].trim();
    if (father !== 'null') {
        const fatherParts = father.split(' ');
        relatives.push({
            "FirstName": fatherParts[0],
            "LastName": fatherParts[1],
            "Relationship": "Father"
        });
    }
    
    // Mother
    const mother = fields[4].trim();
    if (mother !== 'null') {
        const motherParts = mother.split(' ');
        relatives.push({
            "FirstName": motherParts[0],
            "LastName": motherParts[1],
            "Relationship": "Mother"
        });
    }
    
    // Brother
    const brother = fields[5].trim();
    if (brother !== 'null') {
        const brotherParts = brother.split(' ');
        relatives.push({
            "FirstName": brotherParts[0],
            "LastName": brotherParts[1],
            "Relationship": "Brother"
        });
    }
    
    // Sister
    const sister = fields[6].trim();
    if (sister !== 'null') {
        const sisterParts = sister.split(' ');
        relatives.push({
            "FirstName": sisterParts[0],
            "LastName": sisterParts[1],
            "Relationship": "Sister"
        });
    }
    
    // Construct the object
    const person = {
        "FirstName": firstName,
        "LastName": lastName,
        "Birthday": birthday,
        "Age": age,
        "Relatives": relatives
    };
    
    return person;
}

const results = [];
for (const line of lines) {
    if (line.trim() === '') continue;
    const parsed = parseLine(line);
    results.push(parsed);
}

// Output JSON to stdout
console.log(JSON.stringify(results, null, 2));