import * as fs from 'fs';
import * as path from 'path';

// Read the CSV file
const filePath = path.join('input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');

// Split into lines and filter out empty lines
const lines = content.split('\n').filter(line => line.trim() !== '');

// Skip the header (first line)
const dataLines = lines.slice(1);

// Function to parse CSV line handling quoted fields
function parseCSVLine(line: string): string[] {
  const result: string[] = [];
  let current = '';
  let inQuotes = false;

  for (let i = 0; i < line.length; i++) {
    const char = line[i];
    
    if (char === '"') {
      inQuotes = !inQuotes;
    } else if (char === ',' && !inQuotes) {
      result.push(current.trim());
      current = '';
    } else {
      current += char;
    }
  }
  
  if (current.length > 0 || line.endsWith(',')) {
    result.push(current.trim());
  }
  
  return result;
}

// Parse each line into an object
const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const line of dataLines) {
  const values = parseCSVLine(line);
  result.push({
    Name: values[0],
    Age: parseInt(values[1], 10),
    Email: values[2],
    Score: parseFloat(values[3])
  });
}

// Output as JSON to stdout
console.log(JSON.stringify(result, null, 2));