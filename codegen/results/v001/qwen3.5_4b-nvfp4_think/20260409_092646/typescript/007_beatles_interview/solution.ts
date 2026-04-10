import { readFileSync } from 'fs';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const inputPath = fileURLToPath(import.meta.url);
const filePath = dirname(inputPath) + '/input/input.csv';
const expectedFormatPath = dirname(filePath) + '/input/expected_format.json';

function parseCSV(content: string): any[] {
  const lines = content.trim().split('\n');
  if (lines.length === 0) return [];
  
  const headers = lines[0].split(',').map(h => h.trim());
  const rows: any[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    const row: any = {};
    for (let j = 0; j < headers.length; j++) {
      if (values[j]) {
        row[headers[j]] = values[j];
      }
    }
    rows.push(row);
  }
  
  return rows;
}

function parseDate(dateStr: string): Date | null {
  if (!dateStr || dateStr.trim() === '') return null;
  
  const dateRegex = /^(?:(\d{2})(-\d{2})|(\d{2}-(\d{1,2})))/;
  const match = dateStr.match(dateRegex);
  
  if (match) {
    let year, month, day;
    
    if (match[3]) {
      // Format: MM-DD-YYYY or DD-MM-YYYY
      year = parseInt(match[4] || '0000');
      month = parseInt(match[3]);
      day = parseInt(match[2]);
    } else {
      // Format: YYYY-MM-DD
      year = parseInt(match[1]);
      month = parseInt(match[2]);
      day = parseInt(match[3]);
    }
    
    return new Date(year, month - 1, day);
  }
  
  return null;
}

function calculateAgeAsOf(birthDate: Date, targetDate: Date): number {
  let age = targetDate.getFullYear() - birthDate.getFullYear();
  
  const targetBirthday = new Date(targetDate.getFullYear(), targetDate.getMonth(), targetDate.getDate());
  const currentBirthday = new Date(birthDate.getFullYear(), birthDate.getMonth(), birthDate.getDate());
  
  if (currentBirthday > targetBirthday) {
    age--;
  }
  
  return age;
}

const expectedFormat = JSON.parse(readFileSync(expectedFormatPath, 'utf-8'));

// Define target date: July 1, 2025
const TARGET_DATE = new Date(2025, 6, 1);

const csvContent = readFileSync(filePath, 'utf-8');
const data = parseCSV(csvContent);

// Ensure the output matches expected format structure
const columns = Object.keys(data[0] || {});

const result = data.map(row => {
  const age = calculateAgeAsOf(parseDate(row['DOB'] || row.birthDate || row.date_of_birth || row.DOB), TARGET_DATE);
  
  return {
    _name: row['name'] || row['Name'] || row['Full Name'] || row['FullName'],
    _age: age,
    _date: new Date(TARGET_DATE).toDateString(),
    _values: row as any
  };
});

console.log(JSON.stringify(result, null, 2));