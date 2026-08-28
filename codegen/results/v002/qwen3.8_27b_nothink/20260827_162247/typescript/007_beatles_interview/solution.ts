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

function parseDate(dateStr: string): Date {
  const parts = dateStr.split('/');
  // MM/DD/YYYY
  const month = parseInt(parts[0], 10);
  const day = parseInt(parts[1], 10);
  const year = parseInt(parts[2], 10);
  return new Date(year, month - 1, day);
}

function calculateAge(birthDate: Date, refDate: Date): number {
  let age = refDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = refDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && refDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

function formatBirthday(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

function nameToFirstAndLast(name: string): { first: string; last: string } {
  const parts = name.trim().split(/\s+/);
  const firstName = parts[0];
  const lastName = parts.length > 1 ? parts[parts.length - 1] : '';
  return { first: firstName, last: lastName };
}

function parseCSV(csvContent: string): string[][] {
  const lines = csvContent.trim().split('\n');
  const rows: string[][] = [];
  
  for (const line of lines) {
    // Simple CSV parsing - handles quoted fields
    const row: string[] = [];
    let current = '';
    let inQuotes = false;
    
    for (let i = 0; i < line.length; i++) {
      const char = line[i];
      
      if (inQuotes) {
        if (char === '"') {
          if (i + 1 < line.length && line[i + 1] === '"') {
            current += '"';
            i++;
          } else {
            inQuotes = false;
          }
        } else {
          current += char;
        }
      } else {
        if (char === '"') {
          inQuotes = true;
        } else if (char === ',') {
          row.push(current);
          current = '';
        } else {
          current += char;
        }
      }
    }
    row.push(current);
    rows.push(row);
  }
  
  return rows;
}

function main() {
  const inputPath = path.resolve('input/input.csv');
  const content = fs.readFileSync(inputPath, 'utf-8');
  
  const rows = parseCSV(content);
  const header = rows[0];
  const dataRows = rows.slice(1);
  
  const referenceDate = new Date(2025, 6, 1); // July 1, 2025 (month is 0-indexed, so 6 = July)
  
  const result: Person[] = [];
  
  for (const row of dataRows) {
    const record: Record<string, string> = {};
    for (let i = 0; i < header.length; i++) {
      record[header[i]] = row[i] || '';
    }
    
    const fullName = record['Name'];
    const { first: firstName, last: lastName } = nameToFirstAndLast(fullName);
    
    const birthDate = parseDate(record['Birthday']);
    const age = calculateAge(birthDate, referenceDate);
    const birthdayStr = formatBirthday(birthDate);
    
    const relatives: Relative[] = [];
    
    const father = record['Father'];
    if (father && father !== 'null') {
      const { first: fFirst, last: fLast } = nameToFirstAndLast(father);
      relatives.push({
        FirstName: fFirst,
        LastName: fLast,
        Relationship: 'Father'
      });
    }
    
    const mother = record['Mother'];
    if (mother && mother !== 'null') {
      const { first: mFirst, last: mLast } = nameToFirstAndLast(mother);
      relatives.push({
        FirstName: mFirst,
        LastName: mLast,
        Relationship: 'Mother'
      });
    }
    
    const brother = record['Brother'];
    if (brother && brother !== 'null') {
      const { first: bFirst, last: bLast } = nameToFirstAndLast(brother);
      relatives.push({
        FirstName: bFirst,
        LastName: bLast,
        Relationship: 'Brother'
      });
    }
    
    const sister = record['Sister'];
    if (sister && sister !== 'null') {
      const { first: sFirst, last: sLast } = nameToFirstAndLast(sister);
      relatives.push({
        FirstName: sFirst,
        LastName: sLast,
        Relationship: 'Sister'
      });
    }
    
    result.push({
      FirstName: firstName,
      LastName: lastName,
      Birthday: birthdayStr,
      Age: age,
      Relatives: relatives
    });
  }
  
  process.stdout.write(JSON.stringify(result, null, 2));
}

main();