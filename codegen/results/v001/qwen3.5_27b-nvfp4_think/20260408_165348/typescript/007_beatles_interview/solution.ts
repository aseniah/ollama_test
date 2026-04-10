import * as fs from 'fs';
import * as path from 'path';

const REFERENCE_DATE = new Date('2025-07-01');

function parseCSV(csvContent: string): { headers: string[]; rows: string[][] } {
  const lines = csvContent.trim().split('\n');
  if (lines.length === 0) {
    return { headers: [], rows: [] };
  }

  const headers = lines[0].split(',').map(h => h.trim());
  const rows: string[][] = [];

  for (let i = 1; i < lines.length; i++) {
    if (lines[i].trim()) {
      rows.push(lines[i].split(',').map(cell => cell.trim()));
    }
  }

  return { headers, rows };
}

function calculateAge(birthDateString: string): number | null {
  const birthDate = new Date(birthDateString);
  if (isNaN(birthDate.getTime())) {
    return null;
  }

  let age = REFERENCE_DATE.getFullYear() - birthDate.getFullYear();
  const monthDiff = REFERENCE_DATE.getMonth() - birthDate.getMonth();
  
  if (
    monthDiff < 0 ||
    (monthDiff === 0 && REFERENCE_DATE.getDate() < birthDate.getDate())
  ) {
    age--;
  }

  return age;
}

function inferTypeFromSampleValues(values: string[]): 'string' | 'number' {
  const numericCount = values.filter(v => !isNaN(Number(v)) && v.trim() !== '').length;
  return numericCount > 0 ? 'number' : 'string';
}

function transformCSV(csvContent: string, expectedFormat: any[]): any[] {
  const { headers, rows } = parseCSV(csvContent);
  
  if (headers.length === 0 || rows.length === 0) {
    return [];
  }

  const sampleValues: Record<string, string[]> = {};
  for (let i = 0; i < rows.length; i++) {
    headers.forEach((header, j) => {
      if (!sampleValues[header]) {
        sampleValues[header] = [];
      }
      if (rows[i][j]) {
        sampleValues[header].push(rows[i][j]);
      }
    });
  }

  const fieldTypes: Record<string, 'string' | 'number'> = {};
  headers.forEach(header => {
    fieldTypes[header] = inferTypeFromSampleValues(sampleValues[header] || []);
  });

  const result: any[] = [];

  rows.forEach((row, rowIndex) => {
    const record: Record<string, any> = {};

    expectedFormat.forEach(field => {
      const fieldName = field.name;
      
      for (const header of headers) {
        const normalizedHeader = header.toLowerCase().replace(/\s+/g, '').trim();
        
        if (header === 'birthday' || header === 'birth_date' || header === 'birthdate') {
          record[fieldName] = calculateAge(row[headers.indexOf(header)]);
          return;
        }

        if (fieldName.includes('age')) {
          const birthdateIndex = headers.findIndex(h => 
            h.toLowerCase() === 'birthday' || 
            h.toLowerCase() === 'birth_date' || 
            h.toLowerCase() === 'birthdate'
          );
          if (birthdateIndex !== -1) {
            record[fieldName] = calculateAge(row[headers.indexOf(header)]);
          }
          return;
        }

        if (header.toLowerCase() === fieldName.toLowerCase()) {
          const value = row[headers.indexOf(header)];
          if (fieldTypes[header] === 'number' && value) {
            record[fieldName] = Number(value);
          } else {
            record[fieldName] = value;
          }
          return;
        }

        const normalizedField = fieldName.toLowerCase().replace(/\s+/g, '').trim();
        if (normalizedHeader === normalizedField) {
          const value = row[headers.indexOf(header)];
          if (fieldTypes[header] === 'number' && value) {
            record[fieldName] = Number(value);
          } else {
            record[fieldName] = value;
          }
          return;
        }

        // Check for partial matches or common field name variations
        const headerMap: Record<string, string[]> = {
          'name': ['name', 'full_name', 'fullname', 'first_last'],
          'id': ['id', 'employee_id', 'emp_id', 'uid'],
          'email': ['email', 'email_address'],
          'phone': ['phone', 'telephone', 'mobile'],
          'address': ['address', 'street_address', 'full_address'],
          'city': ['city', 'town'],
          'state': ['state', 'province', 'region'],
          'zip': ['zip', 'postal_code', 'postcode', 'zipcode'],
          'country': ['country', 'nation'],
        };

        for (const [mappedField, variations] of Object.entries(headerMap)) {
          if (fieldTypes[mappedField] === 'number') continue;
          
          if (variations.some(v => header.toLowerCase() === v || normalizedHeader === mappedField)) {
            const value = row[headers.indexOf(header)];
            record[fieldName] = value || '';
            return;
          }
        }
      }

      // If no match found, try to find a column that could contain this field
      const matchedHeaderIndex = headers.findIndex((h, idx) => {
        if (row[idx]) {
          return h.toLowerCase() === fieldName.toLowerCase() || 
                 fieldName.toLowerCase() === h.toLowerCase().replace(/\s+/g, '');
        }
        return false;
      });

      if (matchedHeaderIndex !== -1) {
        const value = row[matchedHeaderIndex];
        record[fieldName] = value;
      } else {
        record[fieldName] = null;
      }
    });

    result.push(record);
  });

  return result;
}

try {
  const inputCsvPath = process.argv[2] || 'input/input.csv';
  const expectedFormatPath = process.argv[3] || 'input/expected_format.json';

  const csvContent = fs.readFileSync(inputCsvPath, 'utf-8');
  const expectedFormat = JSON.parse(fs.readFileSync(expectedFormatPath, 'utf-8'));

  const transformedData = transformCSV(csvContent, Array.isArray(expectedFormat) ? expectedFormat : []);

  process.stdout.write(JSON.stringify(transformedData, null, 2));
} catch (error) {
  process.stderr.write('Error: ' + (error as Error).message + '\n');
  process.exit(1);
}