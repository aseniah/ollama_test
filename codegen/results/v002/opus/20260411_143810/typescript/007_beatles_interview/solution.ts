import fs from 'fs';
import path from 'path';

const csvPath = path.join(process.cwd(), 'input', 'input.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8');

const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');

const REF_DATE = new Date('2025-07-01');

function calcAge(birthday: Date, died: Date | null): number {
  const end = died && died < REF_DATE ? died : REF_DATE;
  let age = end.getFullYear() - birthday.getFullYear();
  const m = end.getMonth() - birthday.getMonth();
  if (m < 0 || (m === 0 && end.getDate() < birthday.getDate())) age--;
  return age;
}

function parseName(full: string): { FirstName: string; LastName: string } {
  const parts = full.trim().split(' ');
  const LastName = parts[parts.length - 1];
  const FirstName = parts[0];
  return { FirstName, LastName };
}

function parseDate(val: string): Date {
  const [m, d, y] = val.split('/').map(Number);
  return new Date(y, m - 1, d);
}

function formatDate(d: Date): string {
  const y = d.getFullYear();
  const m = String(d.getMonth() + 1).padStart(2, '0');
  const day = String(d.getDate()).padStart(2, '0');
  return `${y}-${m}-${day}`;
}

const result = lines.slice(1).map(line => {
  const values = line.split(',');
  const row: Record<string, string> = {};
  headers.forEach((h, i) => { row[h.trim()] = (values[i] ?? '').trim(); });

  const { FirstName, LastName } = parseName(row['Name']);
  const birthday = parseDate(row['Birthday']);
  const died = row['Died'] && row['Died'] !== 'null' ? parseDate(row['Died']) : null;
  const age = calcAge(birthday, died);

  const relFields = ['Father', 'Mother', 'Brother', 'Sister'];
  const Relatives = relFields
    .filter(r => row[r] && row[r] !== 'null')
    .map(r => {
      const { FirstName: rFirst, LastName: rLast } = parseName(row[r]);
      return { FirstName: rFirst, LastName: rLast, Relationship: r };
    });

  return {
    FirstName,
    LastName,
    Birthday: formatDate(birthday),
    Age: age,
    Relatives,
  };
});

process.stdout.write(JSON.stringify(result, null, 2) + '\n');
