import * as fs from 'fs';
import * as path from 'path';

const inputFile = 'input/input.csv';
const referenceFile = 'input/expected_format.json';

function parseDate(dateString: string): Date {
    // Simple date parsing: YYYY-MM-DD
    const [year, month, day] = dateString.split('-').map(Number);
    return new Date(year, month - 1, day);
}

function main() {
    if (!fs.existsSync(inputFile)) {
        console.error(`Input file not found: ${inputFile}`);
        process.exit(1);
    }

    const rows = fs.readFileSync(inputFile, 'utf-8').trim().split('\n');
    
    // Assuming first row is a header (Name,BirthDate) based on typical CSV structures for this task.
    // If the first row isn't headers, this might fail logic, but standard assumption applies.
    if (rows.length < 2) {
        console.error('CSV file does not contain headers and data.');
        process.exit(1);
    }

    const header = rows[0].split(',');
    
    // Index map for columns
    const indexMap: Record<string, number> = {};
    header.forEach((col, idx) => {
        indexMap[col.trim()] = idx;
    });

    const birthDateKey = indexMap['BirthDate'] || undefined;

    const referenceContent = fs.readFileSync(referenceFile, 'utf-8');
    // Basic parsing of the expected format to confirm structure (JSON array of objects)
    if (referenceContent && referenceContent.trim()) {
        try {
            JSON.parse(referenceContent);
        } catch {} 
    }

    const results: any[] = [];

    for (let i = 1; i < rows.length; i++) {
        const row = rows[i];
        
        // Handle potential quoted CSV fields or commas inside strings if needed, 
        // but assuming standard comma-separated values for simplicity as per "Node.js built-in" constraint.
        // A simple split by ',' might fail on quotes. However, without a specific delimiter example, 
        // we assume simple splitting. If 'BirthDate' is empty or missing, skip logic that depends on it.
        
        const parts = row.split(',');
        const name = (indexMap['Name'] >= 0) ? parts[indexMap['Name']] : '';
        let birthDateString = '';

        if (indexMap['BirthDate']) {
            birthDateString = parts[indexMap['BirthDate']].trim();
        } else if (birthDateStringKey) { // Fallback if logic changed
             birthDateString = rows[i].split(',')[1]?.trim() || '';
        }

        if (!birthDateString) {
            continue; 
        }

        const birthDate = parseDate(birthDateString);
        
        // Reference date: July 1, 2025 (Time zone neutral calculation)
        const referenceDate = new Date(2025, 6, 1); 

        let ageInYears = 0;
        let birthYear = 0;
        let birthMonth = 0;

        if (birthDate.getFullYear() > referenceDate.getFullYear()) {
            ageInYears = referenceDate.getFullYear() - birthDate.getFullYear() - 1;
        } else if (birthDate.getMonth() > referenceDate.getMonth() || 
                   (birthDate.getMonth() === referenceDate.getMonth() && 
                    birthDate.getDate() >= referenceDate.getDate())) {
            ageInYears = referenceDate.getFullYear() - birthDate.getFullYear();
        } else {
            ageInYears = referenceDate.getFullYear() - birthDate.getFullYear() - 1;
        }

        // Ensure minimum age of 0 if logic is off, though math should handle it.
        if (ageInYears < 0) ageInYears = 0;

        const outputRecord = {
            name: name.trim(),
            birthDate: birthDateString,
            age: ageInYears // The format implies 'age' as the calculated number. Reference shows "age": 25.
        };

        results.push(outputRecord);
    }

    console.log(JSON.stringify(results, null, 2));
}

main();