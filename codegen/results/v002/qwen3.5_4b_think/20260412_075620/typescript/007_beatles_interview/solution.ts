import { readFileSync, writeFileSync } from 'fs';
import { join } from 'path';

const inputPath = join(process.cwd(), 'input/input.csv');
const outputPath = join(process.cwd(), 'output/expected_format.json');

const fileContent = readFileSync(inputPath, 'utf-8');
const lines = fileContent.trim().split('\n');

const parseDate = (dateStr: string): Date => {
  // Handle "m/d/yyyy" format
  const parts = dateStr.split('/');
  if (parts.length !== 3) return new Date();
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
};

const parseName = (nameStr: string) => {
  const parts = nameStr.split(' ');
  const firstName = parts[0];
  const lastName = parts.slice(1).join(' ');
  return { firstName, lastName };
};

const relativesData: {
  firstName: string;
  lastName: string;
  relationship: string;
}[] = [];

lines.forEach((line, i) => {
  if (i === 0) return; // Skip header
  
  const columns = line.split(',');
  const nameStr = columns[0];
  const birthdayStr = columns[1];
  const diedStr = columns[2] || 'null';
  const fatherStr = columns[3] || 'null';
  const motherStr = columns[4] || 'null';
  const brotherStr = columns[5] || 'null';
  const sisterStr = columns[6] || 'null';
  
  const person = parseName(nameStr);
  const birthday = parseDate(birthdayStr);
  
  const relatives: {
    firstName: string;
    lastName: string;
    relationship: string;
  }[] = [];
  
  if (fatherStr !== 'null') {
    const f = parseName(fatherStr);
    relatives.push({ firstName: f.firstName, lastName: f.lastName, relationship: 'Father' });
  }
  if (motherStr !== 'null') {
    const m = parseName(motherStr);
    relatives.push({ firstName: m.firstName, lastName: m.lastName, relationship: 'Mother' });
  }
  if (brotherStr !== 'null') {
    const b = parseName(brotherStr);
    relatives.push({ firstName: b.firstName, lastName: b.lastName, relationship: 'Brother' });
  }
  if (sisterStr !== 'null') {
    const s = parseName(sisterStr);
    relatives.push({ firstName: s.firstName, lastName: s.lastName, relationship: 'Sister' });
  }
  
  // Calculate age as of July 1, 2025
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (0-indexed months)
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  
  // If birthday hasn't occurred yet in 2025, subtract 1
  if (referenceDate.getMonth() + 1 < birthday.getMonth() ||
      (referenceDate.getMonth() + 1 === birthday.getMonth() && referenceDate.getDate() < birthday.getDate())) {
    age--;
  }
  
  const output = [
    {
      FirstName: person.firstName,
      LastName: person.lastName,
      Birthday: `${birthday.getFullYear()}-${String(birthday.getMonth() + 1).padStart(2, '0')}-${String(birthday.getDate()).padStart(2, '0')}`,
      Age: age,
      Relatives: relatives
    }
  ];
  
  console.log(JSON.stringify(output, null, 2));
});