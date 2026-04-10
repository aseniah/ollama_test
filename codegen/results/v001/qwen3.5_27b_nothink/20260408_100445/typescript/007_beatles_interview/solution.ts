import * as fs from 'fs';
import * as path from 'path';

function parseCSV(content: string): string[][] {
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  const result: string[][] = [];

  for (const line of lines) {
    const row: string[] = [];
    let current = '';
    let inQuotes = false;

    for (const char of line) {
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === ',' && !inQuotes) {
        row.push(current);
        current = '';
      } else {
        current += char;
      }
    }
    row.push(current);
    result.push(row.map(cell => cell.trim()));
  }

  return result;
}

function calculateAge(dobStr: string, asOfDate: Date): number {
  const dob = new Date(dobStr);
  if (isNaN(dob.getTime())) return null;

  let age = asOfDate.getFullYear() - dob.getFullYear();
  const m = asOfDate.getMonth() - dob.getMonth();
  if (m < 0 || (m === 0 && asOfDate.getDate() < dob.getDate())) {
    age--;
  }
  return age;
}

function main() {
  const inputPath = path.join('input', 'input.csv');
  const expectedPath = path.join('input', 'expected_format.json');

  try {
    // Read CSV
    const csvContent = fs.readFileSync(inputPath, 'utf-8');
    const rows = parseCSV(csvContent);

    if (rows.length === 0) {
      console.log('[]');
      return;
    }

    const headers = rows[0];
    const dataRows = rows.slice(1);

    // Read expected format to infer keys and structure
    let expectedContent;
    try {
      expectedContent = fs.readFileSync(expectedPath, 'utf-8');
      const expected = JSON.parse(expectedContent);
      
      // Infer transformation rules from expected format
      const expectedKeys = Array.isArray(expected) && expected.length > 0 
        ? Object.keys(expected[0]) 
        : [];

      // Map headers to keys if possible
      // Common mappings based on typical datasets:
      // "first_name" -> "firstName"
      // "last_name" -> "lastName"
      // "birth_date" / "dob" / "date_of_birth" -> "age" (calculated)
      // "email" -> "email"
      // "city" -> "city"
      
      const dateRef = new Date('2025-07-01');
      
      const output: any[] = [];

      for (const row of dataRows) {
        const obj: any = {};

        // Simple column name mapping based on common CSV headers vs JSON keys
        for (let i = 0; i < headers.length; i++) {
          const header = headers[i].toLowerCase().replace(/\s+/g, '_');
          const value = row[i];

          if (!value) continue;

          if (['first_name', 'firstname'].includes(header)) {
            obj.firstName = value;
          } else if (['last_name', 'lastname'].includes(header)) {
            obj.lastName = value;
          } else if (['birth_date', 'date_of_birth', 'dob', 'date'].includes(header)) {
            const age = calculateAge(value, dateRef);
            if (age !== null) {
              obj.age = age;
            }
          } else if (['email', 'e_mail'].includes(header)) {
            obj.email = value;
          } else if (['city'].includes(header)) {
            obj.city = value;
          } else if (['country'].includes(header)) {
            obj.country = value;
          } else {
            // Fallback: try to infer snake_case to camelCase for unknown keys
            const camelCaseKey = header.replace(/_([a-z])/g, (g) => g[1].toUpperCase());
            obj[camelCaseKey] = value;
          }
        }

        // Reorder keys to match typical expected output order if inferred
        // We will just output based on the expected keys order if we have it
        const orderedObj: any = {};
        for (const key of expectedKeys) {
          if (obj[key] !== undefined) {
            orderedObj[key] = obj[key];
          }
        }
        
        // Add any extra keys not in expected (though unlikely based on prompt)
        for (const key of Object.keys(obj)) {
          if (!orderedObj[key]) orderedObj[key] = obj[key];
        }

        output.push(orderedObj);
      }

      console.log(JSON.stringify(output));

    } catch (e) {
      // If we can't read expected, we fall back to a standard assumption based on the CSV
      // This handles the case where expected file might be missing or invalid
      const dateRef = new Date('2025-07-01');
      const output: any[] = [];

      for (const row of dataRows) {
        const obj: any = {};
        for (let i = 0; i < headers.length; i++) {
          const header = headers[i].toLowerCase().replace(/\s+/g, '_');
          const value = row[i];
          if (!value) continue;
          
          if (['first_name', 'firstname'].includes(header)) obj.firstName = value;
          else if (['last_name', 'lastname'].includes(header)) obj.lastName = value;
          else if (['birth_date', 'date_of_birth', 'dob', 'date'].includes(header)) {
            const age = calculateAge(value, dateRef);
            if (age !== null) obj.age = age;
          }
          else if (['email'].includes(header)) obj.email = value;
          else if (['city'].includes(header)) obj.city = value;
          else if (['country'].includes(header)) obj.country = value;
          else {
             const camelCaseKey = header.replace(/_([a-z])/g, (g) => g[1].toUpperCase());
             obj[camelCaseKey] = value;
          }
        }
        output.push(obj);
      }
      console.log(JSON.stringify(output));
    }

  } catch (err) {
    console.error('Error processing files:', err);
    process.exit(1);
  }
}

main();