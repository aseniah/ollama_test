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

// Read the CSV file
const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
const lines = csvContent.trim().split('\n');

// Parse headers
const headers = lines[0].split(',');

// Date for age calculation: July 1, 2025
const referenceDate = new Date(2025, 6, 1); // Month is 0-indexed, so 6 = July

function splitName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  const firstName = parts[0];
  const lastName = parts.length > 1 ? parts.slice(1).join(' ') : '';
  return { firstName, lastName };
}

function parseDate(dateStr: string): Date {
  // Format: M/D/YYYY or MM/DD/YYYY
  const parts = dateStr.split('/');
  const month = parseInt(parts[0], 10) - 1;
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month, day);
}

function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function calculateAge(birthDate: Date, reference: Date): number {
  let age = reference.getFullYear() - birthDate.getFullYear();
  const monthDiff = reference.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && reference.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

const result: Person[] = [];

// Process each data row
for (let i = 1; i < lines.length; i++) {
  const line = lines[i].trim();
  if (!line) continue;
  
  const fields = line.split(',');
  
  const fullName = fields[0].trim();
  const birthdayStr = fields[1].trim();
  const diedStr = fields[2].trim();
  const fatherStr = fields[3].trim();
  const motherStr = fields[4].trim();
  const brotherStr = fields[5].trim();
  const sisterStr = fields[6].trim();
  
  const { firstName, lastName } = splitName(fullName);
  const birthDate = parseDate(birthdayStr);
  const age = calculateAge(birthDate, referenceDate);
  
  const relatives: Relative[] = [];
  
  // Add relatives (skip if "null")
  if (fatherStr !== 'null') {
    const { firstName: fFirst, lastName: fLast } = splitName(fatherStr);
    relatives.push({
      FirstName: fFirst,
      LastName: fLast,
      Relationship: 'Father'
    });
  }
  
  if (motherStr !== 'null') {
    const { firstName: mFirst, lastName: mLast } = splitName(motherStr);
    relatives.push({
      FirstName: mFirst,
      LastName: mLast,
      Relationship: 'Mother'
    });
  }
  
  if (brotherStr !== 'null') {
    const { firstName: bFirst, lastName: bLast } = splitName(brotherStr);
    relatives.push({
      FirstName: bFirst,
      LastName: bLast,
      Relationship: 'Brother'
    });
  }
  
  if (sisterStr !== 'null') {
    const { firstName: sFirst, lastName: sLast } = splitName(sisterStr);
    relatives.push({
      FirstName: sFirst,
      LastName: sLast,
      Relationship: 'Sister'
    });
  }
  
  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: formatDate(birthDate),
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 1));