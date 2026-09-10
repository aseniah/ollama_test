import * as fs from 'fs';
import * as path from 'path';

/**
 * Calculates the age based on a birthday and a reference date (July 1, 2025).
 * If the person died, the age is calculated as of the date of death.
 * If they are alive, it is calculated as of July 1, 2025.
 */
function calculateAge(birthdayStr: string, diedStr: string | null): number {
    const refDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)
    
    // Input format is M/D/YYYY
    const parseDate = (str: string) => {
        const [m, d, y] = str.split('/').map(Number);
        return new Date(y, m - 1, d);
    };

    const birthDate = parseDate(birthdayStr);
    const endDate = diedStr ? parseDate(diedStr) : refDate;

    let age = endDate.getFullYear() - birthDate.getFullYear();
    const m = endDate.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && endDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

function splitFullName(name: string) {
    const parts = name.trim().split(/\s+/);
    if (parts.length === 0) return { first: '', last: '' };
    if (parts.length === 1) return { first: parts[0], last: '' };
    // Assuming format: First [Middle] Last
    return { first: parts[0], last: parts[parts.length - 1] };
}

function parseCsv() {
    const filePath = path.join('input', 'input.csv');
    const content = fs.readFileSync(filePath, 'utf-8');
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    const headers = lines[0].split(',');
    const dataLines = lines.slice(1);

    const result = dataLines.map(line => {
        const values = line.split(',');
        const row: Record<string, string> = {};
        headers.forEach((header, i) => {
            row[header] = values[i] ? values[i].trim() : 'null';
        });

        const nameInfo = splitFullName(row['Name']);
        
        const person = {
            FirstName: nameInfo.first,
            LastName: nameInfo.last,
            Birthday: formatToISO(row['Birthday']),
            Age: calculateAge(row['Birthday'], row['Died'] === 'null' ? null : row['Died']),
            Relatives: [] as any[]
        };

        const relativeCols = ['Father', 'Mother', 'Brother', 'Sister'];
        relativeCols.forEach(col => {
            const relName = row[col];
            if (relName && relName !== 'null') {
                const relInfo = splitFullName(relName);
                person.Relatives.push({
                    FirstName: relInfo.first,
                    LastName: relInfo.last,
                    Relationship: col
                });
            }
        });

        return person;
    });

    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

function formatToISO(dateStr: string): string {
    // Convert M/D/YYYY to YYYY-MM-DD
    const [m, d, y] = dateStr.split('/');
    const month = m.padStart(2, '0');
    const day = d.padStart(2, '0');
    return `${y}-${month}-${day}`;
}

parseCsv();