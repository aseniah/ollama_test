import * as fs from 'fs';
import * as path from 'path';

const inputFile = 'input/input.csv';
const expectedFile = 'input/expected_format.json';
const referenceDate = new Date('2025-07-01');

function readAndParseCSV(filePath: string): { headers: string[], rows: Record<string, string>[] } {
  const content = fs.readFileSync(filePath, 'utf-8');
  const lines = content.trim().split('\n');
  
  if (lines.length === 0) {
    return { headers: [], rows: [] };
  }

  // Parse headers
  const headers = parseCSVLine(lines[0]);
  const rows: Record<string, string>[] = [];

  // Parse body
  for (let i = 1; i < lines.length; i++) {
    if (!lines[i].trim()) continue;
    
    const values = parseCSVLine(lines[i]);
    const row: Record<string, string> = {};
    
    headers.forEach((header, index) => {
      row[header] = values[index] || '';
    });
    
    rows.push(row);
  }

  return { headers, rows };
}

function parseCSVLine(line: string): string[] {
  const result: string[] = [];
  let current = '';
  let inQuotes = false;

  for (let i = 0; i < line.length; i++) {
    const char = line[i];
    
    if (char === '"') {
      inQuotes = !inQuotes;
      continue;
    }

    if (char === ',' && !inQuotes) {
      result.push(current.trim());
      current = '';
      continue;
    }

    current += char;
  }

  result.push(current.trim());
  return result;
}

function calculateAge(dateString: string): number | null {
  const date = new Date(dateString);
  if (isNaN(date.getTime())) return null;

  let age = referenceDate.getFullYear() - date.getFullYear();
  const m = referenceDate.getMonth() - date.getMonth();
  
  if (m < 0 || (m === 0 && referenceDate.getDate() < date.getDate())) {
    age--;
  }

  return age;
}

function inferTransformations(headers: string[], sampleRow: Record<string, string>): (row: Record<string, string>) => any[] {
  // Try to load expected format to understand the schema
  let expectedFormat = [] as any[];
  
  try {
    if (fs.existsSync(expectedFile)) {
      const expectedContent = fs.readFileSync(expectedFile, 'utf-8');
      expectedFormat = JSON.parse(expectedContent);
    }
  } catch (e) {
    // Fallback if file doesn't exist or is invalid
  }

  if (expectedFormat.length > 0) {
    // Use keys from expected format as output keys
    const targetKeys = Object.keys(expectedFormat[0]);
    
    return (row: Record<string, string>) => {
      const outputRecord: any[] = [];
      
      // We'll build an object first, then maybe convert to array if needed
      // But looking at the prompt "produce a JSON array", usually means array of objects
      
      // Let's map headers to expected keys based on naming patterns
      const mapping: Record<string, string> = {};
      
      for (const targetKey of targetKeys) {
        let sourceKey = targetKey;
        
        // Handle case variations
        const lowerTarget = targetKey.toLowerCase().replace(/[_-]/g, '');
        
        // Try exact match first
        if (headers.includes(targetKey)) {
          sourceKey = targetKey;
        } else {
          // Try to find a similar header
          for (const h of headers) {
            const lowerHead = h.toLowerCase().replace(/[_-]/g, '');
            if (lowerHead === lowerTarget || 
                (targetKey.includes('name') && lowerHead.includes('name')) ||
                (targetKey.includes('dob') && lowerHead.includes('birth')) ||
                (targetKey.includes('age') && lowerHead.includes('birth'))) {
              sourceKey = h;
              break;
            }
          }
        }
        
        mapping[targetKey] = sourceKey;
      }

      const result: Record<string, any> = {};
      
      for (const [targetKey, sourceKey] of Object.entries(mapping)) {
        if (sourceKey in row) {
          let value = row[sourceKey];
          
          // Special handling for age calculation
          if (targetKey.toLowerCase().includes('age')) {
            const age = calculateAge(value);
            result[targetKey] = age !== null ? age : null;
          } else {
            // Try to parse numbers
            if (/^\d+$/.test(value)) {
              result[targetKey] = parseInt(value, 10);
            } else if (/^\d+\.\d*$|^\.\d+$/.test(value)) {
              result[targetKey] = parseFloat(value);
            } else {
              result[targetKey] = value;
            }
          }
        } else {
          result[targetKey] = null;
        }
      }
      
      // Return as object (which will be part of the array)
      return Object.values(result); // If expected is an array, or return result if expected is object
    };
  }

  // Fallback: just parse values and calculate age for date fields
  return (row: Record<string, string>) => {
    const transformedRow: Record<string, any> = {};
    
    for (const [key, value] of Object.entries(row)) {
      const lowerKey = key.toLowerCase();
      
      if (lowerKey.includes('dob') || lowerKey.includes('birth')) {
        const age = calculateAge(value);
        transformedRow.age = age !== null ? age : null;
        // Keep the date too if needed
        transformedRow[key] = value;
      } else {
        // Try to parse numbers
        if (/^\d+$/.test(value)) {
          transformedRow[key] = parseInt(value, 10);
        } else if (/^\d+\.\d*$|^\.\d+$/.test(value)) {
          transformedRow[key] = parseFloat(value);
        } else {
          transformedRow[key] = value;
        }
      }
    }
    
    return transformedRow; // Will become an object in the array
  };
}

function main() {
  const { headers, rows } = readAndParseCSV(inputFile);
  
  if (rows.length === 0) {
    console.log(JSON.stringify([]));
    return;
  }

  const transform = inferTransformations(headers, rows[0]);
  const result = rows.map(transform);

  // Check if expected format shows array elements or objects
  let expectedFormat = [] as any[];
  try {
    if (fs.existsSync(expectedFile)) {
      const expectedContent = fs.readFileSync(expectedFile, 'utf-8');
      expectedFormat = JSON.parse(expectedContent);
    }
  } catch (e) {}

  // Determine output format based on expected
  let output: any;
  
  if (expectedFormat.length > 0 && Array.isArray(expectedFormat[0])) {
    // Expected elements are arrays
    output = result.map(row => {
      const transformed = transform(row);
      return typeof transformed === 'object' ? Object.values(transformed) : [transformed];
    });
  } else {
    // Expected elements are objects
    output = result;
  }

  console.log(JSON.stringify(output, null, 2));
}

main();