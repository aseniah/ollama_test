import * as fs from 'fs';
import * as path from 'path';

const inputPath = path.join('input', 'input.csv');
const content = fs.readFileSync(inputPath, 'utf-8');

const lines = content.trim().split('\n');
const headers = lines[0].split(',').map(h => h.trim());

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

function parseName(fullName: string): { first: string; last: string } {
  const parts = fullName.trim().split(' ');
  const first = parts[0];
  const last = parts[parts.length - 1];
  return { first, last };
}

function parseDate(dateStr: string): { month: number; day: number; year: number } {
  const parts = dateStr.split('/').map(Number);
  return { month: parts[0], day: parts[1], year: parts[2] };
}

function toISO(month: number, day: number, year: number): string {
  return `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
}

function calculateAge(birthDate: string, deathDate: string | null, refDate: Date): number {
  const ref = deathDate
    ? parseDate(deathDate)
    : { month: refDate.getMonth() + 1, day: refDate.getDate(), year: refDate.getFullYear() };

  const birth = parseDate(birthDate);

  let age = ref.year - birth.year;
  if (ref.month < birth.month || (ref.month === birth.month && ref.day < birth.day)) {
    age--;
  }

  return age;
}

const referenceDate = new Date(2025, 6, 1); // July 1, 2025

const result: Person[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',').map(v => v.trim());
  const row: Record<string, string> = {};
  headers.forEach((h, idx) => {
    row[h] = values[idx] ?? '';
  });

  const nameParts = parseName(row['Name']);

  const birthday = parseDate(row['Birthday']);
  const birthdayStr = toISO(birthday.month, birthday.day, birthday.year);

  const died = row['Died'] === 'null' || row['Died'] === '' ? null : row['Died'];

  const age = calculateAge(row['Birthday'], died, referenceDate);

  const relatives: Relative[] = [];
  const relColumns = ['Father', 'Mother', 'Brother', 'Sister'];

  for (const col of relColumns) {
    const val = row[col];
    if (val && val !== 'null') {
      const relName = parseName(val);
      relatives.push({
        FirstName: relName.first,
        LastName: relName.last,
        Relationship: col
      });
    }
  }

  result.push({
    FirstName: nameParts.first,
    LastName: nameParts.last,
    Birthday: birthdayStr,
    Age: age,
    Relatives: relatives
  });
}

process.stdout.write(JSON.stringify(result, null, 2) + '\n');