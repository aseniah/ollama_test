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

function parseDate(s: string): Date {
  const [m, d, y] = s.split('/').map(Number);
  return new Date(y, m - 1, d);
}

function formatDate(d: Date): string {
  const y = d.getFullYear();
  const m = String(d.getMonth() + 1).padStart(2, '0');
  const day = String(d.getDate()).padStart(2, '0');
  return `${y}-${m}-${day}`;
}

function calculateAge(birthDate: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function splitName(fullName: string): { first: string; last: string } {
  const parts = fullName.trim().split(/\s+/);
  const first = parts[0];
  const last = parts[parts.length - 1];
  return { first, last };
}

// Read input
const csvPath = path.join('input', 'input.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8').trim();
const lines = csvContent.split('\n');
const headers = lines[0].split(',');
const rows = lines.slice(1);

const referenceDate = new Date(2025, 6, 1); // July 1, 2025

const result: Person[] = [];

for (const row of rows) {
  const fields = row.split(',');
  const record: Record<string, string> = {};
  headers.forEach((h, i) => {
    record[h.trim()] = (fields[i] || '').trim();
  });

  // Name
  const { first, last } = splitName(record['Name']);

  // Birthday
  const birthDate = parseDate(record['Birthday']);
  const birthdayStr = formatDate(birthDate);

  // Age
  let age: number;
  if (record['Died'] && record['Died'] !== 'null') {
    const deathDate = parseDate(record['Died']);
    age = calculateAge(birthDate, deathDate);
  } else {
    age = calculateAge(birthDate, referenceDate);
  }

  // Relatives
  const relatives: Relative[] = [];
  const relFields = ['Father', 'Mother', 'Brother', 'Sister'];
  for (const rel of relFields) {
    const val = record[rel];
    if (val && val !== 'null') {
      const relParts = splitName(val);
      relatives.push({
        FirstName: relParts.first,
        LastName: relParts.last,
        Relationship: rel,
      });
    }
  }

  result.push({
    FirstName: first,
    LastName: last,
    Birthday: birthdayStr,
    Age: age,
    Relatives: relatives,
  });
}

process.stdout.write(JSON.stringify(result, null, 2) + '\n');