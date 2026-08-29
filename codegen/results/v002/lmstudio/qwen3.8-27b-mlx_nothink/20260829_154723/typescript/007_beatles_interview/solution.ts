import * as fs from 'fs';

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr === 'null') return null;
  const [month, day, year] = dateStr.split('/').map(Number);
  if (month === undefined || day === undefined || year === undefined) return null;
  return new Date(year, month - 1, day);
}

function formatDate(date: Date | null): string {
  if (!date) return '';
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function calculateAge(birthday: Date | null, died: Date | null): number {
  if (!birthday) return 0;
  const refDate = died ? died : new Date(2025, 6, 1);
  let age = refDate.getFullYear() - birthday.getFullYear();
  const monthDiff = refDate.getMonth() - birthday.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && refDate.getDate() < birthday.getDate())) {
    age--;
  }
  return age;
}

function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(/\s+/);
  const firstName = parts[0] || '';
  const lastName = parts.length > 1 ? parts[parts.length - 1] : '';
  return { firstName, lastName };
}

function cleanValue(value: string): string {
  const trimmed = value.trim();
  return trimmed === 'null' || trimmed === '' ? '' : trimmed;
}

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

function main() {
  const inputPath = process.argv[2] || 'input/input.csv';
  const content = fs.readFileSync(inputPath, 'utf8');
  const lines = content.split(/\r?\n/).filter((line) => line.trim().length > 0);

  const headers = lines[0].split(',');
  const people: Person[] = [];

  for (let i = 1; i < lines.length; i++) {
    const row = lines[i].split(',');
    if (row.length < headers.length) continue;

    const data: { [key: string]: string } = {};
    headers.forEach((h, idx) => {
      data[h] = row[idx] ?? '';
    });

    const nameParts = parseName(data['Name'] || '');
    const birthday = parseDate(data['Birthday']);
    const died = parseDate(data['Died']);

    const relatives: Relative[] = [];
    const relativeCols: Array<[string, string]> = [
      ['Father', 'Father'],
      ['Mother', 'Mother'],
      ['Brother', 'Brother'],
      ['Sister', 'Sister'],
    ];

    for (const [col, relationship] of relativeCols) {
      const value = cleanValue(data[col] || '');
      if (value) {
        const relParts = parseName(value);
        relatives.push({
          FirstName: relParts.firstName,
          LastName: relParts.lastName,
          Relationship: relationship,
        });
      }
    }

    people.push({
      FirstName: nameParts.firstName,
      LastName: nameParts.lastName,
      Birthday: formatDate(birthday),
      Age: calculateAge(birthday, died),
      Relatives: relatives,
    });
  }

  console.log(JSON.stringify(people, null, 2));
}

main();