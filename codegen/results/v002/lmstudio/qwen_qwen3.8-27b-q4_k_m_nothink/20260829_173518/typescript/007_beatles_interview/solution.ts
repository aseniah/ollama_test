import { readFileSync } from 'fs';

const csvContent = readFileSync('input/input.csv', 'utf-8').trim();
const lines = csvContent.split('\n');
const headers = lines[0].split(',').map(h => h.trim());
const rows = lines.slice(1).map(line => {
  const values = line.split(',').map(v => v.trim());
  const row: Record<string, string> = {};
  headers.forEach((h, i) => { row[h] = values[i]; });
  return row;
});

function splitName(name: string): { first: string; last: string } {
  const parts = name.trim().split(' ');
  const last = parts[parts.length - 1];
  const first = parts.slice(0, -1).join(' ');
  return { first, last };
}

function calculateAge(birthday: string, died: string | null): number {
  // Parse date from M/D/YYYY format
  const [month, day, year] = birthday.split('/').map(Number);
  
  if (died && died !== 'null') {
    const [dMonth, dDay, dYear] = died.split('/').map(Number);
    let age = dYear - year;
    // If death date is before birthday in the year, subtract 1
    if (dMonth < month || (dMonth === month && dDay < day)) {
      age--;
    }
    return age;
  } else {
    // Age as of July 1, 2025
    const refYear = 2025;
    const refMonth = 7;
    const refDay = 1;
    let age = refYear - year;
    if (month > refMonth || (month === refMonth && day > refDay)) {
      age--;
    }
    return age;
  }
}

function buildRelatives(row: Record<string, string>): Array<{ FirstName: string; LastName: string; Relationship: string }> {
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  
  const fields: Array<[string, string]> = [
    ['Father', 'Father'],
    ['Mother', 'Mother'],
    ['Brother', 'Brother'],
    ['Sister', 'Sister']
  ];
  
  for (const [col, relationship] of fields) {
    const value = row[col];
    if (value && value !== 'null') {
      const { first, last } = splitName(value);
      relatives.push({ FirstName: first, LastName: last, Relationship: relationship });
    }
  }
  
  return relatives;
}

const result = rows.map(row => {
  const { first, last } = splitName(row['Name']);
  const died = row['Died'] === 'null' ? null : row['Died'];
  const age = calculateAge(row['Birthday'], died);
  const relatives = buildRelatives(row);
  
  return {
    FirstName: first,
    LastName: last,
    Birthday: row['Birthday'],
    Age: age,
    Relatives: relatives
  };
});

console.log(JSON.stringify(result, null, 2));