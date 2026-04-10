import { readFileSync } from 'fs';

// Read CSV file
const csvContent = readFileSync('input/input.csv', 'utf-8');

// Read expected format
const expectedFormat = JSON.parse(readFileSync('input/expected_format.json', 'utf-8'));

// Parse CSV and calculate ages as of July 1, 2025
function parseCSV(line: string) {
  return line.split(',').map(s => s.trim());
}

function parseDate(dateStr: string) {
  return new Date(dateStr);
}

function calculateAge(birthDate: Date, referenceDate: Date = new Date('2025-07-01')) {
  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  const birthMonth = birthDate.getMonth() + 1;
  const referenceMonth = referenceDate.getMonth() + 1;
  
  if (birthMonth > referenceMonth || (birthMonth === referenceMonth && birthDate.getDate() > referenceDate.getDate())) {
    age -= 1;
  }
  return age;
}

function formatOutput(row: any[]): any[] {
  const data = row.map((r, i) => {
    if (r.length > 0) {
      return {
        id: r[0],
        age: calculateAge(parseDate(r[1]))
      };
    }
    return {};
  }).filter(obj => obj.id && obj.age !== undefined);
  
  return data;
}

// Parse CSV content
const lines = csvContent.trim().split('\n').filter(line => line.trim());

// Handle header and data rows
const linesData = lines.filter(line => !line.includes('Name') || !line.includes('Date') || !line.includes('ID') || !line.includes('Age'));

if (linesData.length === 0) {
  console.log(JSON.stringify([]));
  process.exit(0);
}

const results = linesData.map(line => parseCSV(line));

// Calculate and format
const output = results.map(row => formatOutput(row)).flat();

// Output JSON
console.log(JSON.stringify(output));