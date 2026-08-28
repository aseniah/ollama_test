import * as fs from 'fs';

const inputFile = 'input/input.csv';
const referenceDate = new Date('2025-07-01');

// Parse CSV
const csvContent = fs.readFileSync(inputFile, 'utf-8');
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');
const dataLines = lines.slice(1);

// Helper function to format birthday
function formatBirthday(birthdayStr: string): string {
  const parts = birthdayStr.split('/');
  const month = parts[0].padStart(2, '0');
  const day = parts[1].padStart(2, '0');
  const year = parts[2];
  return `${year}-${month}-${day}`;
}

// Helper function to parse name into firstName and lastName
function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(/\s+/);
  const firstName = parts[0];
  const lastName = parts[parts.length - 1];
  return { firstName, lastName };
}

// Helper function to calculate age as of reference date
function calculateAge(birthday: string): number {
  const birthDate = new Date(birthday);
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  const dayDiff = referenceDate.getDate() - birthDate.getDate();
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }
  return age;
}

// Helper function to create relative object
function createRelative(name: string, relationship: string): {
  FirstName: string;
  LastName: string;
  Relationship: string;
} | null {
  if (name === 'null' || name.trim() === '') {
    return null;
  }
  const parts = name.trim().split(/\s+/);
  return {
    FirstName: parts[0],
    LastName: parts[parts.length - 1],
    Relationship: relationship,
  };
}

// Process each data line
const results: any[] = [];

for (const line of dataLines) {
  const values = line.split(',');
  
  const name = values[0];
  const birthday = values[1];
  const father = values[3];
  const mother = values[4];
  const brother = values[5];
  const sister = values[6];
  
  const { firstName, lastName } = parseName(name);
  const formattedBirthday = formatBirthday(birthday);
  const age = calculateAge(formattedBirthday);
  
  const relatives: any[] = [];
  const fatherRel = createRelative(father, 'Father');
  if (fatherRel) relatives.push(fatherRel);
  const motherRel = createRelative(mother, 'Mother');
  if (motherRel) relatives.push(motherRel);
  const brotherRel = createRelative(brother, 'Brother');
  if (brotherRel) relatives.push(brotherRel);
  const sisterRel = createRelative(sister, 'Sister');
  if (sisterRel) relatives.push(sisterRel);
  
  results.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: formattedBirthday,
    Age: age,
    Relatives: relatives,
  });
}

console.log(JSON.stringify(results));