import * as fs from 'fs';
import * as path from 'path';

function parseCsvLine(line: string): string[] {
  const result: string[] = [];
  let current = '';
  let inQuotes = false;
  for (let i = 0; i < line.length; i++) {
    const char = line[i];
    if (char === '"') {
      inQuotes = !inQuotes;
    } else if (char === ',' && !inQuotes) {
      result.push(current.trim());
      current = '';
    } else {
      current += char;
    }
  }
  result.push(current.trim());
  return result;
}

function readCsv(filePath: string): string[][] {
  const content = fs.readFileSync(filePath, 'utf8').trim();
  const lines = content.split('\n');
  const header = parseCsvLine(lines[0]);
  const rows: string[][] = [];
  for (let i = 1; i < lines.length; i++) {
    if (lines[i].trim() === '') continue;
    const fields = parseCsvLine(lines[i]);
    rows.push(fields);
  }
  return [header, ...rows];
}

function calculateAge(birthday: string, referenceDate: Date): number {
  // birthday format: M/D/YYYY
  const parts = birthday.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  
  const refYear = referenceDate.getFullYear();
  const refMonth = referenceDate.getMonth() + 1; // 1-12
  const refDay = referenceDate.getDate();
  
  let age = refYear - year;
  if (refMonth < month || (refMonth === month && refDay < day)) {
    age--;
  }
  return age;
}

function toIsoDate(birthday: string): string {
  const parts = birthday.split('/');
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  const mm = month.toString().padStart(2, '0');
  const dd = day.toString().padStart(2, '0');
  return `${year}-${mm}-${dd}`;
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
  const inputPath = path.join('input', 'input.csv');
  const [header, ...rows] = readCsv(inputPath);
  
  // Reference date: July 1, 2025
  const referenceDate = new Date(2025, 6, 1); // Month is 0-indexed, so 6 = July
  
  const result: Person[] = [];
  
  for (const row of rows) {
    const record: Record<string, string> = {};
    for (let i = 0; i < header.length; i++) {
      record[header[i]] = row[i] || '';
    }
    
    const fullName = record['Name'].trim();
    const nameParts = fullName.split(' ');
    const firstName = nameParts[0];
    const lastName = nameParts[nameParts.length - 1];
    
    const birthday = record['Birthday'].trim();
    const age = calculateAge(birthday, referenceDate);
    const isoBirthday = toIsoDate(birthday);
    
    const relatives: Relative[] = [];
    
    if (record['Father'] && record['Father'] !== 'null') {
      const fatherName = record['Father'].trim();
      const fatherParts = fatherName.split(' ');
      relatives.push({
        FirstName: fatherParts[0],
        LastName: fatherParts[fatherParts.length - 1],
        Relationship: 'Father'
      });
    }
    
    if (record['Mother'] && record['Mother'] !== 'null') {
      const motherName = record['Mother'].trim();
      const motherParts = motherName.split(' ');
      relatives.push({
        FirstName: motherParts[0],
        LastName: motherParts[motherParts.length - 1],
        Relationship: 'Mother'
      });
    }
    
    if (record['Brother'] && record['Brother'] !== 'null') {
      const brotherName = record['Brother'].trim();
      const brotherParts = brotherName.split(' ');
      relatives.push({
        FirstName: brotherParts[0],
        LastName: brotherParts[brotherParts.length - 1],
        Relationship: 'Brother'
      });
    }
    
    if (record['Sister'] && record['Sister'] !== 'null') {
      const sisterName = record['Sister'].trim();
      const sisterParts = sisterName.split(' ');
      relatives.push({
        FirstName: sisterParts[0],
        LastName: sisterParts[sisterParts.length - 1],
        Relationship: 'Sister'
      });
    }
    
    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: isoBirthday,
      Age: age,
      Relatives: relatives
    });
  }
  
  process.stdout.write(JSON.stringify(result, null, 2));
}

main();