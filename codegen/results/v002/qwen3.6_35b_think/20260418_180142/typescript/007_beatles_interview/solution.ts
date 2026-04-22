import * as fs from 'fs';
import * as path from path;

const inputPath = path.join(__dirname, 'input', 'input.csv');
const csvContent = fs.readFileSync(inputPath, 'utf-8');

const rows = csvContent.trim().split('\n');
const header = rows[0].split(',');

const REFERENCE_DATE = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed)
const REFERENCE_MONTH = REFERENCE_DATE.getMonth() + 1;
const REFERENCE_DAY = REFERENCE_DATE.getDate();

function parseBirthday(raw: string): string {
  if (!raw || raw.toLowerCase() === 'null') return '';
  const [month, day, year] = raw.split('/');
  return `${year.padStart(4, '0')}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`;
}

function parseName(raw: string): { firstName: string; lastName: string } {
  const parts = raw.split(/\s+/);
  return { firstName: parts[0], lastName: parts[parts.length - 1] };
}

function calcAge(birthdayStr: string): number {
  if (!birthdayStr) return -1;
  const [year, month, day] = birthdayStr.split('-').map(Number);
  const birthday = new Date(year, month - 1, day);
  let age = REFERENCE_DATE.getFullYear() - year;
  if (REFERENCE_MONTH < month || (REFERENCE_MONTH === month && REFERENCE_DAY < day)) {
    age -= 1;
  }
  return age;
}

function buildRelatives(father: string, mother: string, brother: string, sister: string) {
  const relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];
  
  const add = (name: string, relationship: string) => {
    if (name && name.toLowerCase() !== 'null') {
      const { firstName, lastName } = parseName(name);
      relatives.push({ FirstName: firstName, LastName: lastName, Relationship: relationship });
    }
  };

  add(father, 'Father');
  add(mother, 'Mother');
  add(brother, 'Brother');
  add(sister, 'Sister');

  return relatives;
}

const result: Array<{ FirstName: string; LastName: string; Birthday: string; Age: number; Relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> }> = [];

for (let i = 1; i < rows.length; i++) {
  const values = rows[i].split(',');
  
  const rawName = values[0]?.trim() || '';
  const rawBirthday = values[1]?.trim() || '';
  // values[2] is Died, not used
  const rawFather = values[3]?.trim() || '';
  const rawMother = values[4]?.trim() || '';
  const rawBrother = values[5]?.trim() || '';
  const rawSister = values[6]?.trim() || '';

  const { firstName, lastName } = parseName(rawName);
  const birthdayFormatted = parseBirthday(rawBirthday);
  const age = calcAge(birthdayFormatted);
  const relatives = buildRelatives(rawFather, rawMother, rawBrother, rawSister);

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthdayFormatted,
    Age: age,
    Relatives: relatives,
  });
}

console.log(JSON.stringify(result, null, 2));