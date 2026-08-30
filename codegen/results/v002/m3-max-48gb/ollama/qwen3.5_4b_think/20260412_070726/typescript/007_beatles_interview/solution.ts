import * as fs from 'fs';

const inputPath = 'input/input.csv';
const expectedPath = 'input/expected_format.json';

// Read input CSV
const csvContent = fs.readFileSync(inputPath, 'utf-8');

// Parse CSV
const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');

const personFields = {};
// Determine field order
for (let i = 0; i < header.length; i++) {
  personFields[header[i]] = i;
}

// Parse relatives from name field
function parseNameParts(name: string): { firstName: string; lastName: string } {
  // John Winston Lennon -> John Lennon
  const lastDotIndex = name.lastIndexOf(' ');
  const lastName = name.substring(lastDotIndex + 1);
  const firstName = name.substring(0, lastDotIndex);
  return { firstName, lastName };
}

// Convert M/D/YYYY to Y-M-D
function parseDate(dateStr: string): string {
  const parts = dateStr.split('/');
  const month = String(parts[0]).padStart(2, '0');
  const day = String(parts[1]).padStart(2, '0');
  const year = parts[2];
  return `${year}-${month}-${day}`;
}

// Calculate age as of July 1, 2025
function calculateAge(birthDate: Date): number {
  const targetDate = new Date('2025-07-01');
  let age = targetDate.getFullYear() - birthDate.getFullYear();
  
  // If birthday hasn't occurred yet in the current year, subtract 1
  if (birthDate.getMonth() > 6 || (birthDate.getMonth() === 6 && birthDate.getDate() > 1)) {
    age--;
  }
  
  return age;
}

// Parse full name for relatives (e.g., "Alfred Lennon" -> "Alfred", "Lennon")
function parseRelativeName(name: string | null): { firstName: string; lastName: string } | null {
  if (!name) return null;
  const parts = name.split(' ');
  if (parts.length >= 2) {
    return { firstName: parts[0], lastName: parts.slice(1).join(' ') };
  }
  return { firstName: parts[0], lastName: '' };
}

const persons: any[] = [];

// Process data rows
for (let i = 1; i < lines.length; i++) {
  const fields = lines[i].split(',');
  
  const firstName = fields[0].split(' ')[0];
  const lastName = fields[0].split(' ').slice(1).join(' ');
  
  const birthday = parseDate(fields[1]);
  const birthDate = new Date(birthday + 'T00:00:00');
  
  const age = calculateAge(birthDate);
  
  const relatives: any[] = [];
  
  // Father
  const fatherName = fields[3] || null;
  if (fatherName) {
    relatives.push(parseRelativeName(fatherName));
  }
  
  // Mother
  const motherName = fields[4] || null;
  if (motherName) {
    relatives.push(parseRelativeName(motherName));
  }
  
  // Brother
  const brotherName = fields[5] || null;
  if (brotherName) {
    relatives.push(parseRelativeName(brotherName));
  }
  
  // Sister
  const sisterName = fields[6] || null;
  if (sisterName) {
    relatives.push(parseRelativeName(sisterName));
  }
  
  // Add relationship labels
  relatives[0] && (relatives[0].Relationship = 'Father');
  relatives[1] && (relatives[1].Relationship = 'Mother');
  relatives[2] && (relatives[2].Relationship = 'Brother');
  relatives[3] && (relatives[3].Relationship = 'Sister');
  
  persons.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthday,
    Age: age,
    Relatives: relatives
  });
}

// Output JSON to stdout
console.log(JSON.stringify(persons, null, 2));