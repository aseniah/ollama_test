import * as fs from 'fs';

// Helper: Parse CSV line into an object
function parseCSVLine(line: string, headers: string[]): object {
  const parts = line.split(',');
  const result: any = {};
  headers.forEach((header, i) => {
    if (parts[i]) {
      result[header] = parts[i].trim();
    }
  });
  return result;
}

// Helper: Calculate age as of July 1, 2025
function calculateAge(dateStr: string, birthDate?: Date): number {
  if (dateStr) {
    const birth = new Date(dateStr);
    const referenceDate = new Date('2025-07-01');
    let age = referenceDate.getFullYear() - birth.getFullYear();
    // Account for not yet having had birthday in the reference year
    if (referenceDate < birth) {
      age--;
    }
    return age;
  }
  if (birthDate) {
    const referenceDate = new Date('2025-07-01');
    let age = referenceDate.getFullYear() - birthDate.getFullYear();
    // Account for not yet having had birthday in the reference year
    if (referenceDate < birthDate) {
      age--;
    }
    return age;
  }
  return 0;
}

// Read input file
const inputPath = 'input/input.csv';
const content = fs.readFileSync(inputPath, 'utf-8');
const lines = content.trim().split('\n').filter(line => line.trim() !== '');

if (lines.length === 0) {
  console.log('[]');
  process.exit(0);
}

// Read expected format to infer headers
const expectedPath = 'input/expected_format.json';
const expectedContent = fs.readFileSync(expectedPath, 'utf-8');
const expectedData = JSON.parse(expectedContent);

// Extract headers from expected output if available, otherwise assume first line of CSV
const headers = expectedData.length > 0 ? Object.keys(expectedData[0]) : null;

if (!headers) {
  // Infer from first CSV line
  const firstLine = lines[0];
  // Simple comma splitting for headers
  headers = firstLine.split(',').map(h => h.trim());
}

const outputData = lines.map(line => {
  const obj = parseCSVLine(line, headers);
  let age = 0;
  let dateStr: string | undefined;
  let birthDate: Date | undefined;

  // Try to find a date field
  for (const key in obj) {
    const val = obj[key];
    if (val && typeof val === 'string') {
      const d = new Date(val);
      if (!isNaN(d.getTime())) {
        dateStr = val;
        birthDate = d;
        break;
      }
    }
  }

  // Calculate age as of July 1, 2025
  if (birthDate) {
    age = calculateAge('', birthDate);
  } else {
    age = calculateAge(dateStr);
  }

  // Construct output object
  const result: any = { ...obj };
  result.age = age;
  return result;
});

// Sort by age (optional, but good practice if multiple people have same age or if expected format has sorting)
// Let's assume natural sorting by name if present, or just leave as is per input order
// Since the prompt doesn't specify sorting, we stick to input order but calculate age correctly.

// Output JSON
console.log(JSON.stringify(outputData, null, 2));