import * as fs from 'fs';
import * as path from 'path';

interface Person {
  FirstName: string;
  LastName: string;
  Birthday: string;
  Age: number;
  Relatives: Relative[];
}

interface Relative {
  FirstName: string;
  LastName: string;
  Relationship: string;
}

function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(/\s+/);
  if (parts.length === 1) {
    return { firstName: parts[0], lastName: '' };
  }
  return {
    firstName: parts[0],
    lastName: parts.slice(1).join(' ')
  };
}

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr.toLowerCase() === 'null') return null;
  
  const match = dateStr.match(/^(\d+)\/(\d+)\/(\d+)$/);
  if (!match) return null;
  
  const [, month, day, year] = match.map(Number);
  const date = new Date(year, month - 1, day);
  return isNaN(date.getTime()) ? null : date;
}

function formatBirthday(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function calculateAge(birthday: Date, deathDate: Date | null): number {
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025
  
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  
  const birthdayThisYear = new Date(referenceDate.getFullYear(), birthday.getMonth(), birthday.getDate());
  if (referenceDate < birthdayThisYear) {
    age--;
  }
  
  if (deathDate && deathDate < referenceDate) {
    let deathAge = deathDate.getFullYear() - birthday.getFullYear();
    const birthdayAtDeath = new Date(deathDate.getFullYear(), birthday.getMonth(), birthday.getDate());
    if (deathDate < birthdayAtDeath) {
      deathAge--;
    }
    return Math.max(0, deathAge);
  }
  
  return age;
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const csvContent = fs.readFileSync(inputPath, 'utf-8');
  
  const lines = csvContent.trim().split('\n');
  const headers = lines[0].split(',');
  
  const results: Person[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    const row: Record<string, string> = {};
    
    headers.forEach((header, index) => {
      row[header] = values[index] || '';
    });

    const nameParts = parseName(row['Name']);
    const birthday = parseDate(row['Birthday']);
    const deathDate = parseDate(row['Died']);
    const age = birthday ? calculateAge(birthday, deathDate) : 0;

    const relatives: Relative[] = [];
    
    if (row['Father'] && row['Father'].toLowerCase() !== 'null') {
      const relative = parseName(row['Father']);
      relatives.push({
        FirstName: relative.firstName,
        LastName: relative.lastName,
        Relationship: 'Father'
      });
    }
    
    if (row['Mother'] && row['Mother'].toLowerCase() !== 'null') {
      const relative = parseName(row['Mother']);
      relatives.push({
        FirstName: relative.firstName,
        LastName: relative.lastName,
        Relationship: 'Mother'
      });
    }
    
    if (row['Brother'] && row['Brother'].toLowerCase() !== 'null') {
      const relative = parseName(row['Brother']);
      relatives.push({
        FirstName: relative.firstName,
        LastName: relative.lastName,
        Relationship: 'Brother'
      });
    }
    
    if (row['Sister'] && row['Sister'].toLowerCase() !== 'null') {
      const relative = parseName(row['Sister']);
      relatives.push({
        FirstName: relative.firstName,
        LastName: relative.lastName,
        Relationship: 'Sister'
      });
    }
    
    results.push({
      FirstName: nameParts.firstName,
      LastName: nameParts.lastName,
      Birthday: birthday ? formatBirthday(birthday) : '',
      Age: age,
      Relatives: relatives
    });
  }

  console.log(JSON.stringify(results));
}

main();