import * as fs from 'fs';

/**
 * Parses a single line of CSV, handling quoted values.
 */
function parseCSVLine(line: string): string[] {
    const result: string[] = [];
    let current = '';
    let inQuotes = false;
    for (let i = 0; i < line.length; i++) {
        const char = line[i];
        if (char === '"') {
            in
            inQuotes = !inQuotes;
        } else if (char === ',' && !inQuotes) {
            result.push(current.trim());
            current = '';
        } else {
            current += char;
        }
    }
    result.push(current.trim());
    return result;
}

/**
 * Calculates age as of July 1, 2025.
 */
function calculateAge(birthDateStr: string): number {
    const refDate = new Date('2025-07-01T00:00:00');
    const birthDate = new Date(birthDateStr);
    
    if (isNaN(birthDate.getTime())) {
        return 0;
    }

    let age = refDate.getFullYear() - birthDate.getFullYear();
    const monthDiff = refDate.getMonth() - birthDate.getMonth();
    
    // If the birthday hasn't occurred yet in the reference year
    if (monthDiff < 0 || (monthDiff === 0 && refDate.getDate() < birthDate.getDate())) {
        age--;
    }
    
    return age;
}

/**
 * Main execution function.
 */
function main() {
    const csvFilePath = 'input/input.csv';
    const templateFilePath = 'input/expected_format.json';

    try {
        if (!fs.existsSync(csvFilePath) || !fs.existsSync(templateFilePath)) {
            return;
        }

        // 1. Load the expected format to understand the target keys and structure
        const expectedFormat = JSON.parse(fs.readFileSync(templateFilePath, 'utf-pass'));
        // We assume the first element of the array contains the template keys
        const templateKeys = Array.isArray(expectedFormat) && expectedFormat.length > 0 
            ? Object.keys(expectedFormat[0]) 
            : [];

        // 2. Load and parse the CSV file
        const csvContent = fs.readFileSync(csvFilePath, 'utf-8').trim();
        if (!csvContent) return;

        const csvLines = csvContent.split(/\r?\n/);
        const csvHeaders = parseCSVLine(csvLines[0]);

        const dataRows = csvLines.slice(1);
        const outputArray: any[] = [];

        // 3. Transform each CSV row into the expected JSON format
        for (const row of dataRows) {
            if (!row.trim()) continue;
            const values = parseCSVLine(row);
            const transformedEntry: any = {};

            // For every key required by the expected format
            for (const targetKey of templateKeys) {
                let foundValue: string | null = null;
                let isDateColumn = false;

                // Search through CSV headers to find a matching column for the targetKey
                for (let i = 0; i < csvHeaders.length; i++) {
                    const header = csvHeaders[i];
                    const lowerTarget = targetKey.toLowerCase();
                    const lowerHeader = header.toLowerCase();

                    // Logic for 'age' key: look for columns containing 'date' or 'birth'
                    if (lowerTarget === 'age') {
                        if (lowerHeader.includes('date') || lowerHeader.includes('birth') || lowerHeader.includes('dob')) {
                            foundValue = values[i];
                            isDateColumn = true;
                            break;
                        }
                    } else {
                        // Logic for other keys: check for direct or fuzzy string matches
                        if (lowerTarget === lowerHeader || 
                            lowerHeader.includes(lowerTarget) || 
                            lowerTarget.includes(lowerHeader)) {
                            foundValue = values[i];
                            break;
                        }
                    }
                }

                // If we found a matching column, apply transformation
                if (foundValue !== null) {
                    if (targetKey.toLowerCase() === 'age' && isDateColumn) {
                        transformedEntry[targetKey] = calculateAge(foundValue);
                    } else {
                        transformedEntry[targetKey] = foundValue;
                    }
                }
            }

            // Only add if the entry is not empty
            if (Object.keys(transformedEntry).length > 0) {
                outputArray.push(transformedEntry);
            }
        }

        // 4. Output the final JSON array to stdout
        process.stdout.write(JSON.stringify(outputArray, null, 2) + '\n');

    } catch (error) {
        // Error handling: ensure no output other than JSON is sent to stdout
        // Error messages would typically go to stderr in a production environment
    }
}

main();