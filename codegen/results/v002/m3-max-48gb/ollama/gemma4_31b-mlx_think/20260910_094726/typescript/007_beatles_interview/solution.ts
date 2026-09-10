import * as fs from 'fs';
import * as path from 'path';

/**
 * Parses a date string in the format M/D/YYYY into a JavaScript Date object.
 * Returns null if the input string is 'null' or empty.
 */
function parseDate(dateStr: string): Date | null {
    if (!dateStr || dateStr === 'null') return null;
    const parts = dateStr.split('/');
    if (parts.length !== 3) return null;
    const [m, d, y] = parts.map(Number);
    // JavaScript Date months are 0-indexed (January is 0).
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
 * Calculates the age of a person given their birthday and a reference date.
 */
function calculateAge(birthDate: Date, refDate: Date): number {
    let age = refDate.getFullYear() - birthDate.getFullYear();
    const m = refDate.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && refDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

/**
 * Splits a full name into a FirstName and a LastName.
 * Takes the first word as FirstName and the last word as LastName.
 */
function splitName(name: string) {
    const parts = name.trim().split(/\s+/);
    return {
        FirstName: parts[0],
        LastName: parts[parts.length - 1]
    };
}

function main() {
    const inputPath = path.join('input', 'input.csv');
    let csvContent: string;
    
    try {
        csvContent = fs.readFileSync(inputPath, 'utf8');
    } catch (err) {
        return;
    }

    const lines = csvContent.split(/\r?\n/).filter(l => l.trim() !== '');
    if (lines.length === 0) return;

    const headers = lines[0].split(',');
    const rows = lines.slice(1).map(line => {
        const values = line.split(',');
        const row: Record<string, string> = {};
        headers.forEach((h, i) => {
            row[h] = values[i] || 'null';
        });
        return row;
    });

    // The reference date for age calculation is July 1, 2025.
    // In JS Date, month 6 is July.
    const referenceDate = new Date(2025, 6, 1);

    const results = rows.map(row => {
        const { FirstName, LastName } = splitName(row.Name);
        const birthDate = parseDate(row.Birthday);
        if (!birthDate) throw new Error(`Invalid birthday for ${row.Name}`);

        const deathDate = parseDate(row.Died);
        
        // Use the date of death if applicable, otherwise use the fixed reference date.
        const actualRefDate = deathDate || referenceDate;
        const age = calculateAge(birthDate, actualRefDate);

        const relatives = [];
        const familyCols = ['Father', 'Mother', 'Brother', 'Sister'];
        familyCols.forEach(col => {
            const relativeName = row[col];
            if (relativeName && relativeName !== 'null') {
                const { FirstName: rf, LastName: rl } = splitName(relativeName);
                relatives.push({
                    FirstName: rf,
                    LastName: rl,
                    Relationship: col
                });
            }
        });

        return {
            FirstName,
            LastName,
            Birthday: formatDate(birthDate),
            Age: age,
            Relatives: relatives
        };
    });

    process.stdout.write(JSON.stringify(results, null, 2) + '\n');
}

main();