import * as fs from 'fs';
import * as path from 'path';

// Read the file
const filePath = 'input/data.csv';
let data: string;

try {
  // Ensure we read from the relative path correctly
  const fullPath = path.join(process.cwd(), filePath);
  data = fs.readFileSync(fullPath, 'utf8');
} catch (error) {
  // Fallback in case the script is run directly and cwd handling differs
  data = fs.readFileSync(filePath, 'utf8');
}

// Split into lines and filter out empty ones
const lines = data.split('\n').filter(line => line.trim() !== '');

// Skip header
const rows = lines.slice(1);

// Helper to handle CSV parsing including commas inside quotes (if any)
function parseCSVRow(line: string): string[] {
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
  result.push(current.trim());
  return result;
}

const parsedData: { Name: string; Age: number; Email: string; Score: number }[] = [];

for (const row of rows) {
  if (!row.trim()) continue;

  const fields = parseCSVRow(row);
  
  // Ensure we have at least 4 fields to avoid runtime errors on malformed lines
  if (fields.length < 4) continue;

  parsedData.push({
    Name: fields[0],
    Age: parseInt(fields[1], 10),
    Email: fields[2],
    Score: parseFloat(fields[3])
  });
}

// Output as valid JSON to stdout
console.log(JSON.stringify(parsedData));