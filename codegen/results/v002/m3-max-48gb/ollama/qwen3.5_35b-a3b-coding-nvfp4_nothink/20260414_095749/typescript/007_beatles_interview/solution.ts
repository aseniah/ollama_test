import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join(__dirname, 'input', 'input.csv');
const csvContent = fs.readFileSync(inputPath, 'utf-8');

const lines = csvContent.trim().split('\n');
const header = lines[0].split(',');
const dataLines = lines.slice(1);

const referenceDate = new Date('2025-07-01');

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null') return null;
  // Handle formats like "10/9/1940" (MM/D/YYYY)
  const parts = dateStr.split('/');
  if (parts.length !== 3) return null;
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  if (isNaN(month) || isNaN(day) || isNaN(year)) return null;
  return new Date(year, month - 1, day);
}

function calculateAge(birthday: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthday.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function extractFirstName(fullName: string): string {
  // Handle "FirstName MiddleName LastName" -> take first word
  const parts = fullName.trim().split(/\s+/);
  return parts[0] || '';
}

function extractLastName(fullName: string): string {
  const parts = fullName.trim().split(/\s+/);
  return parts[parts.length - 1] || '';
}

function parseRelative(field: string): { firstName: string; lastName: string } | null {
  if (!field || field === 'null' || field === '') return null;
  return {
    firstName: extractFirstName(field),
    lastName: extractLastName(field)
  };
}

const result: any[] = [];

dataLines.forEach(line => {
  const values = line.split(',');
  const name = values[0];
  const birthdayStr = values[1];
  const diedStr = values[2]; // Not used in output, but present
  const fatherStr = values[3];
  const motherStr = values[4];
  const brotherStr = values[5];
  const sisterStr = values[6];

  const birthday = parseDate(birthdayStr);
  if (!birthday) return;

  const person: any = {
    FirstName: extractFirstName(name),
    LastName: extractLastName(name),
    Birthday: birthday.toISOString().split('T')[0],
    Age: calculateAge(birthday),
    Relatives: []
  };

  const father = parseRelative(fatherStr);
  if (father) {
    person.Relatives.push({
      FirstName: father.firstName,
      LastName: father.lastName,
      Relationship: 'Father'
    });
  }

  const mother = parseRelative(motherStr);
  if (mother) {
    person.Relatives.push({
      FirstName: mother.firstName,
      LastName: mother.lastName,
      Relationship: 'Mother'
    });
  }

  const brother = parseRelative(brotherStr);
  if (brother) {
    person.Relatives.push({
      FirstName: brother.firstName,
      LastName: brother.lastName,
      Relationship: 'Brother'
    });
  }

  const sister = parseRelative(sisterStr);
  if (sister) {
    person.Relatives.push({
      FirstName: sister.firstName,
      LastName: sister.lastName,
      Relationship: 'Sister'
    });
  }

  result.push(person);
});

console.log(JSON.stringify(result, null, 2));