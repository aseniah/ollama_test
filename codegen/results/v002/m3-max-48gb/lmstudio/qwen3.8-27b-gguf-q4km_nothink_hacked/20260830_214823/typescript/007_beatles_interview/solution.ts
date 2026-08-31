import * as fs from 'fs';

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

function parseName(fullName: string): { first: string; last: string } {
    const parts = fullName.trim().split(' ');
    const first = parts[0];
    const last = parts[parts.length - 1];
    return { first, last };
}

function parseDate(dateStr: string): Date {
    const [month, day, year] = dateStr.split('/').map(Number);
    return new Date(year, month - 1, day);
}

function calculateAge(birthday: string, asOf: Date): number {
    const birth = parseDate(birthday);
    let age = asOf.getFullYear() - birth.getFullYear();
    const monthDiff = asOf.getMonth() - birth.getMonth();
    if (monthDiff < 0 || (monthDiff === 0 && asOf.getDate() < birth.getDate())) {
        age--;
    }
    return age;
}

function main() {
    const asOfDate = new Date(2025, 6, 1); // July 1, 2025

    const csvContent = fs.readFileSync('input/input.csv', 'utf8');
    const lines = csvContent.trim().split('\n');
    const headers = lines[0].split(',');
    const dataRows = lines.slice(1);

    const result: Person[] = [];

    for (const row of dataRows) {
        const values = row.split(',');
        const record: { [key: string]: string } = {};
        headers.forEach((header, i) => {
            record[header] = values[i];
        });

        const { first, last } = parseName(record['Name']);
        
        const relatives: Relative[] = [];

        const relativeFields: { column: string; relationship: string }[] = [
            { column: 'Father', relationship: 'Father' },
            { column: 'Mother', relationship: 'Mother' },
            { column: 'Brother', relationship: 'Brother' },
            { column: 'Sister', relationship: 'Sister' }
        ];

        for (const { column, relationship } of relativeFields) {
            const value = record[column].trim();
            if (value && value.toLowerCase() !== 'null') {
                const { first: relFirst, last: relLast } = parseName(value);
                relatives.push({
                    FirstName: relFirst,
                    LastName: relLast,
                    Relationship: relationship
                });
            }
        }

        result.push({
            FirstName: first,
            LastName: last,
            Birthday: record['Birthday'],
            Age: calculateAge(record['Birthday'], asOfDate),
            Relatives: relatives
        });
    }

    console.log(JSON.stringify(result, null, 2));
}

main();