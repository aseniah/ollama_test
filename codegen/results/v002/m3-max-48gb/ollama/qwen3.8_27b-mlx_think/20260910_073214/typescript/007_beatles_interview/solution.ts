import { readFileSync } from 'fs';

// Read CSV
const csvContent = readFileSync('input/input.csv', 'utf-8').trim();
const lines = csvContent.split('\n');
const headers = lines[0].split(',');

const referenceDate = new Date(2025, 6, 1); // July 1, 2025

const results = lines.slice(1).map(line => {
  const fields = line.split(',');
  const record: Record<string, string> = {};
  headers.forEach((h, i) => {
    record[h.trim()] = (fields[i] || '').trim();
  });

  // Split name into FirstName and LastName
  const nameParts = record['Name'].split(' ');
  const firstName = nameParts[0];
  const lastName = nameParts[nameParts.length - 1];

  // Convert birthday from MM/DD/YYYY to YYYY-MM-DD
  const [birthMonth, birthDay, birthYear] = record['Birthday'].split('/').map(Number);
  const birthday = `${birthYear}-${String(birthMonth).padStart(2, '0')}-${String(birthDay).padStart(2, '0')}`;

  // Calculate age: at death if deceased, otherwise at July 1, 2025
  let age: number;
  const died = record['Died'];
  if (died !== 'null') {
    const [diedMonth, diedDay, diedYear] = died.split('/').map(Number);
    const deathDate = new Date(diedYear, diedMonth - 1, diedDay);
    const birthDate = new Date(birthYear, birthMonth - 1, birthDay);
    age = deathDate.getFullYear() - birthDate.getFullYear();
    const birthdayInDeathYear = new Date(deathDate.getFullYear(), birthMonth - 1, birthDay);
    if (deathDate < birthdayInDeathYear) {
      age--;
    }
  } else {
    const birthDate = new Date(birthYear, birthMonth - 1, birthDay);
    age = referenceDate.getFullYear() - birthDate.getFullYear();
    const birthdayThisYear = new Date(referenceDate.getFullYear(), birthMonth - 1, birthDay);
    if (referenceDate < birthdayThisYear) {
      age--;
    }
  }

  // Build relatives array
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];

  const addRelative = (name: string, relationship: string): void => {
    if (name !== 'null') {
      const parts = name.split(' ');
      relatives.push({
        FirstName: parts[0],
        LastName: parts[parts.length - 1],
        Relationship: relationship,
      });
    }
  };

  addRelative(record['Father'], 'Father');
  addRelative(record['Mother'], 'Mother');
  addRelative(record['Brother'], 'Brother');
  addRelative(record['Sister'], 'Sister');

  return {
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthday,
    Age: age,
    Relatives: relatives,
  };
});

console.log(JSON.stringify(results, null, 2));