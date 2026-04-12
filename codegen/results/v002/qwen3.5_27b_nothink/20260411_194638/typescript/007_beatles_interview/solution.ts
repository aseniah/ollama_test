import * as fs from 'fs';
import * as path from 'path';

const DATE_CUTOFF = new Date('2025-07-01T00:00:00Z');

interface RelativeData {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

interface OutputRecord {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: RelativeData[];
}

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null') return null;
  // Format: M/D/YYYY
  const parts = dateStr.split('/').map(Number);
  // Note: Month is 0-indexed in JS Date constructor
  return new Date(parts[2], parts[0] - 1, parts[1]);
}

function calculateAge(birthDate: Date): number {
  let age = DATE_CUTOFF.getFullYear() - birthDate.getFullYear();
  
  const monthCheck = DATE_CUTOFF.getMonth() - birthDate.getMonth();
  if (
    monthCheck < 0 ||
    (monthCheck === 0 && DATE_CUTOFF.getDate() < birthDate.getDate())
  ) {
    age--;
  }
  return age;
}

function formatISODate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function parseFullNames(fullName: string): [string, string] {
  const parts = fullName.trim().split(' ');
  const lastName = parts.pop()!;
  const firstName = parts.join(' ');
  return [firstName, lastName];
}

function parseRelative(
  fullName: string | undefined | null,
  relationship: string
): RelativeData | null {
  if (!fullName || fullName === 'null') return null;
  const [firstName, lastName] = parseFullNames(fullName);
  return {
    FirstName: firstName,
    LastName: lastName,
    Relationship: relationship,
  };
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  
  if (!fs.existsSync(inputPath)) {
    console.error('Input file not found.');
    process.exit(1);
  }

  const content = fs.readFileSync(inputPath, 'utf-8');
  const lines = content.split('\n').filter(line => line.trim() !== '');

  const headers = lines[0].split(',');
  const output: OutputRecord[] = [];

  for (let i = 1; i < lines.length; i++) {
    const row = lines[i].split(',');

    // Map columns to variables
    const fullName = row[headers.indexOf('Name')];
    const birthdayStr = row[headers.indexOf('Birthday')];
    // Died is at index 2, but not used in output
    const fatherName = row[headers.indexOf('Father')];
    const motherName = row[headers.indexOf('Mother')];
    const brotherName = row[headers.indexOf('Brother')];
    const sisterName = row[headers.indexOf('Sister')];

    if (!fullName) continue;

    const [firstName, lastName] = parseFullNames(fullName);
    const birthDate = parseDate(birthdayStr);

    if (!birthDate) continue;

    const age = calculateAge(birthDate);

    const relatives: RelativeData[] = [];

    if (fatherName && fatherName !== 'null') {
      relatives.push({ FirstName: fatherName, LastName: 'Father', Relationship: 'Father' });
    }
    // Wait, looking at the expected output again:
    // "FirstName": "Alfred", "LastName": "Lennon"
    // The CSV "Father" column contains "Alfred Lennon".
    // So I need to split the name string, not use "Father" as the last name.
    
    const relativesRaw: RelativeData[] = [];
    
    // Re-parse father
    const parsedFather = parseRelative(fatherName, 'Father');
    if (parsedFather) relativesRaw.push(parsedFather);

    const parsedMother = parseRelative(motherName, 'Mother');
    if (parsedMother) relativesRaw.push(parsedMother);

    const parsedBrother = parseRelative(brotherName, 'Brother');
    if (parsedBrother) relativesRaw.push(parsedBrother);

    const parsedSister = parseRelative(sisterName, 'Sister');
    if (parsedSister) relativesRaw.push(parsedSister);

    output.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: formatISODate(birthDate),
      Age: age,
      Relatives: relativesRaw,
    });
  }

  console.log(JSON.stringify(output, null, 2));
}

main();