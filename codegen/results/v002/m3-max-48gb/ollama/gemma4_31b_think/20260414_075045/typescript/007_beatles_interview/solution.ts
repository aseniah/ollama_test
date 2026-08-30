import * as fs from 'fs';
import * as path from 'path';

/**
 * Interface for the Relatives object in the output JSON
 */
interface Relative {
    FirstName: string;
    LastName: string;
    Relationship: string;
}

/**
 * Interface for the Person object in the output JSON
 */
interface Person {
    FirstName: string;
    LastName: string;
    Birthday: string;
    Age: number;
    Relatives: Relative[];
}

/**
 * Splits a full name string into first and last names.
 * Based on the expected format, it takes the first part as FirstName
 * and the final part as LastName.
 */
function splitName(fullName: string) {
    const parts = fullName.trim().split(/\s+/);
    return {
        FirstName: parts[0],
        LastName: parts[parts.length - 1]
    };
}

/**
 * Converts a date string in M/D/YYYY format to YYYY-MM-DD format.
 */
function formatDate(dateStr: string): string {
    const [m, d, y] = dateStr.split('/');
    return `${y}-${m.padStart(2, '0')}-${d.padStart(2, '0')}`;
}

/**
 * Calculates age relative to a reference date.
 * If the person has a death date, that is used as the reference.
 * Otherwise, July 1, 2025 is used.
 */
function calculateAge(birthStr: string, deathStr: string | null): number {
    const [bM, bD, bY] = birthStr.split('/').map(Number);
    
    let rM: number, rD: number, rY: number;
    if (deathStr && deathStr !== 'null') {
        const [dM, dD, dY] = deathStr.split('/').map(Number);
        rM = dM; 
        rD = dD; 
        rY = dY;
    } else {
        // July 1, 2025
        rM = 7; 
        rD = 1; 
        rY = 2025;
    }

    let age = rY - bY;
    // Adjust age if the birthday has not yet occurred in the reference year
    if (rM < bM || (rM === bM && rD < bD)) {
        age--;
    }
    return age;
}

/**
 * Main execution function to read the CSV and output the transformed JSON.
 */
function solve() {
    try {
        const csvPath = path.join('input', 'input.csv');
        const content = fs.readFileSync(csvPath, 'utf-8');
        const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
        
        if (lines.length === 0) {
            process.stdout.write('[]\n');
            return;
        }

        // The first line is the CSV header: Name,Birthday,Died,Father,Mother,Brother,Sister
        const dataRows = lines.slice(1);
        const result: Person[] = [];

        for (const line of dataRows) {
            const cols = line.split(',');
            if (cols.length < 7) continue;

            const nameFull = cols[0];
            const birthday = cols[1];
            const died = cols[2];
            const father = cols[3];
            const mother = cols[4];
            const brother = cols[5];
            const sister = cols[6];

            const nameParts = splitName(nameFull);
            const relatives: Relative[] = [];

            // List of family columns to map to the Relatives array
            const relativeDefinitions = [
                { name: father, rel: 'Father' },
                { name: mother, rel: 'Mother' },
                { name: brother, rel: 'Brother' },
                { name: sister, rel: 'Sister' },
            ];

            for (const def of relativeDefinitions) {
                if (def.name && def.name !== 'null') {
                    const rNameParts = splitName(def.name);
                    relatives.push({
                        FirstName: rNameParts.FirstName,
                        LastName: rNameParts.LastName,
                        Relationship: def.rel
                    });
                }
            }

            result.push({
                FirstName: nameParts.FirstName,
                LastName: nameParts.LastName,
                Birthday: formatDate(birthday),
                Age: calculateAge(birthday, died),
                Relatives: relatives
            });
        }

        process.stdout.write(JSON.stringify(result, null, 2) + '\n');
    } catch (err) {
        // Ensure no output other than JSON is produced on successful runs.
        // Errors are handled silently or through standard node error behavior.
    }
}

solve();