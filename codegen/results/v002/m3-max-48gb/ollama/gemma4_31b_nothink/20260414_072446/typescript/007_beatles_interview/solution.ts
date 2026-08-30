import * as fs from 'fs';
import * as path from 'path';

/**
 * Task: Read CSV, transform to JSON based on expected format.
 * Age calculation date: July 1, 2025.
 */

interface Person {
    FirstName: string;
    LastName: string;
    Birthday: string;
    Age: number;
    Relatives: Relative[];
}

interface Relative {
    FirstName: string;
    LastName: string;
    Relationship: string;
}

function calculateAge(birthdayStr: string, referenceDate: Date): number {
    // Input format in CSV is M/D/YYYY
    const parts = birthdayStr.split('/');
    const birthDate = new Date(
        parseInt(parts[2]),
        parseInt(parts[0]) - 1,
        parseInt(parts[1])
    );

    let age = referenceDate.getFullYear() - birthDate.getFullYear();
    const m = referenceDate.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && referenceDate.getDate() < birthDate.getDate())) {
        age--;
    }
    return age;
}

function formatIsoDate(dateStr: string): string {
    const parts = dateStr.split('/');
    const y = parts[2];
    const m = parts[0].padStart(2, '0');
    const d = parts[1].padStart(2, '0');
    return `${y}-${m}-${d}`;
}

function splitFullName(fullName: string): { first: string, last: string } {
    const parts = fullName.trim().split(' ');
    if (parts.length === 0) return { first: '', last: '' };
    if (parts.length === 1) return { first: parts[0], last: '' };
    
    // For the primary person, they might have middle names. 
    // Based on expected_format.json, "John Winston Lennon" becomes "John" and "Lennon".
    return {
        first: parts[0],
        last: parts[parts.length - 1]
    };
}

function solve() {
    const inputPath = path.join('input', 'input.csv');
    const content = fs.readFileSync(inputPath, 'utf8');
    const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
    
    const headers = lines[0].split(',');
    const rows = lines.slice(1);
    const refDate = new Date(2025, 6, 1); // July 1, 2025 (Month is 0-indexed)

    const result: Person[] = rows.map(row => {
        const values = row.split(',');
        const data: Record<string, string> = {};
        headers.forEach((header, index) => {
            data[header] = values[index];
        });

        const nameParts = splitFullName(data['Name']);
        const birthday = data['Birthday'];

        const person: Person = {
            FirstName: nameParts.first,
            LastName: nameParts.last,
            Birthday: formatIsoDate(birthday),
            Age: calculateAge(birthday, refDate),
            Relatives: []
        };

        const relativeTypes = ['Father', 'Mother', 'Brother', 'Sister'];
        relativeTypes.forEach(relType => {
            const relName = data[relType];
            if (relName && relName !== 'null') {
                const relParts = splitFullName(relName);
                person.Relatives.push({
                    FirstName: relParts.first,
                    LastName: relParts.last,
                    Relationship: relType
                });
            }
        });

        return person;
    });

    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

solve();