import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface for a relative of a person.
 */
interface Relative {
    FirstName: string;
    LastName: string;
    Relationship: string;
}

/**
 * Interface for a person as defined in the expected output format.
 */
interface Person {
    FirstName: string;
    LastName: string;
    Birthday: string;
    Age: number;
    Relatives: Relative[];
}

/**
 * Parses a date string in M/D/YYYY format into a JavaScript Date object.
 * Returns null if the input is invalid or explicitly "null".
 */
function parseCSVDate(dateStr: string): Date | null {
    if (!dateStr || dateStr.trim() === 'null') return null;
    const parts = dateStr.split('/');
    if (parts.length !== 3) return null;
    
    const m = parseInt(parts[0], 10);
    const d = parseInt(parts[1], 10);
    const y = parseInt(parts[2], 10);
    
    // JavaScript months are 0-indexed (0 = January)
    return new Date(y, m - 1, d);
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
 * Calculates age based on birth date and a target date.
 */
function calculateAge(birthDate: Date, targetDate: Date): number {
    let age = targetDate.getFullYear() - birthDate.getFullYear();
    const monthDiff = targetDate.getMonth() - birthDate.getMonth();
    
    // Adjust age if the birthday has not yet occurred in the target year
    if (monthDiff < 0 || (monthDiff === 0 && targetDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

/**
 * Splits a full name into first and last name.
 * If the name has multiple parts, the first part is FirstName and the last part is LastName.
 */
function splitName(fullName: string) {
    const parts = fullName.trim().split(/\s+/);
    return {
        first: parts[0],
        last: parts.length > 1 ? parts[parts.length - 1] : parts[0]
    };
}

/**
 * Main function to read the CSV, transform the data, and output the JSON array.
 */
function main() {
    const inputFilePath = path.join('input', 'input.csv');
    
    let csvContent: string;
    try {
        csvContent = fs.readFileSync(inputFilePath, 'utf-8');
    } catch (err) {
        process.exit(1);
    }

    const lines = csvContent.split(/\r?\n/).filter(line => line.trim() !== '');
    if (lines.length === 0) {
        console.log('[]');
        return;
    }

    // The first line is the header: Name, Birthday, Died, Father, Mother, Brother, Sister
    const dataLines = lines.slice(1);
    
    // Target date for age calculation: July 1, 2025
    const targetDate = new Date(2025, 6, 1); 
    
    const results: Person[] = [];

    for (const line of dataLines) {
        const cols = line.split(',');
        if (cols.length < 7) continue;

        const [name, birthdayStr, diedStr, fatherStr, motherStr, brotherStr, sisterStr] = cols;

        const { first, last } = splitName(name);
        const bday = parseCSVDate(birthdayStr);
        if (!bday) continue;

        // Age calculation: if the person died, age is as of the death date; otherwise, as of July 1, 2025.
        const deathDate = parseCSVDate(diedStr);
        const referenceDate = deathDate ? deathDate : targetDate;
        const age = calculateAge(bday, referenceDate);

        const relatives: Relative[] = [];
        const relSources = [
            { val: fatherStr, rel: 'Father' },
            { val: motherStr, rel: 'Mother' },
            { val: brotherStr, rel: 'Brother' },
            { val: sisterStr, rel: 'Sister' },
        ];

        for (const source of relSources) {
            if (source.val && source.val.trim() !== 'null') {
                const { first: rFirst, last: rLast } = splitName(source.val);
                relatives.push({
                    FirstName: rFirst,
                    LastName: rLast,
                    Relationship: source.rel
                });
            }
        }

        results.push({
            FirstName: first,
            LastName: last,
            Birthday: formatDate(bday),
            Age: age,
            Relatives: relatives
        });
    }

    process.stdout.write(JSON.stringify(results, null, 2) + '\n');
}

main();