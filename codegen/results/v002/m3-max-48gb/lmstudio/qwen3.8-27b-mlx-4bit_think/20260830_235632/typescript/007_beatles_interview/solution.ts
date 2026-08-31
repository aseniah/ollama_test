import * as fs from 'fs';

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

const csvContent = fs.readFileSync('input/input.csv', 'utf-8').trim();
const lines = csvContent.split('\n');
const rows = lines.slice(1);

function parseDate(dateStr: string): { year: number; month: number; day: number } {
  const parts = dateStr.split('/').map(Number);
  return { month: parts[0], day: parts[1], year: parts[2] };
}

function formatDate(dateStr: string): string {
  const { year, month, day } = parseDate(dateStr);
  return `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
}

function calcAge(birthStr: string, refYear: number, refMonth: number, refDay: number): number {
  const { year, month, day } = parseDate(birthStr);
  let age = refYear - year;
  if (refMonth < month || (refMonth === month && refDay < day)) {
    age--;
  }
  return age;
}

function splitName(fullName: string): { first: string; last: string } {
  const parts = fullName.trim().split(' ');
  const first = parts[0];
  const last = parts[parts.length - 1];
  return { first, last };
}

// July 1, 2025 (1-indexed month for comparison)
const REF = { year: 2025, month: 7, day: 1 };

const result: Person[] = [];

for (const row of rows) {
  const fields = row.split(',');
  const name = fields[0];
  const birthday = fields[1].trim();
  const died = fields[2].trim();
  const father = fields[3].trim();
  const mother = fields[4].trim();
  const brother = fields[5].trim();
  const sister = fields[6].trim();

  const { first: FirstName, last: LastName } = splitName(name);
  const Birthday = formatDate(birthday);

  let Age: number;
  if (died && died !== 'null') {
    const d = parseDate(died);
    Age = calcAge(birthday, d.year, d.month, d.day);
  } else {
    Age = calcAge(birthday, REF.year, REF.month, REF.day);
  }

  const relatives: Relative[] = [];
  const relData: [string, string][] = [
    [father, 'Father'],
    [mother, 'Mother'],
    [brother, 'Brother'],
    [sister, 'Sister'],
  ];

  for (const [val, relationship] of relData) {
    if (val && val !== 'null') {
      const { first, last } = splitName(val);
      relatives.push({ FirstName: first, LastName: last, Relationship: relationship });
    }
  }

  result.push({ FirstName, LastName, Birthday, Age, Relatives: relatives });
}

process.stdout.write(JSON.stringify(result, null, 2) + '\n');