import * as fs from 'fs';
import * as path from 'path';

function parseDate(dateStr: string): { year: number; month: number; day: number } {
  const parts = dateStr.split('/').map(Number);
  return { month: parts[0], day: parts[1], year: parts[2] };
}

function formatBirthday(dateStr: string): string {
  const { month, day, year } = parseDate(dateStr);
  return `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
}

function calculateAge(birthdayStr: string, diedStr: string | null): number {
  const refDate = diedStr ? parseDate(diedStr) : { year: 2025, month: 7, day: 1 };
  const birthday = parseDate(birthdayStr);

  let age = refDate.year - birthday.year;

  // If the birthday hasn't occurred yet in the reference year, subtract 1
  if (refDate.month < birthday.month || (refDate.month === birthday.month && refDate.day < birthday.day)) {
    age--;
  }

  return age;
}

function parseName(nameStr: string): { FirstName: string; LastName: string } {
  const parts = nameStr.trim().split(/\s+/);
  return {
    FirstName: parts[0],
    LastName: parts[parts.length - 1]
  };
}

function parseRelatives(fields: string[]): Array<{ FirstName: string; LastName: string; Relationship: string }> {
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  const relationshipMap: Array<{ index: number; relationship: string }> = [
    { index: 3, relationship: 'Father' },
    { index: 4, relationship: 'Mother' },
    { index: 5, relationship: 'Brother' },
    { index: 6, relationship: 'Sister' },
  ];

  for (const { index, relationship } of relationshipMap) {
    const value = fields[index];
    if (value && value !== 'null') {
      const { FirstName, LastName } = parseName(value);
      relatives.push({ FirstName, LastName, Relationship: relationship });
    }
  }

  return relatives;
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  const lines = content.trim().split('\n');

  const dataLines = lines.slice(1);
  const result: any[] = [];

  for (const line of dataLines) {
    if (!line.trim()) continue;
    const fields = line.split(',');
    const name = fields[0];
    const birthday = fields[1];
    const died = fields[2];

    const { FirstName, LastName } = parseName(name);
    const Birthday = formatBirthday(birthday);
    const Died = died !== 'null' ? died : null;
    const Age = calculateAge(birthday, Died);
    const Relatives = parseRelatives(fields);

    result.push({ FirstName, LastName, Birthday, Age, Relatives });
  }

  process.stdout.write(JSON.stringify(result, null, 2) + '\n');
}

main();