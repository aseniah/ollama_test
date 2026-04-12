import fs from 'fs';
import path from 'path';

const csvPath = path.join(process.cwd(), 'input', 'input.csv');
const csv = fs.readFileSync(csvPath, 'utf-8');

const lines = csv.trim().split('\n');
const headers = lines[0].split(',');

const referenceDate = new Date(2025, 6, 1);

function calcAge(birthday: Date, died: Date | null): number {
  const end = died && died < referenceDate ? died : referenceDate;
  let age = end.getFullYear() - birthday.getFullYear();
  const m = end.getMonth() - birthday.getMonth();
  if (m < 0 || (m === 0 && end.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function parseName(full: string): { FirstName: string; LastName: string } {
  const parts = full.trim().split(' ');
  const LastName = parts[parts.length - 1];
  const FirstName = parts[0];
  return { FirstName, LastName };
}

function formatDate(mmddyyyy: string): string {
  const [m, d, y] = mmddyyyy.split('/');
  return `${y}-${m.padStart(2, '0')}-${d.padStart(2, '0')}`;
}

const result = lines.slice(1).map(line => {
  const values = line.split(',');
  const row: Record<string, string> = {};
  headers.forEach((h, i) => { row[h.trim()] = (values[i] || '').trim(); });

  const nameParts = row['Name'].split(' ');
  const LastName = nameParts[nameParts.length - 1];
  const FirstName = nameParts[0];

  const birthdayFormatted = formatDate(row['Birthday']);
  const birthday = new Date(row['Birthday']);

  const diedRaw = row['Died'];
  const died = diedRaw && diedRaw !== 'null' ? new Date(diedRaw) : null;

  const age = calcAge(birthday, died);

  const relFields = ['Father', 'Mother', 'Brother', 'Sister'];
  const relatives = relFields
    .filter(rel => row[rel] && row[rel] !== 'null')
    .map(rel => {
      const { FirstName: fn, LastName: ln } = parseName(row[rel]);
      return { FirstName: fn, LastName: ln, Relationship: rel };
    });

  return {
    FirstName,
    LastName,
    Birthday: birthdayFormatted,
    Age: age,
    Relatives: relatives
  };
});

process.stdout.write(JSON.stringify(result, null, 2) + '\n');
