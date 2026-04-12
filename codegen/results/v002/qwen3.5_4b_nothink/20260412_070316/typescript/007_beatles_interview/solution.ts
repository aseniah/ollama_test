import * as fs from 'fs';
import * as path from 'path';

// Reference date for age calculation: July 1, 2025
const REFERENCE_DATE = new Date('2025-07-01');

// Helper to parse date from string "M/D/YYYY"
function parseDateString(dateStr: string): Date {
    const [month, day, year] = dateStr.split('/').map(Number);
    return new Date(year, month - 1, day);
}

// Helper to calculate age as of July 1, 2025
function calculateAge(birthDate: Date): number {
    let age = REFERENCE_DATE.getFullYear() - birthDate.getFullYear();
    // If birthday hasn't occurred yet this year, subtract 1
    if (
        birthDate.getMonth() > REFERENCE_DATE.getMonth() ||
        (birthDate.getMonth() === REFERENCE_DATE.getMonth() && birthDate.getDate() > REFERENCE_DATE.getDate())
    ) {
        age -= 1;
    }
    return age;
}

// Helper to format date as "YYYY-MM-DD"
function formatDate(date: Date): string {
    return date.toISOString().split('T')[0];
}

// Main logic
const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
const lines = csvContent.trim().split('\n');

// Skip header line
if (lines.length <= 1) {
    console.log(JSON.stringify([]));
    process.exit(0);
}

const results: Array<{
    FirstName: string;
    LastName: string;
    Birthday: string;
    Age: number;
    Relatives: Array<{
        FirstName: string;
        LastName: string;
        Relationship: string;
    }>;
}> = [];

for (let i = 1; i < lines.length; i++) {
    const fields = lines[i].split(',');
    
    // Basic validation
    if (fields.length < 5) {
        continue;
    }

    const name = fields[0].trim(); // First name + last name combined (e.g., "John Winston Lennon")
    const birthdayStr = fields[1].trim(); // "M/D/YYYY"
    const fatherStr = fields[3]?.trim();
    const motherStr = fields[4]?.trim();
    const brotherStr = fields[5]?.trim();
    const sisterStr = fields[6]?.trim();

    // Parse name
    // Format: "First Middle Last" -> split by space, last part is surname, rest are given names
    const nameParts = name.split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];
    const middleName = nameParts.slice(1, -1).join(' ');

    // Parse birthday
    const birthDate = parseDateString(birthdayStr);
    const age = calculateAge(birthDate);

    // Parse relatives
    const relatives: Array<{
        FirstName: string;
        LastName: string;
        Relationship: string;
    }> = [];

    // Father
    if (fatherStr) {
        const fatherParts = fatherStr.split(' ');
        const fatherName = fatherParts[0];
        const fatherSurname = fatherParts[fatherParts.length - 1];
        relatives.push({
            FirstName: fatherName,
            LastName: fatherSurname,
            Relationship: 'Father'
        });
    }

    // Mother
    if (motherStr) {
        const motherParts = motherStr.split(' ');
        const motherName = motherParts[0];
        const motherSurname = motherParts[motherParts.length - 1];
        relatives.push({
            FirstName: motherName,
            LastName: motherSurname,
            Relationship: 'Mother'
        });
    }

    // Brother
    if (brotherStr) {
        const brotherParts = brotherStr.split(' ');
        const brotherName = brotherParts[0];
        const brotherSurname = brotherParts[brotherParts.length - 1];
        relatives.push({
            FirstName: brotherName,
            LastName: brotherSurname,
            Relationship: 'Brother'
        });
    }

    // Sister
    if (sisterStr) {
        const sisterParts = sisterStr.split(' ');
        const sisterName = sisterParts[0];
        const sisterSurname = sisterParts[sisterParts.length - 1];
        relatives.push({
            FirstName: sisterName,
            LastName: sisterSurname,
            Relationship: 'Sister'
        });
    }

    results.push({
        FirstName,
        LastName,
        Birthday: formatDate(birthDate),
        Age,
        relatives
    });
}

// Output JSON array
console.log(JSON.stringify(results, null, 2));