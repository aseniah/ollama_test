import * as fs from 'fs';
import * as path from 'path';

function parseCSV(content: string): string[][] {
  const lines = content.trim().split('\n');
  const result: string[][] = [];

  for (const line of lines) {
    const row: string[] = [];
    let current = '';
    let inQuotes = false;

    for (const char of line) {
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

function calculateAge(dob: string): number {
  const targetDate = new Date('2025-07-01T00:00:00');
  const birthDate = new Date(dob);
  
  if (isNaN(birthDate.getTime())) {
    return 0;
  }

  let age = targetDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = targetDate.getMonth() - birthDate.getMonth();
  
  if (monthDiff < 0 || (monthDiff === 0 && targetDate.getDate() < birthDate.getDate())) {
    age--;
  }

  return age;
}

function main() {
  const inputPath = path.join(__dirname, 'input', 'input.csv');
  const formatPath = path.join(__dirname, 'input', 'expected_format.json');

  try {
    const csvContent = fs.readFileSync(inputPath, 'utf-8');
    const formatContent = fs.readFileSync(formatPath, 'utf-8');
    
    // Parse expected format to understand structure
    const expectedExample = JSON.parse(formatContent);
    // We assume the first item defines the schema if it exists
    const sampleKeyMapping: Record<string, string> = {};
    
    // Heuristic: Try to map CSV headers to JSON keys based on similarity or position
    // Since we don't have explicit mapping rules, we'll infer from the CSV header row
    const rows = parseCSV(csvContent);
    if (rows.length === 0) {
      console.log('[]');
      return;
    }

    const headers = rows[0].map(h => h.toLowerCase());
    const dataRows = rows.slice(1);

    // Map headers to expected JSON keys
    // Common mappings: first name, first, fname -> firstName
    // last name, last, lname -> lastName
    // date of birth, dob, birthdate -> dateOfBirth
    // email -> email
    const keyMap: Record<string, string> = {
      'first name': 'firstName', 'first': 'firstName', 'fname': 'firstName', 'given': 'firstName',
      'last name': 'lastName', 'last': 'lastName', 'lname': 'lastName', 'surname': 'lastName',
      'date of birth': 'dateOfBirth', 'dob': 'dateOfBirth', 'birth date': 'dateOfBirth', 'birthdate': 'dateOfBirth',
      'email': 'email', 'email address': 'email'
    };

    // Fallback: if we have expected format, try to match headers to it
    if (Array.isArray(expectedExample) && expectedExample.length > 0) {
      const sample = expectedExample[0];
      const expectedKeys = Object.keys(sample);
      
      // Simple mapping strategy: try to match header to expected key
      headers.forEach((header, index) => {
        const lowerHeader = header.toLowerCase();
        let mappedKey = keyMap[lowerHeader];
        
        if (!mappedKey) {
          // Try to find a key in the expected format that contains parts of the header
          for (const expKey of expectedKeys) {
            // Check if expKey (split by camelCase) exists in header
            const expKeyWords = expKey.replace(/([A-Z])/g, ' $1').toLowerCase().split(' ');
            const headerWords = lowerHeader.split(/[\s-]+/);
            
            // Simple overlap check
            const overlap = expKeyWords.filter(w => headerWords.some(h => h.includes(w) || w.includes(h)));
            if (overlap.length > 0) {
              mappedKey = expKey;
              break;
            }
          }
        }
        
        if (mappedKey) {
          sampleKeyMapping[lowerHeader] = mappedKey;
        }
      });
    } else {
      // Use default mappings
      headers.forEach((header, index) => {
        const lowerHeader = header.toLowerCase();
        if (keyMap[lowerHeader]) {
          sampleKeyMapping[lowerHeader] = keyMap[lowerHeader];
        } else {
          // Fallback: use header as key (cleaned)
          sampleKeyMapping[lowerHeader] = header; // Keep original or cleaned?
          // Let's clean it to snake_case or camelCase depending on expected format
          // If expected format exists, match its style
          const sample = Array.isArray(expectedExample) ? expectedExample[0] : {};
          const keys = Object.keys(sample);
          const isCamelCase = keys.some(k => /[A-Z]/.test(k));
          if (isCamelCase) {
            sampleKeyMapping[lowerHeader] = header.replace(/_\w/g, c => c.slice(1).toUpperCase());
          } else {
            sampleKeyMapping[lowerHeader] = header.toLowerCase().replace(/ /g, '_');
          }
        }
      });
    }

    const output: any[] = [];

    for (const row of dataRows) {
      const record: any = {};
      
      for (let i = 0; i < headers.length; i++) {
        const header = headers[i];
        const value = row[i];
        
        // Skip empty headers
        if (!header || !value) continue;

        let key = sampleKeyMapping[header];
        if (!key) continue;

        // Calculate age if it's the date of birth field
        if (header.toLowerCase().includes('dob') || header.toLowerCase().includes('birth')) {
          record['age'] = calculateAge(value);
          // Also keep dateOfBirth if it's expected
          if (Array.isArray(expectedExample) && expectedExample.length > 0 && expectedExample[0].hasOwnProperty('dateOfBirth')) {
             record['dateOfBirth'] = value;
          }
        } else {
          record[key] = value;
        }
      }

      output.push(record);
    }

    console.log(JSON.stringify(output, null, 2));

  } catch (error) {
    console.error('Error processing files:', error);
    process.exit(1);
  }
}

main();