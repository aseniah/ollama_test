import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const csvPath = path.join('input', 'input.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8');

// Parse CSV
const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');
const dataLines = lines.slice(1);

// Parse birthday to YYYY-MM-DD format
function parseBirthday(dateStr: string): string {
  if (!dateStr || dateStr === 'null') return '';
  const parts = dateStr.split('/');
  const month = parts[0].padStart(2, '0');
  const day = parts[1].padStart(2, '0');
  const year = parts[2];
  return `${year}-${month}-${day}`;
}

// Calculate age as of July 1, 2025
function calculateAge(birthday: string): number {
  if (!birthday) return 0;
  const birthDate = new Date(birthday + 'T00:00:00');
  const referenceDate = new Date('2025-07-01T00:00:00');
  
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

// Parse name into first and last
function parseName(name: string): { firstName: string, lastName: string } {
  const parts = name.trim().split(' ');
  const firstName = parts[0];
  const lastName = parts[parts.length - 1];
  return { firstName, lastName };
}

// Build relatives array
function buildRelatives(row: string[]): any[] {
  const relatives: any[] = [];
  
  const father = row[3];
  const mother = row[4];
  const brother = row[5];
  const sister = row[6];
  
  if (father && father !== 'null') {
    const parts = father.trim().split(' ');
    relatives.push({
      FirstName: parts[0],
      LastName: parts[parts.length - 1],
      Relationship: 'Father'
    });
  }
  
  if (mother && mother !== 'null') {
    const parts = mother.trim().split(' ');
    relatives.push({
      FirstName: parts[0],
      LastName: parts[parts.length - 1],
      Relationship: 'Mother'
    });
  }
  
  if (brother && brother !== 'null') {
    const parts = brother.trim().split(' ');
    relatives.push({
      FirstName: parts[0],
      LastName: parts[parts.length - 1],
      Relationship: 'Brother'
    });
  }
  
  if (sister && sister !== 'null') {
    const parts = sister.trim().split(' ');
    relatives.push({
      FirstName: parts[0],
      LastName: parts[parts.length - 1],
      Relationship: 'Sister'
    });
  }
  
  return relatives;
}

// Process each person
const result: any[] = [];

for (const line of dataLines) {
  const row = line.split(',');
  
  const { firstName, lastName } = parseName(row[0]);
  const birthday = parseBirthday(row[1]);
  const age = calculateAge(birthday);
  const relatives = buildRelatives(row);
  
  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthday,
    Age: age,
    Relatives: relatives
  });
}

// Output JSON
console.log(JSON.stringify(result, null, 2));