import * as fs from 'fs';
import * as path from 'path';

const REFERENCE_DATE = new Date(2025, 6, 1); // July 1, 2025

function calculateAge(birthdayStr: string): number {
  const [mStr, dStr, yStr] = birthdayStr.split('/');
  const month = parseInt(mStr, 10);
  const day = parseInt(dStr, 10);
  const year = parseInt(yStr, 10);
  const birthdayThisYear = new Date(REFERENCE_DATE.getFullYear(), month - 1, day);
  let age = REFERENCE_DATE.getFullYear() - year;
  if (birthdayThisYear > REFERENCE_DATE) {
    age--;
  }
  return age;
}

function formatBirthday(birthdayStr: string): string {
  const [m, d, y] = birthdayStr.split('/');
  return `${y}-${m.padStart(2, '0')}-${d.padStart(2, '0')}`;
}

function splitName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(/\s+/);
  return {
    firstName: parts[0],
    lastName: parts[parts.length - 1],
  };
}

function splitRelative(relative: string): { firstName: string; lastName: string } | null {
  if (!relative || relative === 'null') return null;
  const parts = relative.trim().split(/\s+/);
  return { firstName: parts[0], lastName: parts[parts.length - 1] };
}

const csvContent = fs.readFileSync(path.join('input', 'input.csv'), 'utf-8');
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');

const result = lines.slice(1).map((line) => {
  const values = line.split(',');
  const row: Record<string, string> = {};
  headers.forEach((h, i) => {
    row[h.trim()] = values[i] ? values[i].trim() : '';
  });

  const { firstName, lastName } = splitName(row.Name);
  const birthday = formatBirthday(row.Birthday);
  const age = calculateAge(row.Birthday);

  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  const relationshipMap: Array<[string, string]> = [
    ['Father', 'Father'],
    ['Mother', 'Mother'],
    ['Brother', 'Brother'],
    ['Sister', 'Sister'],
  ];

  for (const [col, rel] of relationshipMap) {
    const parsed = splitRelative(row[col]);
    if (parsed) {
      relatives.push({ FirstName: parsed.firstName, LastName: parsed.lastName, Relationship: rel });
    }
  }

  return { FirstName: firstName, LastName: lastName, Birthday: birthday, Age: age, Relatives: relatives };
});

console.log(JSON.stringify(result, null, 2));