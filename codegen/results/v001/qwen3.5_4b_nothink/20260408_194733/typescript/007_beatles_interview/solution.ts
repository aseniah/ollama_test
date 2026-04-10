import * as fs from 'fs';

// Read input CSV
const inputData = fs.readFileSync('input/input.csv', 'utf-8');
const lines = inputData.trim().split('\n').filter(line => line.trim() !== '');

// Read expected format to infer column structure and output format
const expected = fs.readFileSync('input/expected_format.json', 'utf-8');
const expectedOutput = JSON.parse(expected);

// Infer column names from input
const headers = lines[0].split(',').map(h => h.trim().toLowerCase());

// Infer output fields from expected
const outputFields = Object.keys(expectedOutput[0]);

// Helper to find index by name (case-insensitive)
function findIndex(arr: string[], name: string): number {
  return arr.findIndex(h => h.toLowerCase() === name.toLowerCase());
}

// Helper to parse date
function parseDate(dateStr: string): Date {
  // Handle various date formats. Common ones: YYYY-MM-DD, DD/MM/YYYY, YYYY/MM/DD, Month DD, YYYY
  dateStr = dateStr.trim();
  const dateMatch = dateStr.match(/(\d{1,2})[\/-](\d{1,2})[\/-](\d{4})/);
  if (dateMatch) {
    const [_, day, month, year] = dateMatch;
    return new Date(+year, +month - 1, +day);
  }
  
  // Try DD/MM/YYYY
  const match2 = dateStr.match(/(\d{1,2})[\/](\d{1,2})[\/](\d{4})/);
  if (match2) {
    const [_, day, month, year] = match2;
    return new Date(+year, +month - 1, +day);
  }
  
  // Try Month DD, YYYY (e.g., "July 15, 2023")
  const match3 = dateStr.match(/^([A-Za-z]+)\s+(\d{1,2}),?\s+(\d{4})/);
  if (match3) {
    const [_, monthName, day, year] = match3;
    const months: Record<string, number> = {
      january: 0, february: 1, march: 2, april: 3, may: 4, june: 5,
      july: 6, august: 7, september: 8, october: 9, november: 10, december: 11
    };
    return new Date(+year, months[monthName.toLowerCase()], +day);
  }
  
  // Try just DD MM YYYY or DD-MM-YYYY
  const match4 = dateStr.match(/(\d{1,2})[\/\-](\d{1,2})[\/\-](\d{4})/);
  if (match4) {
    const [_, day, month, year] = match4;
    return new Date(+year, +month - 1, +day);
  }

  throw new Error(`Date parsing failed for: ${dateStr}`);
}

// Calculate age as of July 1, 2025
const targetDate = new Date(2025, 6, 1); // July 1, 2025

const result: any[] = [];

for (const line of lines) {
  const values = line.split(',').map(v => v.trim());
  
  // Find indices of relevant columns
  const birthDateIndex = findIndex(headers, 'birth_date');
  const firstNameIndex = findIndex(headers, 'first_name');
  const lastNameIndex = findIndex(headers, 'last_name');
  const sexIndex = findIndex(headers, 'sex');
  const nationalityIndex = findIndex(headers, 'nationality');
  const countryOfBirthIndex = findIndex(headers, 'country_of_birth');
  const stateOfBirthIndex = findIndex(headers, 'state_of_birth');
  
  // Ensure we have all necessary columns
  if (birthDateIndex === -1 || firstNameIndex === -1) {
    // Skip line if critical fields missing
    continue;
  }

  const birthDateStr = values[birthDateIndex] || '';
  const birthDate = parseDate(birthDateStr);
  
  const firstName = values[firstNameIndex] || '';
  const lastName = values[lastNameIndex] || '';
  
  // Calculate age
  let age = targetDate.getFullYear() - birthDate.getFullYear();
  if (targetDate < birthDate) {
    age--;
  }
  // Also check for birthday passed in current year
  if (birthDate.getMonth() > targetDate.getMonth() || 
      (birthDate.getMonth() === targetDate.getMonth() && birthDate.getDate() >= targetDate.getDate())) {
    age--;
  }
  
  // Prepare output object
  const outputObj: any = {
    first_name: firstName,
    last_name: lastName,
    age: age,
  };

  // Add other fields if they exist in expected format
  const extraFields = outputFields.filter(f => f !== 'first_name' && f !== 'last_name' && f !== 'age');
  
  for (const field of extraFields) {
    let value: any;
    
    switch (field) {
      case 'sex':
        value = values[sexIndex] || '';
        break;
      case 'nationality':
        value = values[nationalityIndex] || '';
        break;
      case 'country_of_birth':
        value = values[countryOfBirthIndex] || '';
        break;
      case 'state_of_birth':
        value = values[stateOfBirthIndex] || '';
        break;
      default:
        // Try to find similar field if exact match not found (case-insensitive)
        const similarField = headers.findIndex(h => h.toLowerCase().includes(field.toLowerCase()));
        if (similarField !== -1) {
          value = values[similarField] || '';
        } else {
          value = '';
        }
        break;
    }
    
    // Clean up special characters in strings
    if (typeof value === 'string') {
      value = value.replace(/,/g, ','); // Remove commas that might be part of date separators if misparsed
    }
    
    outputObj[field] = value;
  }
  
  result.push(outputObj);
}

// Output JSON
console.log(JSON.stringify(result));