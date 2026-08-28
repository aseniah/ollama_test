import * as fs from 'fs';
import * as path from 'path';

const REFERENCE_DATE = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed)

function parseCSV(content: string): any[] {
  const lines = content.trim().split('\n');
  const data: any[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;
    
    const values = line.split(',');
    data.push({
      name: values[0],
      birthday: values[1],
      died: values[2] === 'null' ? null : values[2],
      father: values[3] === 'null' ? null : values[3],
      mother: values[4] === 'null' ? null : values[4],
      brother: values[5] === 'null' ? null : values[5],
      sister: values[6] === 'null' ? null : values[6]
    });
  }
  
  return data;
}

function parseBirthday(birthdayStr: string): string {
  const parts = birthdayStr.split('/');
  const month = parts[0].padStart(2, '0');
  const day = parts[1].padStart(2, '0');
  const year = parts[2];
  return `${year}-${month}-${day}`;
}

function calculateAge(birthdayStr: string): number {
  const parts = birthdayStr.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  
  const birthday = new Date(year, month - 1, day);
  const age = REFERENCE_DATE.getFullYear() - birthday.getFullYear();
  
  if (month > REFERENCE_DATE.getMonth() || 
      (month === REFERENCE_DATE.getMonth() && day > REFERENCE_DATE.getDate())) {
    return age - 1;
  }
  
  return age;
}

function parseName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(' ');
  return { firstName: parts[0], lastName: parts.slice(1).join(' ') };
}

function parseRelativeName(name: string): { firstName: string; lastName: string } {
  const parts = name.trim().split(' ');
  return { firstName: parts[0], lastName: parts.slice(1).join(' ') };
}

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  const csvData = parseCSV(content);
  
  const result: any[] = [];
  
  for (const row of csvData) {
    const { firstName, lastName } = parseName(row.name);
    const birthdayFormatted = parseBirthday(row.birthday);
    const age = calculateAge(row.birthday);
    
    const relatives: any[] = [];
    
    if (row.father) {
      const { firstName: fatherFirstName, lastName: fatherLastName } = parseRelativeName(row.father);
      relatives.push({ firstName: fatherFirstName, lastName: fatherLastName, relationship: 'Father' });
    }
    
    if (row.mother) {
      const { firstName: motherFirstName, lastName: motherLastName } = parseRelativeName(row.mother);
      relatives.push({ firstName: motherFirstName, lastName: motherLastName, relationship: 'Mother' });
    }
    
    if (row.brother) {
      const { firstName: brotherFirstName, lastName: brotherLastName } = parseRelativeName(row.brother);
      relatives.push({ firstName: brotherFirstName, lastName: brotherLastName, relationship: 'Brother' });
    }
    
    if (row.sister) {
      const { firstName: sisterFirstName, lastName: sisterLastName } = parseRelativeName(row.sister);
      relatives.push({ firstName: sisterFirstName, lastName: sisterLastName, relationship: 'Sister' });
    }
    
    result.push({ firstName, lastName, birthday: birthdayFormatted, age, relatives });
  }
  
  console.log(JSON.stringify(result));
}

main();