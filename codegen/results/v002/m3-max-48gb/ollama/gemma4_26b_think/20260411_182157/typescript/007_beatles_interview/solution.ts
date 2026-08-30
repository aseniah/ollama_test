import * as fs from 'fs';
import * as path from 'path';

/**
 * Parses a date string in the format M/D/YYYY into a Date object.
 * Returns null if the format is invalid or the string is 'null'.
 */
function parseDate(dateStr: string): Date | null {
    if (!dateStr || dateStr.trim().toLowerCase() === 'null') return null;
    const parts = dateStr.trim().split('/');
    if (parts.length !== 3) return null;
    
    const month = parseInt(parts[0], 10) - 1;
    const day = parseInt(parts[1], 10);
    const year = parseInt(parts[2], 10);
    
    const date = new Date(year, month, day);
    // Check if the date is valid (e.g., not Feb 31st)
    if (isNaN(date.getTime())) return null;
    return date;
}

/**
 * Formats a Date object into a YYYY-MM-DD string.
 */
function formatDate(date: Date): string {
    const y = date.getFullYear();
    const m = String(date.getMonth() + 1).padStart(2, '0');
    const d = String(date.getDate()).padStart(2, '0');
    return `${y}-${m}-${d}`;
}

/**
 * Calculates the age in years between a birth date and an end date.
 */
function calculateAge(birthDate: Date, endDate: Date): number {
    let age = endDate.getFullYear() - birthDate.getFullYear();
    const m = endDate.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && endDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

/**
 * Extracts FirstName and LastName from a full name string.
 * FirstName is the first part, LastName is the last part.
 */
function parseNameParts(fullName: string) {
    if (!fullName || fullName.trim().toLowerCase() === 'null') return null;
    const parts = fullName.trim().split(/\s+/);
    if (parts.length === 0) return null;
    
    return {
        FirstName: parts[0],
        LastName: parts[parts.length - 1]
    };
}

/**
 * Main function to process the CSV and produce the JSON output.
 */
function solve() {
    const inputFilePath = path.join('input', 'input.csv');
    
    // Check if input file exists
    if (!fs.existsSync(inputFilePath)) {
        return;
    }

    const csvContent = fs.readFileSync(inputFilePath, 'utf-8');
    const lines = csvContent.split(/\r?\n/).map(l => l.trim()).filter(l => l.length > 0);
    
    if (lines.length <= 1) {
        process.stdout.write('[]\n');
        return;
    }

    // Target date for calculating age if 'Died' is null
    const targetDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

    // Skip the header line
    const dataRows = lines.slice(1);

    const results = dataRows.map(line => {
        const cols = line.split(',').map(c => c.trim());
        
        // Column Mapping based on input.csv:
        // 0: Name, 1: Birthday, 2: Died, 3: Father, 4: Mother, 5: Brother, 6: Sister
        const nameParts = parseNameParts(cols[0]);
        if (!nameParts) return null;

        const birthday = parseDate(cols[1]);
        const died = parseDate(cols[2]);

        // The requirements suggest Birthday is mandatory for the record
        if (!birthday) return null;

        // If 'Died' is not null, calculate age at death. Otherwise, calculate age as of July 1, 2025.
        const age = died ? calculateAge(birthday, died) : calculateAge(birthday, targetDate);

        const relatives: any[] = [];
        const relationMappings = [
            { index: 3, relationship: 'Father' },
            { index: 4, relationship: 'Mother' },
            { index: 5, relationship: 'Brother' },
            { index: 6, relationship: 'Sister' }
        ];

        for (const mapping of relationMappings) {
            const relNameStr = cols[mapping.index];
            if (relNameStr && relNameStr.toLowerCase() !== 'null') {
                const relParts = parseNameParts(relNameStr);
                if (relParts) {
                    relatives.push({
                        FirstName: relParts.FirstName,
                        LastName: relParts.LastName,
                        Relationship: mapping.relationship
                    });
                }
            }
        }

        return {
            FirstName: nameParts.FirstName,
            LastName: nameParts.LastName,
            Birthday: formatDate(birthday),
            Age: age,
            Relatives: relatives
        };
    }).filter((item): item is any => item !== null);

    // Print the JSON array to stdout
    process.stdout.write(JSON.stringify(results, null, 2) + '\n');
}

// Execute the script
solve();