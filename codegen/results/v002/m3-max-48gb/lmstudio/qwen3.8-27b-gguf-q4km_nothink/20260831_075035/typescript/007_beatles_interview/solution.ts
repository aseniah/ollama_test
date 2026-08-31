import * as fs from 'fs';
import * as path from 'path';

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

function parseName(fullName: string): { firstName: string; lastName: string } {
  const parts = fullName.trim().split(/\s+/);
  const lastName = parts[parts.length - 1];
  const firstName = parts.slice(0, -1).join(' ');
  return { firstName, lastName };
}

function calculateAge(birthday: string, referenceDate: Date): number {
  const [month, day, year] = birthday.split('/').map(Number);
  const birth = new Date(year, month - 1, day);
  
  let age = referenceDate.getFullYear() - birth.getFullYear();
  const refMonth = referenceDate.getMonth();
  const refDay = referenceDate.getDate();
  const birthMonth = birth.getMonth();
  const birthDay = birth.getDate();
  
  if (refMonth < birthMonth || (refMonth === birthMonth && refDay < birthDay)) {
    age--;
  }
  
  return age;
}

function formatBirthday(birthday: string): string {
  const [month, day, year] = birthday.split('/').map(Number);
  const mm = month.toString().padStart(2, '0');
  const dd = day.toString().padStart(2, '0');
  const yyyy = year.toString().padStart(4, '0');
  return `${yyyy}-${mm}-${dd}`;
}

function main(): void {
  const inputPath = path.join(__dirname, 'input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  const lines = content.trim().split('\n');
  
  const header = lines[0].split(',');
  const people: Person[] = [];
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed, so 6 = July)
  
  for (let i = 1; i < lines.length; i++) {
    const fields = lines[i].split(',');
    const row: Record<string, string> = {};
    for (let j = 0; j < header.length; j++) {
      row[header[j]] = (fields[j] || '').trim();
    }
    
    const name = row['Name'];
    const birthday = row['Birthday'];
    
    const { firstName, lastName } = parseName(name);
    const age = calculateAge(birthday, referenceDate);
    const formattedBirthday = formatBirthday(birthday);
    
    const relatives: Relative[] = [];
    
    const relationshipFields: [string, string][] = [
      ['Father', 'Father'],
      ['Mother', 'Mother'],
      ['Brother', 'Brother'],
      ['Sister', 'Sister']
    ];
    
    for (const [field, relationship] of relationshipFields) {
      const value = row[field];
      if (value && value !== 'null' && value.trim() !== '') {
        const { firstName: relFirstName, lastName: relLastName } = parseName(value);
        relatives.push({
          FirstName: relFirstName,
          LastName: relLastName,
          Relationship: relationship
        });
      }
    }
    
    people.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: formattedBirthday,
      Age: age,
      Relatives: relatives
    });
  }
  
  process.stdout.write(JSON.stringify(people, null, 2));
}

main();