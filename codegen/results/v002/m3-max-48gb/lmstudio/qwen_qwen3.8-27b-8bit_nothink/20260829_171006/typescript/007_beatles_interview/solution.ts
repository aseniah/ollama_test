import * as fs from 'fs';

function calculateAge(birthday: Date, asOf: Date): number {
  let age = asOf.getFullYear() - birthday.getFullYear();
  const monthDiff = asOf.getMonth() - birthday.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && asOf.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function parseDate(str: string): Date {
  // Format: MM/DD/YYYY
  const parts = str.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
}

function formatDate(date: Date): string {
  const year = date.getFullYear().toString().padStart(4, '0');
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function splitName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  const lastName = parts[parts.length - 1];
  const firstName = parts[0];
  return { firstName, lastName };
}

function buildRelatives(
  father: string,
  mother: string,
  brother: string,
  sister: string
): Array<{ FirstName: string; LastName: string; Relationship: string }> {
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];

  if (father && father !== 'null') {
    const { firstName, lastName } = splitName(father);
    relatives.push({ FirstName: firstName, LastName: lastName, Relationship: 'Father' });
  }
  if (mother && mother !== 'null') {
    const { firstName, lastName } = splitName(mother);
    relatives.push({ FirstName: firstName, LastName: lastName, Relationship: 'Mother' });
  }
  if (brother && brother !== 'null') {
    const { firstName, lastName } = splitName(brother);
    relatives.push({ FirstName: firstName, LastName: lastName, Relationship: 'Brother' });
  }
  if (sister && sister !== 'null') {
    const { firstName, lastName } = splitName(sister);
    relatives.push({ FirstName: firstName, LastName: lastName, Relationship: 'Sister' });
  }

  return relatives;
}

const csvContent = fs.readFileSync('input/input.csv', 'utf-8').trim();
const lines = csvContent.split('\n');

const header = lines[0].split(',');
const dataLines = lines.slice(1);

const result: Array<{
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Array<{ FirstName: string; LastName: string; Relationship: string }>;
}> = [];

const referenceDate = new Date(2025, 6, 1); // July 1, 2025

for (const line of dataLines) {
  const fields = line.split(',');
  const name = fields[0];
  const birthdayStr = fields[1];
  const diedStr = fields[2];
  const father = fields[3];
  const mother = fields[4];
  const brother = fields[5];
  const sister = fields[6];

  const { firstName, lastName } = splitName(name);
  const birthday = parseDate(birthdayStr);
  const birthdayFormatted = formatDate(birthday);

  let age: number;
  if (diedStr && diedStr !== 'null') {
    const deathDate = parseDate(diedStr);
    age = calculateAge(birthday, deathDate);
  } else {
    age = calculateAge(birthday, referenceDate);
  }

  const relatives = buildRelatives(father, mother, brother, sister);

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthdayFormatted,
    Age: age,
    Relatives: relatives,
  });
}

console.log(JSON.stringify(result, null, 2));