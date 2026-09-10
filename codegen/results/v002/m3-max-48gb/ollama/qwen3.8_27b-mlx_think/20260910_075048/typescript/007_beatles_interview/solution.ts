import * as fs from 'fs';

function calculateAge(birthDate: Date, asOf: Date): number {
  let age = asOf.getFullYear() - birthDate.getFullYear();
  const monthDiff = asOf.getMonth() - birthDate.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && asOf.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
}

function parseDate(mdy: string): Date {
  const [m, d, y] = mdy.split('/').map(Number);
  return new Date(y, m - 1, d);
}

function formatDate(date: Date): string {
  const y = date.getFullYear();
  const m = String(date.getMonth() + 1).padStart(2, '0');
  const d = String(date.getDate()).padStart(2, '0');
  return `${y}-${m}-${d}`;
}

const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');
const dataLines = lines.slice(1);

const asOfDate = new Date(2025, 6, 1); // July 1, 2025

const result: any[] = [];

for (const line of dataLines) {
  const fields = line.split(',');
  const record: Record<string, string> = {};
  headers.forEach((h, i) => {
    record[h] = fields[i]?.trim() ?? '';
  });

  const [firstName, ...lastNameParts] = record['Name'].split(' ');
  const lastName = lastNameParts.join(' ');

  const birthday = parseDate(record['Birthday']);
  const diedStr = record['Died'];
  const died = diedStr && diedStr !== 'null' ? parseDate(diedStr) : null;

  const age = died ? calculateAge(birthday, died) : calculateAge(birthday, asOfDate);

  const relatives: any[] = [];

  const relativeFields: [string, string][] = [
    ['Father', record['Father']],
    ['Mother', record['Mother']],
    ['Brother', record['Brother']],
    ['Sister', record['Sister']],
  ];

  for (const [relationship, nameStr] of relativeFields) {
    if (nameStr && nameStr !== 'null') {
      const [relFirst, ...relLastParts] = nameStr.split(' ');
      const relLast = relLastParts.join(' ');
      relatives.push({
        FirstName: relFirst,
        LastName: relLast,
        Relationship: relationship,
      });
    }
  }

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: formatDate(birthday),
    Age: age,
    Relatives: relatives,
  });
}

process.stdout.write(JSON.stringify(result, null, 2) + '\n');