import * as fs from 'fs';
import * as path from 'path';

const referenceDate = new Date('2025-07-01');
const inputPath = './input/input.csv';

// Parse birthday string (MM/DD/YYYY) to Date object
function parseDate(dateStr: string): Date {
  const parts = dateStr.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
}

// Calculate age as of July 1, 2025
function calculateAge(birthdate: Date): number {
  const age = referenceDate.getFullYear() - birthdate.getFullYear();
  const currentMonth = referenceDate.getMonth() + 1;
  const birthMonth = birthdate.getMonth() + 1;
  
  if (currentMonth < birthMonth || (currentMonth === birthMonth && referenceDate.getDate() < birthdate.getDate())) {
    age--;
  }
  return age;
}

// Parse a full name (first space = last, last word = last name)
function parseFullName(nameStr: string): { first: string, last: string } {
  const parts = nameStr.trim().split(/[\s]+/);
  // If there are multiple parts (like "John Winston Lennon"), use first and last
  if (parts.length >= 2) {
    return { first: parts[0], last: parts[parts.length - 1] };
  } else {
    return { first: parts[0], last: '' };
  }
}

// Read and parse CSV
const csvContent = fs.readFileSync(inputPath, 'utf-8');
const lines = csvContent.trim().split('\n');

// Parse header
const headers = lines[0].split(',').map(h => h.trim().toLowerCase());
const headerIndex = {
  firstName: headers.findIndex(h => h.includes('birthday') || h.includes('name')),
  birthday: headers.findIndex(h => h.includes('birthday')),
  father: headers.findIndex(h => h.includes('father')),
  mother: headers.findIndex(h => h.includes('mother')),
  brother: headers.findIndex(h => h.includes('brother')),
  sister: headers.findIndex(h => h.includes('sister'))
};

const dataLines = lines.slice(1);

// Process each data row
const results: Array<{
  firstName: string;
  lastName: string;
  birthday: string;
  age: number;
  relatives: Array<{
    firstName: string;
    lastName: string;
    relationship: string;
  }>
}> = [];

for (const line of dataLines) {
  const fields = line.split(',');
  
  // Parse name (first and last word)
  const nameStr = fields[0].trim();
  const { first, last } = parseFullName(nameStr);
  
  // Parse birthday
  const birthdayStr = fields[1].trim();
  const birthday = parseDate(birthdayStr);
  
  // Calculate age
  const age = calculateAge(birthday);
  
  // Parse relatives
  const relatives: Array<{ firstName: string; lastName: string; relationship: string }> = [];
  
  const fatherStr = fields[3]?.trim();
  const motherStr = fields[4]?.trim();
  const brotherStr = fields[5]?.trim();
  const sisterStr = fields[6]?.trim();
  
  // Helper to parse relative name
  function parseRelative(relStr: string, relationship: string): { firstName: string; lastName: string } | null {
    if (!relStr || relStr === 'null') {
      return null;
    }
    const parts = relStr.split(/[\s]+/);
    if (parts.length >= 2) {
      return { first: parts[0], last: parts[parts.length - 1] };
    } else {
      return { first: parts[0], last: '' };
    }
  }
  
  const father = parseRelative(fatherStr || 'null', 'Father');
  const mother = parseRelative(motherStr || 'null', 'Mother');
  const brother = parseRelative(brotherStr || 'null', 'Brother');
  const sister = parseRelative(sisterStr || 'null', 'Sister');
  
  if (father) relatives.push({ firstName: father.first, lastName: father.last, relationship: 'Father' });
  if (mother) relatives.push({ firstName: mother.first, lastName: mother.last, relationship: 'Mother' });
  if (brother) relatives.push({ firstName: brother.first, lastName: brother.last, relationship: 'Brother' });
  if (sister) relatives.push({ firstName: sister.first, lastName: sister.last, relationship: 'Sister' });
  
  results.push({
    firstName: first,
    lastName: last,
    birthday: birthdayStr,
    age: age,
    relatives
  });
}

// Output JSON array to stdout
console.log(JSON.stringify(results, null, 2));