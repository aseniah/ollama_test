import * as fs from 'fs';
import * as path from 'path';

// Read the input CSV and expected format
const csvPath = path.join('input', 'input.csv');
const expectedPath = path.join('input', 'expected_format.json');

// Parse CSV content into an array of objects
function parseCSV(csvContent: string): any[] {
  const lines = csvContent.trim().split('\n');
  if (lines.length === 0) return [];

  const headers = lines[0].split(',').map(h => h.trim());
  const data: any[] = [];

  for (let i = 1; i < lines.length; i++) {
    // Handle quoted fields with commas
    let currentLine = lines[i];
    const values: string[] = [];
    let inQuotes = false;
    let field = '';

    for (let j = 0; j < currentLine.length; j++) {
      const char = currentLine[j];
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === ',' && !inQuotes) {
        values.push(field.trim().replace(/^"|"$/g, ''));
        field = '';
      } else {
        field += char;
      }
    }
    values.push(field.trim().replace(/^"|"$/g, ''));

    const row: any = {};
    headers.forEach((header, index) => {
      row[header] = values[index] || '';
    });
    data.push(row);
  }

  return data;
}

// Calculate age as of July 1, 2025
function calculateAge(birthDateString: string): number {
  if (!birthDateString) return 0;

  let birthDate: Date | null = null;

  // Try parsing in multiple formats
  const patterns = [
    /^(\d{4})-(\d{1,2})-(\d{1,2})$/,
    /^(\d{1,2})[\/-](\d{1,2})[\/-](\d{2,4})$/,
    /^(\d{1,2}) (\w+) (\d{4})$/,
  ];

  for (const pattern of patterns) {
    const match = birthDateString.match(pattern);
    if (match) {
      if (pattern === patterns[0]) {
        // YYYY-MM-DD
        birthDate = new Date(`${match[1]}-${parseInt(match[2], 10).toString().padStart(2, '0')}-${parseInt(match[3], 10).toString().padStart(2, '0')}`);
      } else if (pattern === patterns[1]) {
        // MM/DD/YYYY or DD/MM/YYYY - try to infer
        const m1 = parseInt(match[1], 10);
        const m2 = parseInt(match[2], 10);
        let y = parseInt(match[3], 10);
        if (y < 100) y += 2000;

        if (m1 > 12) birthDate = new Date(`${y}-${m1.toString().padStart(2, '0')}-${m2.toString().padStart(2, '0')}`);
        else birthDate = new Date(`${y}-${m2.toString().padStart(2, '0')}-${m1.toString().padStart(2, '0')}`);
      } else if (pattern === patterns[2]) {
        // DD MonthName YYYY
        const monthNames: string[] = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
        let monthIndex = -1;
        for (let i = 0; i < monthNames.length; i++) {
          if (monthNames[i].toLowerCase() === match[2].toLowerCase()) {
            monthIndex = i + 1;
            break;
          }
        }
        if (monthIndex !== -1) {
          birthDate = new Date(`${match[3]}-${monthIndex.toString().padStart(2, '0')}-${parseInt(match[1], 10).toString().padStart(2, '0')}`);
        }
      }
      break;
    }
  }

  if (!birthDate || isNaN(birthDate.getTime())) return 0;

  const referenceDate = new Date('2025-07-01');
  let age = referenceDate.getFullYear() - birthDate.getFullYear();

  const monthDayComparison =
    (referenceDate.getMonth() > birthDate.getMonth()) ||
    (referenceDate.getMonth() === birthDate.getMonth() && referenceDate.getDate() >= birthDate.getDate());

  if (!monthDayComparison) {
    age--;
  }

  return age;
}

// Read expected format to understand the output structure
const expectedContent = fs.readFileSync(expectedPath, 'utf8', 'utf-8');
const expectedFormat = JSON.parse(expectedContent);

// Read and parse CSV
const csvContent = fs.readFileSync(csvPath, 'utf8');
const data = parseCSV(csvContent);

// Determine target field names from expected format
function mapFields(row: any): any {
  const transformed: any = {};
  if (!expectedFormat || !Array.isArray(expectedFormat) || expectedFormat.length === 0) {
    // Fallback: use original keys
    Object.keys(row).forEach(key => {
      transformed[key] = row[key];
    });

    const birthKey = Object.keys(transformed).find(k => k.toLowerCase().includes('birth') || k.toLowerCase().includes('dob'));
    if (birthKey) transformed.age = calculateAge(transformed[birthKey]);
    return transformed;
  }

  const sampleKeys = Object.keys(expectedFormat[0]);

  for (const targetKey of sampleKeys) {
    if (targetKey === 'age') {
      const birthKeys = ['birthday', 'dob', 'date_of_birth', 'birth_date', 'birth'];
      let value = '';
      for (const key of Object.keys(row)) {
        if (birthKeys.some(bk => bk.toLowerCase() === key.toLowerCase())) {
          value = row[key];
          break;
        }
      }
      transformed[targetKey] = calculateAge(value);
    } else {
      // Try direct match or case-insensitive match
      const matchedKey = Object.keys(row).find(k => k.toLowerCase() === targetKey.toLowerCase());
      if (matchedKey) {
        transformed[targetKey] = row[matchedKey];
      }
    }
  }

  return transformed;
}

const result: any[] = data.map(mapFields);

console.log(JSON.stringify(result));