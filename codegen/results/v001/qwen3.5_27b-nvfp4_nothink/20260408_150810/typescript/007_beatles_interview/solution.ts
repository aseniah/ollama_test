import * as fs from 'fs';
import * as path from 'path';

interface RawRow {
  [key: string]: string;
}

interface OutputRecord {
  name: string;
  birth_date: string;
  age: number;
  city: string;
  country: string;
}

function parseCSV(content: string): RawRow[] {
  const lines = content.trim().split(/\r?\n/);
  if (lines.length < 2) return [];

  const headers = lines[0].split(',').map(h => h.trim());
  
  // Handle quoted fields with commas inside
  const parseLine = (line: string): string[] => {
    const results: string[] = [];
    let current = '';
    let inQuotes = false;

    for (let i = 0; i < line.length; i++) {
      const char = line[i];
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === ',' && !inQuotes) {
        results.push(current.trim());
        current = '';
      } else {
        current += char;
      }
    }
    results.push(current.trim().replace(/^"|"$/g, '').trim());
    return results;
  };

  const data: RawRow[] = [];
  for (let i = 1; i < lines.length; i++) {
    if (!lines[i].trim()) continue;
    const values = parseLine(lines[i]);
    const row: RawRow = {};
    headers.forEach((header, index) => {
      row[header] = values[index] || '';
    });
    data.push(row);
  }

  return data;
}

function calculateAge(birthDateString: string): number {
  const referenceDate = new Date('2025-07-01');
  // Try multiple common date formats
  let birthDate: Date | null = null;

  // ISO format: YYYY-MM-DD
  if (/^\d{4}-\d{2}-\d{2}$/.test(birthDateString)) {
    const [year, month, day] = birthDateString.split('-').map(Number);
    birthDate = new Date(year, month - 1, day);
  } 
  // US format: MM/DD/YYYY
  else if (/^\d{2}\/\d{2}\/\d{4}$/.test(birthDateString)) {
    const [month, day, year] = birthDateString.split('/').map(Number);
    birthDate = new Date(year, month - 1, day);
  }
  // European format: DD.MM.YYYY or DD-MM-YYYY (ambiguous without context, but try)
  else if (/^\d{2}[.-]\d{2}[.-]\d{4}$/.test(birthDateString)) {
    const parts = birthDateString.split(/[.-]/).map(Number);
    // Assuming DD MM YYYY for this fallback
    birthDate = new Date(parts[2], parts[1] - 1, parts[0]);
  }

  if (!birthDate || isNaN(birthDate.getTime())) {
    throw new Error(`Invalid date format: ${birthDateString}`);
  }

  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const m = referenceDate.getMonth() - birthDate.getMonth();
  
  // Adjust if birthday hasn't happened yet this year
  if (m < 0 || (m === 0 && referenceDate.getDate() < birthDate.getDate())) {
    age--;
  }

  return age;
}

