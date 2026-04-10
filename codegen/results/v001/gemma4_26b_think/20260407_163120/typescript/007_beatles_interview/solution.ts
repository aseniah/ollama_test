import * as fs from 'fs';
import * as path from 'path';

/**
 * Parses a single line of CSV, respecting double quotes.
 * 
 * @param line The raw CSV line string.
 * @returns An array of trimmed, unquoted cell strings.
 */
function parseCSVLine(line: string): string[] {
    const result: string[] = [];
    let current = '';
    let inQuotes = false;

    for (let i = 0; i < line.length; i++) {
        const char = line[i];
        if (char === '"') {
            inQuotes = !inQuotes;
        } else if (char === ',' && !inQuotes) {
            result.push(current.trim());
            current = '';
        } else {
            current += char;
        }
    }
    result.push(current.trim());
    // Remove surrounding quotes from cells
    return result.map(s => s.replace(/^"|"$/g, ''));
}

/**
 * Calculates the age of a person based on a birthdate string,
 * relative to the reference date of July 1, 2025.
 * 
 * @param dateStr The string representation of the birthdate.
 * @returns The calculated age as an integer.
 */
function calculateAge(dateStr: string): number {
    const refDate = new Date('2025-07-01');
    const birthDate = new Date(dateStr);

    // If date is invalid, return 0 or a default value
    if (isNaN(birthDate.getTime())) {
        return 0;
    }

    let age = refDate.getFullYear() - birthDate.getFullYear();
    const monthDiff = refDate.getMonth() - birthDate.getMonth();
    const dayDiff = refDate.getDate() - birthDate.getDate();

    // Adjust age if the birthday hasn't occurred yet in the reference year
    if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
        age--;
    }

    return age;
}

/**
 * Main execution function to read the CSV, transform it according to
 * the inferred rules, and output the resulting JSON array.
 */
function main() {
    const inputFilePath = path.join('input', 'input.csv');

    // Check if input file exists
    if (!fs.existsSync(inputFilePath)) {
        return;
    }

    let csvContent: string;
    try {
        csvContent = fs.readFileSync(inputFilePath, 'utf8');
    } catch (err) {
        return;
    }

    // Split content into lines and remove empty lines
    const lines = csvContent.split(/\r?\n/).filter(line => line.trim() !== '');

    if (lines.length === 0) {
        process.stdout.write('[]\n');
        return;
    }

    // The first line contains the CSV headers
    const headers = parseCSVLine(lines[0]);
    const dataRows = lines.slice(1);

    // Transform each row into an object
    const jsonOutput = dataRows.map(line => {
        const values = parseCSVLine(line);
        const obj: Record<string, any> = {};

        headers.forEach((header, index) => {
            const rawValue = values[index] !== undefined ? values[index] : '';
            // Normalize key to lowercase to match common JSON patterns
            let key = header.toLowerCase();
            let finalValue: any = rawValue;

            // Rule Inference: If the header indicates a date/birth, transform to 'age'
            if (key.includes('date') || key.includes('birth')) {
                key = 'age';
                finalValue = calculateAge(rawValue);
            } else {
                // Rule Inference: Convert numeric strings to actual numbers
                const numericValue = Number(rawValue);
                if (rawValue !== '' && !isNaN(numericValue)) {
                    finalValue = numericValue;
                }
            }

            obj[key] = finalValue;
        });

        return obj;
    });

    // Output the final JSON array to stdout
    process.stdout.write(JSON.stringify(jsonOutput, null, 2) + '\n');
}

// Execute the script
main();