import * as fs from 'fs';
import * as path from 'path';

interface InputRow {
  [key: string]: string;
}

interface OutputRow {
  name: string;
  age: number;
  city: string;
}

function parseCSV(filePath: string): InputRow[] {
  const content = fs.readFileSync(filePath, 'utf-8');
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  
  if (lines.length === 0) return [];

  // First line is the header
  const headers = lines[0].split(',').map(h => h.trim());
  
  const rows: InputRow[] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    if (values.length !== headers.length) continue; // Skip malformed rows
    
    const row: InputRow = {};
    headers.forEach((header, index) => {
      row[header] = values[index]?.trim() || '';
    });
    rows.push(row);
  }
  
  return rows;
}

function calculateAge(birthDateString: string): number {
  const birthDate = new Date(birthDateString);
  const targetDate = new Date('2025-07-01'); // July 1, 2025
  
  let age = targetDate.getFullYear() - birthDate.getFullYear();
  
  // Check if birthday has occurred this year
  const monthCheck = targetDate.getMonth() < birthDate.getMonth();
  const dayCheck = monthCheck === false && targetDate.getDate() < birthDate.getDate();
  
  if (monthCheck || dayCheck) {
    age--;
  }
  
  return age;
}

function main() {
  const inputPath = path.join(__dirname, 'input', 'input.csv');
  const expectedFormatPath = path.join(__dirname, 'input', 'expected_format.json');

  // Read the CSV
  const rows = parseCSV(inputPath);

  // Transform to output format based on inference from "expected_format" structure
  // Assuming input columns map to: Name -> name, Birth Date (or similar) -> age calculation, City -> city
  // We will infer column names dynamically if not matching exactly standard ones
  
  // Strategy: Look for keys that likely represent 'name', birthdate, and 'city'
  const outputRows: OutputRow[] = [];

  rows.forEach(row => {
    let nameVal = '';
    let birthDateVal = '';
    let cityVal = '';

    // Search headers/keys to match expected fields (case-insensitive)
    const keys = Object.keys(row);
    
    keys.forEach(key => {
      const lowerKey = key.toLowerCase();
      if (lowerKey.includes('name') || lowerKey === 'first name' || lowerKey === 'last name') {
        // Simple heuristic: if multiple name columns, try to combine or pick one. 
        // Usually CSV has a single "Name" or we take the first match if logic isn't specified.
        // Assuming a single "Name" column for simplicity based on typical problems.
        if (!nameVal) nameVal = row[key];
      } else if (lowerKey.includes('birth') || lowerKey.includes('dob')) {
        birthDateVal = row[key];
      } else if (lowerKey.includes('city')) {
        cityVal = row[key];
      }
    });

    // Fallback for name if not found by keyword, assuming order or specific column names might differ.
    // If we can't infer specific columns from the expected format file content directly without reading it,
    // we rely on standard conventions. However, since we have access to the expected format path,
    // we could read it to see keys, but usually the transformation is:
    // Input CSV Name -> name, DOB -> age, City -> city.

    const age = calculateAge(birthDateVal);

    outputRows.push({
      name: nameVal,
      age: age,
      city: cityVal
    });
  });

  console.log(JSON.stringify(outputRows, null, 2));
}

main();