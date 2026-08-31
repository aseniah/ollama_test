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

function parseDate(dateStr: string): { year: number; month: number; day: number } {
  const parts = dateStr.split('/');
  return { month: parseInt(parts[0], 10), day: parseInt(parts[1], 10), year: parseInt(parts[2], 10) };
}

function formatDate(year: number, month: number, day: number): string {
  const mm = month.toString().padStart(2, '0');
  const dd = day.toString().padStart(2, '0');
  return `${year}-${mm}-${dd}`;
}

function calculateAge(birthYear: number, birthMonth: number, birthDay: number, referenceYear: number, referenceMonth: number, referenceDay: number): number {
  let age = referenceYear - birthYear;
  if (referenceMonth < birthMonth || (referenceMonth === birthMonth && referenceDay < birthDay)) {
    age--;
  }
  return age;
}

function splitName(fullName: string): { first: string; last: string } {
  const parts = fullName.trim().split(/\s+/);
  const first = parts[0];
  const last = parts[parts.length - 1];
  return { first, last };
}

function main() {
  const csvContent = fs.readFileSync('input/input.csv', 'utf-8').trim();
  const lines = csvContent.split('\n');
  const headers = lines[0].split(',');
  
  const result: Person[] = [];

  for (let i = 1; i < lines.length; i++) {
    const fields = lines[i].split(',').map(f => f.trim());
    
    const [name, birthday, died, father, mother, brother, sister] = fields;
    
    const { first, last } = splitName(name);
    
    const bd = parseDate(birthday);
    const birthdayFormatted = formatDate(bd.year, bd.month, bd.day);
    
    let age: number;
    if (died === 'null' || died === '') {
      // Alive: age as of July 1, 2025
      age = calculateAge(bd.year, bd.month, bd.day, 2025, 7, 1);
    } else {
      // Deceased: age at death
      const dd = parseDate(died);
      age = calculateAge(bd.year, bd.month, bd.day, dd.year, dd.month, dd.day);
    }
    
    const relatives: Relative[] = [];
    
    if (father !== 'null' && father !== '') {
      const { first: fn, last: ln } = splitName(father);
      relatives.push({ FirstName: fn, LastName: ln, Relationship: 'Father' });
    }
    if (mother !== 'null' && mother !== '') {
      const { first: fn, last: ln } = splitName(mother);
      relatives.push({ FirstName: fn, LastName: ln, Relationship: 'Mother' });
    }
    if (brother !== 'null' && brother !== '') {
      const { first: fn, last: ln } = splitName(brother);
      relatives.push({ FirstName: fn, LastName: ln, Relationship: 'Brother' });
    }
    if (sister !== 'null' && sister !== '') {
      const { first: fn, last: ln } = splitName(sister);
      relatives.push({ FirstName: fn, LastName: ln, Relationship: 'Sister' });
    }
    
    result.push({
      FirstName: first,
      LastName: last,
      Birthday: birthdayFormatted,
      Age: age,
      Relatives: relatives
    });
  }
  
  console.log(JSON.stringify(result, null, 2));
}

main();