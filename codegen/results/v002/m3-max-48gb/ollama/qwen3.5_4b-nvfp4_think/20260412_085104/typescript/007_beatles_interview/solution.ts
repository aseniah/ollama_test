import * as fs from 'fs';

const referenceDate = new Date('2025-07-01');

function parseDate(dateStr: string): Date {
  const [month, day, year] = dateStr.split('/').map(Number);
  return new Date(year, month - 1, day);
}

function formatDate(date: Date): string {
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${pad(date.getFullYear())}-${pad(date.getMonth() + 1)}-${pad(date.getDate())}`;
}

function parseName(name: string): { first: string; last: string } {
  const parts = name.split(' ');
  return {
    first: parts[0],
    last: parts.slice(parts.length - 1).join(' ')
  };
}

function calculateAge(birthday: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const refBirthday = new Date(referenceDate.getFullYear(), birthday.getMonth(), birthday.getDate());
  if (refBirthday < referenceDate) {
    age--;
  }
  return Math.max(0, age);
}

function parseRelative(rel: string): { first: string; last: string; relationship: string } {
  if (!rel || rel === 'null') return null as any;
  const parts = rel.split(' ');
  return {
    first: parts[0],
    last: parts.length > 1 ? parts.slice(parts.length - 1).join(' ') : '',
    relationship: ''
  };
}

const csvContent = fs.readFileSync('input/input.csv', 'utf-8');
const lines = csvContent.toString().split('\n');
const headers = lines[0].split(',');

const data: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const values = lines[i].split(',');
  
  const [name, birthdayStr, diedStr, fatherStr, motherStr, brotherStr, sisterStr] = values.map(v => v.trim());
  
  if (!name) continue;

  const birthDate = parseDate(birthdayStr);
  
  let age: number;
  if (diedStr === 'null') {
    age = calculateAge(birthDate, referenceDate);
  } else {
    const deathDate = parseDate(diedStr);
    const deathBirthday = new Date(deathDate.getFullYear(), birthDate.getMonth(), birthDate.getDate());
    if (deathBirthday < deathDate) {
      age = deathDate.getFullYear() - deathDate.getFullYear();
    } else {
      age = Math.max(0, deathDate.getFullYear() - birthDate.getFullYear() - 1);
    }
  }

  const relatives: any[] = [];
  
  if (fatherStr !== 'null') {
    const rel = parseRelative(fatherStr);
    relatives.push({
      first: rel.first,
      last: rel.last || '',
      relationship: 'Father'
    });
  }
  if (motherStr !== 'null') {
    const rel = parseRelative(motherStr);
    relatives.push({
      first: rel.first,
      last: rel.last || '',
      relationship: 'Mother'
    });
  }
  if (brotherStr !== 'null') {
    const rel = parseRelative(brotherStr);
    relatives.push({
      first: rel.first,
      last: rel.last || '',
      relationship: 'Brother'
    });
  }
  if (sisterStr !== 'null') {
    const rel = parseRelative(sisterStr);
    relatives.push({
      first: rel.first,
      last: rel.last || '',
      relationship: 'Sister'
    });
  }

  data.push({
    FirstName: name.split(' ')[0],
    LastName: parseName(name).last,
    Birthday: formatDate(birthDate),
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(data, null, 2));