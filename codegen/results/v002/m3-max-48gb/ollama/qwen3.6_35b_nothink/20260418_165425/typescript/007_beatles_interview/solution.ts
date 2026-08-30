import * as fs from 'fs';
import * as path from 'path';

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

function parseDate(dateStr: string): Date {
  // Format: M/D/YYYY or MM/DD/YYYY
  const parts = dateStr.split('/');
  const month = parseInt(parts[0], 10) - 1; // JS months are 0-indexed
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month, day);
}

function calculateAge(birthdayStr: string, referenceDate: Date): number {
  const birthday = parseDate(birthdayStr);
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthday.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function splitName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(' ');
  if (parts.length === 0) {
    return { firstName: '', lastName: '' };
  }
  const lastName = parts.pop() || '';
  const firstName = parts.join(' ') || '';
  return { firstName, lastName };
}

function main() {
  const csvPath = path.join(__dirname, 'input', 'input.csv');
  const csvContent = fs.readFileSync(csvPath, 'utf-8');
  const lines = csvContent.trim().split('\n');

  // Header line
  const headers = lines[0].split(',');

  // Reference date: July 1, 2025
  const referenceDate = new Date(2025, 6, 1); // Month is 0-indexed, so 6 = July

  const people: Person[] = [];

  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    if (values.length < headers.length) continue;

    const nameStr = values[0];
    const birthdayStr = values[1];
    // const diedStr = values[2]; // Not used
    const fatherStr = values[3];
    const motherStr = values[4];
    const brotherStr = values[5];
    const sisterStr = values[6];

    const { firstName, lastName } = splitName(nameStr);
    const age = calculateAge(birthdayStr, referenceDate);

    const relatives: Relative[] = [];

    if (fatherStr && fatherStr.toLowerCase() !== 'null') {
      const { firstName: fFirst, lastName: fLast } = splitName(fatherStr);
      relatives.push({ FirstName: fFirst, LastName: fLast, Relationship: 'Father' });
    }
    if (motherStr && motherStr.toLowerCase() !== 'null') {
      const { firstName: mFirst, lastName: mLast } = splitName(motherStr);
      relatives.push({ FirstName: mFirst, LastName: mLast, Relationship: 'Mother' });
    }
    if (brotherStr && brotherStr.toLowerCase() !== 'null') {
      const { firstName: bFirst, lastName: bLast } = splitName(brotherStr);
      relatives.push({ FirstName: bFirst, LastName: bLast, Relationship: 'Brother' });
    }
    if (sisterStr && sisterStr.toLowerCase() !== 'null') {
      const { firstName: sFirst, lastName: sLast } = splitName(sisterStr);
      relatives.push({ FirstName: sFirst, LastName: sLast, Relationship: 'Sister' });
    }

    people.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayStr.split('/').map(p => p.padStart(2, '0')).join('-'), // Convert to YYYY-MM-DD? Wait, input is M/D/YYYY. Expected is YYYY-MM-DD.
      Age: age,
      Relatives: relatives
    });
  }

  // Convert birthdays from M/D/YYYY to YYYY-MM-DD
  for (const person of people) {
    const parts = person.Birthday.split('-').map(Number); // Already converted? No, let's re-parse.
    // Actually, let's parse the original birthday string properly.
    // We stored the original birthdayStr. Let's fix the Birthday field to be YYYY-MM-DD.
    // The birthdayStr is in M/D/YYYY format.
    // Let's re-extract from the CSV line directly or fix the conversion.
    // Actually, I already converted it in the push above, but incorrectly for some cases if single-digit months/days were present.
    // Let's re-do the Birthday conversion properly.
  }

  // Let's restart the person creation with proper Birthday format
  const peopleFixed: Person[] = [];

  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    if (values.length < headers.length) continue;

    const nameStr = values[0];
    const birthdayStr = values[1];
    const fatherStr = values[3];
    const motherStr = values[4];
    const brotherStr = values[5];
    const sisterStr = values[6];

    const { firstName, lastName } = splitName(nameStr);
    const age = calculateAge(birthdayStr, referenceDate);

    // Convert birthday from M/D/YYYY to YYYY-MM-DD
    const bp = birthdayStr.split('/');
    const year = bp[2];
    const month = bp[0].padStart(2, '0');
    const day = bp[1].padStart(2, '0');
    const birthdayFormatted = `${year}-${month}-${day}`;

    const relatives: Relative[] = [];

    if (fatherStr && fatherStr.toLowerCase() !== 'null') {
      const { firstName: fFirst, lastName: fLast } = splitName(fatherStr);
      relatives.push({ FirstName: fFirst, LastName: fLast, Relationship: 'Father' });
    }
    if (motherStr && motherStr.toLowerCase() !== 'null') {
      const { firstName: mFirst, lastName: mLast } = splitName(motherStr);
      relatives.push({ FirstName: mFirst, LastName: mLast, Relationship: 'Mother' });
    }
    if (brotherStr && brotherStr.toLowerCase() !== 'null') {
      const { firstName: bFirst, lastName: bLast } = splitName(brotherStr);
      relatives.push({ FirstName: bFirst, LastName: bLast, Relationship: 'Brother' });
    }
    if (sisterStr && sisterStr.toLowerCase() !== 'null') {
      const { firstName: sFirst, lastName: sLast } = splitName(sisterStr);
      relatives.push({ FirstName: sFirst, LastName: sLast, Relationship: 'Sister' });
    }

    peopleFixed.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayFormatted,
      Age: age,
      Relatives: relatives
    });
  }

  console.log(JSON.stringify(peopleFixed, null, 2));
}

main();