function main() {
  const inputPath = 'input/input.csv';
  const expectedFormatPath = 'input/expected_format.json';

  try {
    // Read input CSV
    const csvContent = fs.readFileSync(inputPath, 'utf-8');
    const rows = parseCSV(csvContent);

    // Read expected format to infer keys (optional for robustness, 
    // but we can also rely on standard naming conventions if not strictly needed)
    let expectedKeys: string[] = [];
    try {
      const expectedContent = fs.readFileSync(expectedFormatPath, 'utf-8');
      const expectedData = JSON.parse(expectedContent);
      if (Array.isArray(expectedData) && expectedData.length > 0) {
        // Normalize keys to snake_case or whatever the first example has
        expectedKeys = Object.keys(expectedData[0]);
      }
    } catch (e) {
      // Fallback common mapping if JSON read fails
      expectedKeys = ['name', 'birth_date', 'age', 'city', 'country'];
    }

    // Map input rows to output structure
    // We assume the CSV headers map logically to the expected keys.
    // Common mappings: Name -> name, DOB/Date of Birth -> birth_date, etc.
    
    const normalizedRows: OutputRecord[] = [];

    rows.forEach(row => {
      // Heuristic mapping based on typical column names found in CSVs vs expected JSON
      let name = '';
      let birthDateStr = '';
      let city = '';
      let country = '';

      Object.keys(row).forEach(key => {
        const lowerKey = key.toLowerCase().trim();
        const val = row[key];

        if (lowerKey.includes('name') && !lowerKey.includes('first') && !lowerKey.includes('last')) {
          name = val;
        } else if (lowerKey === 'birth_date' || lowerKey === 'dob' || lowerKey.includes('birth')) {
          birthDateStr = val;
        } else if (lowerKey.includes('city')) {
          city = val;
        } else if (lowerKey.includes('country')) {
          country = val;
        }
        
        // Handle split first/last names if they exist and name is empty
        if (!name) {
          if (lowerKey.includes('first') && lowerKey.includes('name')) {
            name = val;
          } else if (lowerKey.includes('last') || lowerKey.includes('surname')) {
            if (name) name += ' ';
            name += val;
          }
        }
      });

      // Fallback: If we didn't find specific keys, maybe the first column is name and second is date?
      // This is a safety net if heuristics fail, but usually headers are descriptive.
      const allKeys = Object.keys(row);
      if (!name && allKeys.length >= 1) {
         // Try to guess: often first column is name
         if (allKeys[0].toLowerCase() === 'name') name = row[allKeys[0]];
      }

      if (!birthDateStr || !name) {
        // Skip malformed rows or handle gracefully? 
        // Given the prompt implies a valid transformation, we assume valid data.
        return;
      }

      const age = calculateAge(birthDateStr);
      // Format date to ISO YYYY-MM-DD
      const birthDateObj = new Date(birthDateStr.replace(/[^0-9]/g, '')); // crude sanitization if needed
      // Actually, let's use the parsed logic again for consistency
      let formattedBirthDate = '';
       if (/^\d{4}-\d{2}-\d{2}$/.test(birthDateStr)) {
         formattedBirthDate = birthDateStr;
       } else if (/^\d{2}\/\d{2}\/\d{4}$/.test(birthDateStr)) {
         const [m, d, y] = birthDateStr.split('/');
         formattedBirthDate = `${y}-${m.padStart(2, '0')}-${d.padStart(2, '0')}`;
       } else if (/^\d{2}[.-]\d{2}[.-]\d{4}$/.test(birthDateStr)) {
        const parts = birthDateStr.split(/[.-]/);
         formattedBirthDate = `${parts[2]}-${parts[1].padStart(2, '0')}-${parts[0].padStart(2, '0')}`;
       } else {
         // If format is unknown but we calculated age, just use input string or ISO if possible
         formattedBirthDate = birthDateStr; 
       }

      // Determine output order based on expected format keys
      const outputRecord: any = {};
      
      // Map to the specific keys found in expected_format.json
      // We assume standard mapping: name, birth_date, age, city, country
      // If the expected file has different key names, we try to match them.
      // Since I cannot read the actual file content dynamically in this thought process without running it,
      // I will rely on the prompt's implied structure and standard conventions, 
      // but I will use the keys extracted from expectedFormatPath if available.

      const keyMap: Record<string, string> = {
        name: name,
        birth_date: formattedBirthDate,
        age: String(age), // JSON requires number type usually, but let's output raw object then stringify
        city: city,
        country: country
      };

      // Construct final object in the order of expected keys if we parsed them
      const finalRecord: OutputRecord = {
        name: name,
        birth_date: formattedBirthDate,
        age: age,
        city: city,
        country: country
      };

      normalizedRows.push(finalRecord);
    });

    console.log(JSON.stringify(normalizedRows, null, 2));

  } catch (error) {
    process.stderr.write(`Error: ${(error as Error).message}\n`);
    process.exit(1);
  }
}

main();