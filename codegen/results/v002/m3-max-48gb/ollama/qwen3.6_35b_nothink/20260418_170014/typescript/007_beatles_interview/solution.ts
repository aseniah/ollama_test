import * as fs from 'fs';
import * as path from 'path';

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: {
    FirstName: string;
    LastName: string;
    Relationship: string;
  }[];
}

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

function parseBirthday(birthdayStr: string): Date {
  const parts = birthdayStr.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
}

function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function calculateAge(birthday: string): number {
  const birthDate = parseBirthday(birthday);
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025

  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  const dayDiff = referenceDate.getDate() - birthDate.getDate();

  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age -= 1;
  }

  return age;
}

function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(/\s+/);
  if (parts.length < 2) {
    throw new Error(`Invalid name format: ${name}`);
  }
  const lastName = parts.pop()!;
  const firstName = parts.join(' ');
  return { firstName, lastName };
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  const lines = csvContent.trim().split('\n');
  const headers = lines[0].split(',');

  const results: Person[] = [];

  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    if (values.length !== headers.length) {
      continue;
    }

    const nameStr = values[headers.indexOf('Name')].trim();
    const birthdayStr = values[headers.indexOf('Birthday')].trim();
    const diedStr = values[headers.indexOf('Died')].trim();
    const fatherStr = values[headers.indexOf('Father')].trim();
    const motherStr = values[headers.indexOf('Mother')].trim();
    const brotherStr = values[headers.indexOf('Brother')].trim();
    const sisterStr = values[headers.indexOf('Sister')].trim();

    const { firstName, lastName } = parseName(nameStr);
    const age = calculateAge(birthdayStr);

    const relatives: Relative[] = [];

    if (fatherStr && fatherStr !== 'null') {
      const { firstName: fFirstName, lastName: fLastName } = parseName(fatherStr);
      relatives.push({ FirstName: fFirstName, LastName: fLastName, Relationship: 'Father' });
    }

    if (motherStr && motherStr !== 'null') {
      const { firstName: mFirstName, lastName: mLastName } = parseName(motherStr);
      relatives.push({ FirstName: mFirstName, LastName: mLastName, Relationship: 'Mother' });
    }

    if (brotherStr && brotherStr !== 'null') {
      const { firstName: bFirstName, lastName: bLastName } = parseName(brotherStr);
      relatives.push({ FirstName: bFirstName, LastName: bLastName, Relationship: 'Brother' });
    }

    if (sisterStr && sisterStr !== 'null') {
      const { firstName: sFirstName, lastName: sLastName } = parseName(sisterStr);
      relatives.push({ FirstName: sFirstName, LastName: sLastName, Relationship: 'Sister' });
    }

    const person: Person = {
      FirstName: firstName,
      LastName: lastName,
      Birthday: formatDate(parseBirthday(birthdayStr)),
      Age: age,
      Relatives: relatives
    };

    results.push(person);
  }

  console.log(JSON.stringify(results, null, 2));
}

main();