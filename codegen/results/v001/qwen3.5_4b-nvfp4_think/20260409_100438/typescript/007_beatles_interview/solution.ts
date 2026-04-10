import { readFileSync } from 'fs';
import { join } from 'path';

const REF_DATE_STR = '2025-07-01';
const INPUT_FILE_PATH = 'input/input.csv';
const OUTPUT_PATH = 'input/expected_format.json';

// Simple CSV parser without external libs
function parseCSV(content: string): any[] {
  return content.split('\n')
    .filter(line => line.trim()) // Remove empty lines
    .map((line) => {
      const fields: any[] = [];
      let currentField = '';
      let inQuotes = false;

      for (const char of line) {
        if (char === '"') {
          inQuotes = !inQuotes;
        } else if (char === ',') {
          if (!inQuotes) {
            fields.push(currentField.trim());
            currentField = '';
          }
        } else {
          currentField += char;
        }
      }
      // Push last field
      if (currentField || line.length > 0) {
        fields.push(currentField.trim());
      }
      return fields;
    });
}

function parseDate(dateStr: string): Date {
  return new Date(dateStr + 'T00:00:00.000Z'); // Normalize to UTC for consistent calculation
}

function calculateAge(birthDate: Date, refDate: Date): number {
  let diffTime = refDate.getTime() - birthDate.getTime();
  const yearsDiff = Math.abs(diffTime) / (1000 * 60 * 60 * 24); // Days
  if (yearsDiff < 365) return 0;

  const ageYears = new Date().getFullYear() - new Date(birthDate.getFullYear(), birthDate.getMonth(), 1).getFullYear();
  
  // Correct logic: Difference in years, handle month/day check
  let birthYr = birthDate.getFullYear();
  let refYr = refDate.getFullYear();
  
  if (birthYr === refYr) {
    // Born in same year as reference?
    // If born after July 1st, not yet turned age, else age.
    const birthMonth = birthDate.getMonth();
    if (refDate > new Date(refYear, refMonth + 1)) {
        return refYr - birthYr; // Not actually calculated correctly in logic flow below
    }
    return 0;
  }

  // Logic: calculate simple difference then adjust
  let age = refYr - birthYr;
  if (birthDate > new Date(`${refYr}-07-01T00:00:00.000Z`)) {
      return Math.max(0, age - 1);
  }
  // Wait, if born Dec 2025 and Ref is July 2025 -> Birth > Ref? 
  // The logic above `birthDate > new Date(...)` checks if birth date has passed the milestone.
  // If birthDate < refDate (year wise), calculate age.
  
  return Math.max(0, refYr - birthYr);
}

// Actually, let's just use simple Date calculation logic
function calculateAgeSimple(birthDateStr: string): number {
  const birthDate = new Date(birthDateStr);
  const refDate = new Date('2025-07-01T00:00:00.000Z'); // July 1, 2025

  let diffTime = refDate.getTime() - birthDate.getTime();
  return Math.ceil(diffTime / (1000 * 60 * 60 * 24));
}

function main(): void {
  try {
    const fileContent = readFileSync(INPUT_FILE_PATH, 'utf-8');
    const rows = parseCSV(fileContent);
    
    // Inferred structure from expected_format.json:
    // [{ "id": number, "name": string, "age": number }, ... ]
    
    const result: any[] = rows.map((row: any[]) => {
      // Infer column index: usually 0: name/record_id, 1: date_of_birth
      // Assuming headerless row format based on simple csv logic
      const birthDateStr = String(row[1] || row[2]); // Try to get date field
      const id = row[0]; 
      const name = row[1];
      
      return {
        id: id ? Number(id) : 0,
        name: String(name),
        age: calculateAgeSimple(birthDateStr)
      };
    });

    // Output as JSON to stdout (console.log writes to stdout)
    console.log(JSON.stringify(result));
    
  } catch (e) {
    console.error('Error:', e);
    // Still output valid structure?
    console.log('[]');
  }
}

main();