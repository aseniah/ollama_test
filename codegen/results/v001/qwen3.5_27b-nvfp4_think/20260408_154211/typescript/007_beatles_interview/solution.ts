import * as fs from 'fs';
import * as path from 'path';

// Read input files
const csvPath = path.join(__dirname, '../input/input.csv');
const formatPath = path.join(__dirname, '../input/expected_format.json');

const csvContent = fs.readFileSync(csvPath, 'utf-8');
const expectedFormat = JSON.parse(fs.readFileSync(formatPath, 'utf-8'));

// Parse CSV content into rows
function parseCSV(content: string): string[][] {
  const lines = content.trim().split('\n');
  const result: string[][] = [];
  
  for (const line of lines) {
    const row: string[] = [];
    let current = '';
    let inQuotes = false;
    
    for (let i = 0; i < line.length; i++) {
      const char = line[i];
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === ',' && !inQuotes) {
        row.push(current.trim());
        current = '';
      } else {
        current += char;
      }
    }
    row.push(current.trim());
    result.push(row);
  }
  
  return result;
}

// Calculate age as of July 1, 2025
function calculateAge(birthDateString: string): number {
  const referenceDate = new Date('2025-07-01');
  
  let birthDate: Date;
  if (birthDateString.includes('/')) {
    // Format: MM/DD/YYYY
    const [month, day, year] = birthDateString.split('/').map(Number);
    birthDate = new Date(year, month - 1, day);
  } else if (birthDateString.includes('-')) {
    // Format: YYYY-MM-DD or DD-MM-YYYY
    const parts = birthDateString.split('-').map(Number);
    if (parts[0] > 2000) {
      // YYYY-MM-DD
      birthDate = new Date(parts[0], parts[1] - 1, parts[2]);
    } else {
      // DD-MM-YYYY or other format, try to infer
      birthDate = new Date(parts[2], parts[1] - 1, parts[0]);
    }
  } else {
    throw new Error(`Unknown date format: ${birthDateString}`);
  }
  
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }
  
  return age;
}

// Infer column mapping from expected format
const firstRow = parseCSV(csvContent);
const headers = firstRow[0];
const dataRows = firstRow.slice(1);

// Map CSV columns to JSON properties based on expected format
const columnNameMap: Record<string, string> = {};

// Common date columns that might be in the CSV
const dateColumnNames = ['birth_date', 'birthdate', 'dob', 'date_of_birth', 'Date of Birth'];

for (const header of headers) {
  const lowerHeader = header.toLowerCase().trim();
  
  for (const name of dateColumnNames) {
    if (lowerHeader.includes(name)) {
      columnNameMap[lowerHeader] = 'birth_date';
      break;
    }
  }
  
  // Map other common names
  if (lowerHeader.includes('name') || lowerHeader === 'full_name' || lowerHeader === 'name_full') {
    columnNameMap[lowerHeader] = 'name';
  } else if (lowerHeader.includes('city')) {
    columnNameMap[lowerHeader] = 'city';
  } else if (lowerHeader.includes('country')) {
    columnNameMap[lowerHeader] = 'country';
  } else if (lowerHeader.includes('email')) {
    columnNameMap[lowerHeader] = 'email';
  } else if (lowerHeader.includes('phone') || lowerHeader.includes('telephone')) {
    columnNameMap[lowerHeader] = 'phone';
  }
}

// Build output based on expected format
const output: any[] = [];

for (const row of dataRows) {
  const entry: Record<string, unknown> = {};
  
  // Map each header to the corresponding value
  for (let i = 0; i < headers.length; i++) {
    const header = headers[i];
    const mappedName = columnNameMap[header.toLowerCase().trim()] || header.toLowerCase().trim();
    
    let value: any = row[i]?.trim() || '';
    
    // Calculate age if this is a birth date column
    if (mappedName === 'birth_date' && value) {
      entry['age'] = calculateAge(value);
      // Also include original birth_date if expected format shows it
      if (expectedFormat.length > 0 && expectedFormat[0].hasOwnProperty('birth_date')) {
        entry['birth_date'] = value;
      }
    } else if (value) {
      entry[mappedName] = value;
    }
  }
  
  // If expected format has specific fields, ensure they're present in correct order
  if (expectedFormat.length > 0) {
    const expectedFields = Object.keys(expectedFormat[0]);
    const formattedEntry: Record<string, unknown> = {};
    
    for (const field of expectedFields) {
      if (entry[field] !== undefined) {
        formattedEntry[field] = entry[field];
      } else if (field === 'age' && entry['birth_date']) {
        // We already calculated age above
      }
    }
    
    output.push(formattedEntry);
  } else {
    output.push(entry);
  }
}

// Output JSON array to stdout
console.log(JSON.stringify(output, null, 2));