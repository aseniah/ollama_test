import * as fs from 'fs';

const inputPath = 'input/input.csv';
const outputPath = 'input/expected_format.json';

// Read the input CSV
const csvContent = fs.readFileSync(inputPath, 'utf-8');
const lines = csvContent.trim().split('\n');

// Skip header
const dataLines = lines.slice(1);

// Reference date for age calculation
const refDate = new Date('2025-07-01');

function calculateAge(birthdayStr: string): number {
  // Parse various date formats: MM/DD/YYYY or M/D/YYYY
  const parts = birthdayStr.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);

  const birthDate = new Date(year, month - 1, day);

  let age = refDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = refDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && refDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(' ');
  if (parts.length >= 2) {
    const lastName = parts.pop()!;
    const firstName = parts.join(' ');
    return { firstName, lastName };
  }
  return { firstName: name.trim(), lastName: '' };
}

function parseRelative(name: string, relationship: string): { firstName: string; lastName: string; relationship: string } | null {
  if (!name || name === 'null') return null;
  const { firstName, lastName } = parseName(name);
  return { firstName, lastName, relationship };
}

const result: any[] = [];

for (const line of dataLines) {
  const parts = line.split(',');
  if (parts.length < 2) continue;

  const name = parts[0];
  const birthday = parts[1];
  const died = parts[2]; // Not used in output, but present in CSV
  const fatherName = parts[3];
  const motherName = parts[4];
  const brotherName = parts[5];
  const sisterName = parts[6];

  const { firstName, lastName } = parseName(name);
  const age = calculateAge(birthday);

  const relatives: { firstName: string; lastName: string; relationship: string }[] = [];

  const father = parseRelative(fatherName, 'Father');
  if (father) relatives.push(father);

  const mother = parseRelative(motherName, 'Mother');
  if (mother) relatives.push(mother);

  const brother = parseRelative(brotherName, 'Brother');
  if (brother) relatives.push(brother);

  const sister = parseRelative(sisterName, 'Sister');
  if (sister) relatives.push(sister);

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthday,
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 2));