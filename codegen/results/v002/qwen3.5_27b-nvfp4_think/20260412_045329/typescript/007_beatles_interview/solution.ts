import * as fs from 'fs';
import * as path from 'path';

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Array<{
    FirstName: string;
    LastName: string;
    Relationship: string;
  }>;
}

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null') return null;
  const parts = dateStr.split('/');
  // Format: M/D/YYYY or MM/DD/YYYY
  const month = parseInt(parts[0], 10) - 1;
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month, day);
}

function formatDate(date: Date): string {
  const yyyy = date.getFullYear();
  const mm = String(date.getMonth() + 1).padStart(2, '0');
  const dd = String(date.getDate()).padStart(2, '0');
  return `${yyyy}-${mm}-${dd}`;
}

function calculateAge(birthday: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  // Adjust if birthday hasn't occurred yet in reference year
  const monthDayRef = referenceDate.getMonth() * 100 + referenceDate.getDate();
  const monthDayBirth = birthday.getMonth() * 100 + birthday.getDate();
  if (monthDayBirth > monthDayRef) {
    age--;
  }
  return age;
}

function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(/\s+/);
  if (parts.length >= 2) {
    return {
      firstName: parts[0],
      lastName: parts[parts.length - 1]
    };
  }
  return { firstName: parts[0], lastName: '' };
}

function parseRelativeName(name: string): { firstName: string; lastName: string } | null {
  if (!name || name === 'null') return null;
  return parseName(name);
}

// Read and parse CSV
const inputPath = path.join('input', 'input.csv');
const csvContent = fs.readFileSync(inputPath, 'utf-8');
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');

const referenceDate = new Date(2025, 6, 1); // July 1, 2025
const result: Person[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  
  // Map header names to indices
  const nameIdx = headers.indexOf('Name');
  const birthdayIdx = headers.indexOf('Birthday');
  const diedIdx = headers.indexOf('Died');
  const fatherIdx = headers.indexOf('Father');
  const motherIdx = headers.indexOf('Mother');
  const brotherIdx = headers.indexOf('Brother');
  const sisterIdx = headers.indexOf('Sister');
  
  const name = parseName(values[nameIdx]);
  const birthday = parseDate(values[birthdayIdx])!;
  const died = parseDate(values[diedIdx]);
  
  // Calculate age
  const ageCalcRef = died ? died : referenceDate;
  const age = calculateAge(birthday, ageCalcRef);
  
  // Build relatives array
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  
  const fatherRel = parseRelativeName(values[fatherIdx]);
  if (fatherRel) {
    relatives.push({ FirstName: fatherRel.firstName, LastName: fatherRel.lastName, Relationship: 'Father' });
  }
  
  const motherRel = parseRelativeName(values[motherIdx]);
  if (motherRel) {
    relatives.push({ FirstName: motherRel.firstName, LastName: motherRel.lastName, Relationship: 'Mother' });
  }
  
  const brotherRel = parseRelativeName(values[brotherIdx]);
  if (brotherRel) {
    relatives.push({ FirstName: brotherRel.firstName, LastName: brotherRel.lastName, Relationship: 'Brother' });
  }
  
  const sisterRel = parseRelativeName(values[sisterIdx]);
  if (sisterRel) {
    relatives.push({ FirstName: sisterRel.firstName, LastName: sisterRel.lastName, Relationship: 'Sister' });
  }
  
  result.push({
    FirstName: name.firstName,
    LastName: name.lastName,
    Birthday: formatDate(birthday),
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result));