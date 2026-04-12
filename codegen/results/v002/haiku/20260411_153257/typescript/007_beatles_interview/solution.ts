import fs from 'fs';
import path from 'path';

const csvPath = path.join(process.cwd(), 'input/input.csv');
const content = fs.readFileSync(csvPath, 'utf-8');
const lines = content.trim().split('\n');
const headers = lines[0].split(',');

const refDate = new Date(2025, 6, 1); // July 1, 2025

function calcAge(birthday: string, died: string | null): number {
  const [month, day, year] = birthday.split('/').map(Number);
  const birth = new Date(year, month - 1, day);
  const end = died ? (() => { const [m,d,y] = died.split('/').map(Number); return new Date(y, m-1, d); })() : refDate;
  let age = end.getFullYear() - birth.getFullYear();
  if (end < new Date(end.getFullYear(), birth.getMonth(), birth.getDate())) age--;
  return age;
}

function formatDate(dateStr: string): string {
  const [month, day, year] = dateStr.split('/').map(Number);
  return `${year}-${String(month).padStart(2,'0')}-${String(day).padStart(2,'0')}`;
}

function parseName(full: string): { FirstName: string; LastName: string } {
  const parts = full.trim().split(' ');
  const LastName = parts[parts.length - 1];
  const FirstName = parts[0];
  return { FirstName, LastName };
}

const result = [];

for (let i = 1; i < lines.length; i++) {
  const cols = lines[i].split(',');
  const row: Record<string, string> = {};
  headers.forEach((h, idx) => { row[h.trim()] = cols[idx]?.trim() ?? ''; });

  const nameParts = row['Name'].split(' ');
  const FirstName = nameParts[0];
  const LastName = nameParts[nameParts.length - 1];
  const birthday = row['Birthday'];
  const died = row['Died'] === 'null' || !row['Died'] ? null : row['Died'];

  const Age = calcAge(birthday, died);
  const Birthday = formatDate(birthday);

  const Relatives = [];
  for (const rel of ['Father', 'Mother', 'Brother', 'Sister']) {
    const val = row[rel];
    if (val && val !== 'null') {
      const { FirstName: fn, LastName: ln } = parseName(val);
      Relatives.push({ FirstName: fn, LastName: ln, Relationship: rel });
    }
  }

  result.push({ FirstName, LastName, Birthday, Age, Relatives });
}

console.log(JSON.stringify(result, null, 2));
