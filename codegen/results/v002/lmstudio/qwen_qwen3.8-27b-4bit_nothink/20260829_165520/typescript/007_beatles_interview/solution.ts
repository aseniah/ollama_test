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

const csvContent = fs.readFileSync('input/input.csv', 'utf-8').trim();
const lines = csvContent.split('\n');
const headers = lines[0].split(',');
const dataLines = lines.slice(1);

// Reference date: July 1, 2025
const refDate = new Date(2025, 6, 1); // Month is 0-indexed, so 6 = July

// Map from relationship column names to relationship labels
const relationshipMap: { [key: string]: string } = {
    'Father': 'Father',
    'Mother': 'Mother',
    'Brother': 'Brother',
    'Sister': 'Sister'
};

function parseDate(dateStr: string): Date {
    // Format is M/D/YYYY
    const parts = dateStr.split('/');
    const month = parseInt(parts[0], 10) - 1; // 0-indexed
    const day = parseInt(parts[1], 10);
    const year = parseInt(parts[2], 10);
    return new Date(year, month, day);
}

function formatISODate(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
}

function calculateAge(birthday: Date, died?: Date, refDate: Date = new Date(2025, 6, 1)): number {
    let age: number;
    const deathDate = died ?? refDate;
    
    age = deathDate.getFullYear() - birthday.getFullYear();
    
    // Check if birthday hasn't occurred yet in the death/reference year
    if (deathDate.getMonth() < birthday.getMonth() ||
        (deathDate.getMonth() === birthday.getMonth() && deathDate.getDate() < birthday.getDate())) {
        age--;
    }
    
    return age;
}

function parseName(fullName: string): { first: string; last: string } {
    const parts = fullName.trim().split(/\s+/);
    if (parts.length === 1) {
        return { first: parts[0], last: '' };
    } else if (parts.length === 2) {
        return { first: parts[0], last: parts[1] };
    } else {
        // Take first name and last name (last part), ignore middle names
        const first = parts[0];
        const last = parts[parts.length - 1];
        return { first: first, last: last };
    }
}

function parseRelative(fullName: string, relationship: string): Relative | null {
    if (!fullName || fullName.trim() === 'null' || fullName.trim() === '') {
        return null;
    }
    const { first, last } = parseName(fullName);
    return {
        FirstName: first,
        LastName: last,
        Relationship: relationship
    };
}

const result: Person[] = [];

for (const line of dataLines) {
    const cols = line.split(',');
    const name = cols[0].trim();
    const birthdayStr = cols[1].trim();
    const diedStr = cols[2].trim();
    const fatherStr = cols[3].trim();
    const motherStr = cols[4].trim();
    const brotherStr = cols[5].trim();
    const sisterStr = cols[6].trim();

    const { first, last } = parseName(name);
    const birthday = parseDate(birthdayStr);
    const died = diedStr !== 'null' ? parseDate(diedStr) : undefined;
    
    const age = calculateAge(birthday, died);
    
    const relatives: Relative[] = [];
    
    const fatherRel = parseRelative(fatherStr, 'Father');
    if (fatherRel) relatives.push(fatherRel);
    
    const motherRel = parseRelative(motherStr, 'Mother');
    if (motherRel) relatives.push(motherRel);
    
    const brotherRel = parseRelative(brotherStr, 'Brother');
    if (brotherRel) relatives.push(brotherRel);
    
    const sisterRel = parseRelative(sisterStr, 'Sister');
    if (sisterRel) relatives.push(sisterRel);
    
    result.push({
        FirstName: first,
        LastName: last,
        Birthday: formatISODate(birthday),
        Age: age,
        Relatives: relatives
    });
}

console.log(JSON.stringify(result, null, 2));