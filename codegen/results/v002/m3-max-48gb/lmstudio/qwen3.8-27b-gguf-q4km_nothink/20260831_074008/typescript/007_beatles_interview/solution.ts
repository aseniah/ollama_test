import * as fs from 'fs';

const REFERENCE_DATE = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed, so 6 = July)

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

function parseCSVLine(line: string): string[] {
  // Handle quoted fields and null values
  const fields: string[] = [];
  let current = '';
  let inQuotes = false;

  for (let i = 0; i < line.length; i++) {
    const char = line[i];
    if (char === '"') {
      inQuotes = !inQuotes;
    } else if (char === ',' && !inQuotes) {
      fields.push(current.trim());
      current = '';
    } else {
      current += char;
    }
  }
  fields.push(current.trim());
  return fields;
}

function calculateAge(birthday: string, referenceDate: Date): number {
  const birthDate = new Date(birthday);
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function formatBirthday(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function splitName(fullName: string): { first: string; last: string } {
  const parts = fullName.split(' ');
  const last = parts[parts.length - 1];
  const first = parts.slice(0, -1).join(' ');
  return { first, last };
}

function processRelatives(
  father: string,
  mother: string,
  brother: string,
  sister: string
): Relative[] {
  const relatives: Relative[] = [];

  if (father && father !== 'null') {
    const { first, last } = splitName(father);
    relatives.push({ FirstName: first, LastName: last, Relationship: 'Father' });
  }

  if (mother && mother !== 'null') {
    const { first, last } = splitName(mother);
    relatives.push({ FirstName: first, LastName: last, Relationship: 'Mother' });
  }

  if (brother && brother !== 'null') {
    const { first, last } = splitName(brother);
    relatives.push({ FirstName: first, LastName: last, Relationship: 'Brother' });
  }

  if (sister && sister !== 'null') {
    const { first, last } = splitName(sister);
    relatives.push({ FirstName: first, LastName: last, Relationship: 'Sister' });
  }

  return relatives;
}

function main() {
  const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
  const lines = csvContent.trim().split('\n');
  
  // Skip header
  const dataLines = lines.slice(1);
  
  const result: Person[] = [];

  for (const line of dataLines) {
    if (!line.trim()) continue;
    
    const fields = parseCSVLine(line);
    if (fields.length < 7) continue;

    const [fullName, birthday, died, father, mother, brother, sister] = fields;
    
    const { first, last } = splitName(fullName);
    const birthDate = new Date(birthday);
    const age = calculateAge(birthDate, REFERENCE_DATE);
    const relatives = processRelatives(father, mother, brother, sister);

    result.push({
      FirstName: first,
      LastName: last,
      Birthday: formatBirthday(birthDate),
      Age: age,
      Relatives: relatives
    });
  }

  console.log(JSON.stringify(result, null, 2));
}

main();