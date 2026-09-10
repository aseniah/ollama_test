import * as fs from 'fs';
import * as path from 'path';

interface Relative {
    FirstName: string;
    LastName: string;
    Relationship: string;
}

interface Person {
    FirstName: string;
    LastName: string;
    Birthday: string;
    Age: number;
    Relatives: Relative[];
}

function parseDate(dateStr: string): Date {
    // Input format is M/D/YYYY
    const [m, d, y] = dateStr.split('/').map(Number);
    return new Date(y, m - 1, d);
}

function formatDate(date: Date): string {
    const y = date.getFullYear();
    const m = String(date.getMonth() + 1).padStart(2, '0');
    const d = String(date.getDate()).padStart(2, '0');
    return `${y}-${m}-${d}`;
}

function calculateAge(birthdayStr: string, diedStr: string | null): number {
    const birthDate = parseDate(birthdayStr);
    const targetDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)
    
    let effectiveEndDate = targetDate;
    if (diedStr && diedStr !== 'null') {
        effectiveEndDate = parseDate(diedStr);
    }

    let age = effectiveEndDate.getFullYear() - birthDate.getFullYear();
    const mDiff = effectiveEndDate.getMonth() - birthDate.getMonth();
    if (mDiff < 0 || (mDiff === 0 && effectiveEndDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

function splitName(fullName: string): { first: string, last: string } {
    if (!fullName || fullName === 'null') return { first: '', last: '' };
    const parts = fullName.trim().split(' ');
    if (parts.length === 1) return { first: parts[0], last: '' };
    
    // The expected format seems to take the first word as FirstName 
    // and the last word as LastName for the main person.
    // For relatives, it's simple First Last.
    return {
        first: parts[0],
        last: parts[parts.length - 1]
    };
}

function main() {
    const inputPath = path.join('input', 'input.csv');
    const content = fs.readFileSync(inputPath, 'utf8');
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    const headers = lines[0].split(',');
    const dataLines = lines.slice(1);
    
    const result: Person[] = [];

    for (const line of dataLines) {
        const values = line.split(',');
        const row: any = {};
        headers.forEach((header, i) => {
            row[header] = values[i];
        });

        const nameInfo = splitName(row.Name);
        const birthday = row.Birthday;
        const died = row.Died === 'null' ? null : row.Died;
        
        const person: Person = {
            FirstName: nameInfo.first,
            LastName: nameInfo.last,
            Birthday: formatDate(parseDate(birthday)),
            Age: calculateAge(birthday, died),
            Relatives: []
        };

        const relativeCols = ['Father', 'Mother', 'Brother', 'Sister'];
        relativeCols.forEach(col => {
            const relName = row[col];
            if (relName && relName !== 'null') {
                const relInfo = splitName(relName);
                person.Relatives.push({
                    FirstName: relInfo.first,
                    LastName: relInfo.last,
                    Relationship: col
                });
            }
        });

        result.push(person);
    }

    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();